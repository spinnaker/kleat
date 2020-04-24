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
	"os"
	"path/filepath"

	"github.com/spinnaker/kleat/internal/validate_paths"
	"github.com/spinnaker/kleat/pkg/parse_hal"
	"github.com/spinnaker/kleat/pkg/validate_hal"
	"sigs.k8s.io/yaml"
)

func WriteConfigs(hal string, d string) error {
	if err := validate_paths.EnsureFile(hal); err != nil {
		return err
	}
	if err := validate_paths.EnsureDirectory(d); err != nil {
		return err
	}

	h, err := parse_hal.ParseHalConfig(hal)
	if err != nil {
		return err
	}

	if err := validate_hal.ValidateHalConfig(h); err != nil {
		return err
	}

	c := parse_hal.HalToClouddriver(h)
	if err := write(c, d, "clouddriver.yml"); err != nil {
		return err
	}

	e := parse_hal.HalToEcho(h)
	if err := write(e, d, "echo.yml"); err != nil {
		return err
	}

	return nil
}

func write(i interface{}, d string, f string) error {
	w, err := os.Create(filepath.Join(d, f))
	if err != nil {
		return err
	}

	bytes, err := yaml.Marshal(i)
	if err != nil {
		return err
	}
	if _, err = w.Write(bytes); err != nil {
		return err
	}
	return nil
}
