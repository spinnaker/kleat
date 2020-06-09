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
package write

import (
	"github.com/spinnaker/kleat/internal/serializer"
	"google.golang.org/protobuf/proto"
)

func writeDeck(m proto.Message, file string) error {
	js, err := serializer.SettingsJs.Serialize(m)
	if err != nil {
		return err
	}
	if err := writeBytes(js, file); err != nil {
		return err
	}
	return nil
}

func writeDeckEnv(m proto.Message, file string) error {
	js, err := serializer.EnvFile.Serialize(m)
	if err != nil {
		return err
	}
	if err := writeBytes(js, file); err != nil {
		return err
	}
	return nil
}
