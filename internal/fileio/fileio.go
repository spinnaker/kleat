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

// Package fileio supports reading and writing config files to the
// filesystem.
package fileio

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spinnaker/kleat/api/client/config"
	"github.com/spinnaker/kleat/internal/protoyaml"
	"github.com/spinnaker/kleat/internal/validate"
	"github.com/spinnaker/kleat/pkg/transform"
)

// ParseHalConfig reads the YAML file at halPath parses it into a *config.Hal.
// The config.Hal is validated; if there are any errors, they will be returned
// in error and *config.Hal will be nil.
func ParseHalConfig(halPath string) (*config.Hal, error) {
	data, err := ioutil.ReadFile(halPath)
	if err != nil {
		return nil, fmt.Errorf("unable to read %q. Error: %v", halPath, err)
	}

	hal := &config.Hal{}
	if err := protoyaml.Unmarshal(data, hal); err != nil {
		return nil, fmt.Errorf("unable to unmarshal %q: %v", halPath, err)
	}

	if err := validate.HalConfig(hal); err != nil {
		return nil, fmt.Errorf("unable to validate %q: %v", halPath, err)
	}

	return hal, nil
}

// WriteConfigs generates the service configs from the supplied hal, and writes
// them to the directory dir.
func WriteConfigs(hal *config.Hal, dir string) error {
	if err := ensureDirectory(dir); err != nil {
		return err
	}

	services := transform.HalToServiceConfigs(hal)
	configFiles, err := transform.GenerateConfigFiles(services)
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

func ensureDirectory(d string) error {
	stat, err := os.Stat(d)
	if os.IsNotExist(err) {
		return os.MkdirAll(d, 0755)
	} else if err != nil {
		return err
	}
	if !stat.IsDir() {
		return fmt.Errorf("%s is not a directory", d)
	}
	return nil
}
