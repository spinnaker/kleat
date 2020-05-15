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
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spinnaker/kleat/api/client/config"
	"github.com/spinnaker/kleat/internal/protoyaml"
	"github.com/spinnaker/kleat/internal/validate_paths"
	"github.com/spinnaker/kleat/pkg/parse_hal"
	"github.com/spinnaker/kleat/pkg/validate_hal"
	"google.golang.org/protobuf/proto"
)

func WriteConfigs(halPath string, dir string) error {
	if err := validate_paths.EnsureDirectory(dir); err != nil {
		return err
	}

	hal := &config.Hal{}
	if err := read(hal, halPath); err != nil {
		return err
	}

	if err := validate_hal.ValidateHalConfig(hal); err != nil {
		return err
	}

	services := parse_hal.HalToServiceConfigs(hal)
	if err := write(services.GetClouddriver(), filepath.Join(dir, "clouddriver.yml")); err != nil {
		return err
	}
	if err := write(services.GetEcho(), filepath.Join(dir, "echo.yml")); err != nil {
		return err
	}
	if err := write(services.GetFront50(), filepath.Join(dir, "front50.yml")); err != nil {
		return err
	}
	if err := write(services.GetGate(), filepath.Join(dir, "gate.yml")); err != nil {
		return err
	}
	if err := write(services.GetFiat(), filepath.Join(dir, "fiat.yml")); err != nil {
		return err
	}
	return nil
}

func read(m proto.Message, file string) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	return protoyaml.Unmarshal(data, m)
}

func write(m proto.Message, file string) error {
	w, err := os.Create(file)
	if err != nil {
		return err
	}

	bytes, err := protoyaml.Marshal(m)
	if err != nil {
		return err
	}
	if _, err = w.Write(bytes); err != nil {
		return err
	}
	return nil
}
