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
	"path/filepath"

	"github.com/spinnaker/kleat/api/client/config"
	"github.com/spinnaker/kleat/internal/protoyaml"
	"github.com/spinnaker/kleat/internal/validate_paths"
	"github.com/spinnaker/kleat/pkg/parse_hal"
	"github.com/spinnaker/kleat/pkg/validate_hal"
)

func ParseHalConfig(halPath string) (*config.Hal, error) {
	data, err := ioutil.ReadFile(halPath)
	if err != nil {
		return nil, err
	}

	hal := &config.Hal{}
	if err := protoyaml.Unmarshal(data, hal); err != nil {
		return nil, err
	}

	if err := validate_hal.ValidateHalConfig(hal); err != nil {
		return nil, err
	}

	return hal, nil
}

func WriteConfigs(hal *config.Hal, dir string) error {
	if err := validate_paths.EnsureDirectory(dir); err != nil {
		return err
	}

	services := parse_hal.HalToServiceConfigs(hal)
	configFiles, err := parse_hal.GenerateConfigFiles(services)
	if err != nil {
		return err
	}

	for _, file := range configFiles.GetConfigFile() {
		if err := ioutil.WriteFile(filepath.Join(dir, file.GetName()), file.GetContents(), 0666); err != nil {
			return err
		}
	}

	return nil
}
