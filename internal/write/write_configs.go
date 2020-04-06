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
	"github.com/ghodss/yaml"
	"github.com/spinnaker/kleat/internal/parse"
	"github.com/spinnaker/kleat/internal/validate"
	"io"
	"os"
	"path/filepath"
)

func WriteConfigs(hal string, d string) error {
	if err := validate.EnsureFile(hal); err != nil {
		return err
	}
	if err := validate.EnsureDirectory(d); err != nil {
		return err
	}

	h, err := parse.ParseHalConfig(hal)
	if err != nil {
		return err
	}

	if err := validate.ValidateHalConfig(h); err != nil {
		return err
	}

	c, err := parse.HalToClouddriver(*h)
	if err != nil {
		return err
	}

	f, err := os.Create(filepath.Join(d, "clouddriver.yml"))
	if err != nil {
		return err
	}

	if err = write(c, f); err != nil {
		return err
	}

	return nil
}

func write(i interface{}, w io.Writer) error {
	bytes, err := yaml.Marshal(i)
	if err != nil {
		return err
	}
	if _, err = w.Write(bytes); err != nil {
		return err
	}
	return nil
}
