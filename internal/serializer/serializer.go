/*
 * Copyright The Spinnaker Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package serializer

import (
	"bytes"
	"fmt"

	"github.com/spinnaker/kleat/internal/protoyaml"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Serializer interface {
	Serialize(proto.Message) ([]byte, error)
}

type yaml struct{}

var Yaml Serializer = new(yaml)

func (yaml) Serialize(m proto.Message) ([]byte, error) {
	return protoyaml.Marshal(m)
}

type settingsJs struct{}

var SettingsJs Serializer = new(settingsJs)

func (settingsJs) Serialize(m proto.Message) ([]byte, error) {
	json, err := protojson.Marshal(m)
	if err != nil {
		return nil, err
	}

	const jsStart = "window.spinnakerSettings = JSON.parse('"
	const jsEnd = "');"
	js := append([]byte(jsStart), json...)
	js = append(js, []byte(jsEnd)...)

	return js, nil
}

type envFile struct{}

var EnvFile Serializer = new(envFile)

func (envFile) Serialize(m proto.Message) ([]byte, error) {
	var buf bytes.Buffer
	var err error
	r := m.ProtoReflect()
	r.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		if _, err = fmt.Fprintf(&buf, "%s=%s\n", fd.JSONName(), v); err != nil {
			return false
		}
		return true
	})
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
