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
package test

import (
	"bytes"
	"github.com/ghodss/yaml"
	"github.com/spinnaker/kleat/api/client"
	"github.com/spinnaker/kleat/pkg/parse_hal"
	"io/ioutil"
	"testing"
)

func TestHalToClouddriver(t *testing.T) {
	h, err := parse_hal.ParseHalConfig("./data/halconfig.yml")
	if err != nil {
		t.Errorf(err.Error())
	}

	gotC, err := parse_hal.HalToClouddriver(*h)
	if err != nil {
		t.Errorf(err.Error())
	}

	wantC, err := parseClouddriverConfig("./data/clouddriver.yml")
	if err != nil {
		t.Errorf(err.Error())
	}

	want, err := yaml.Marshal(wantC)
	if err != nil {
		t.Errorf(err.Error())
	}

	got, err := yaml.Marshal(gotC)
	if err != nil {
		t.Errorf(err.Error())
	}

	res := bytes.Compare(want, got)
	if res != 0 {
		t.Errorf("Expected generated Clouddriver config to match contents of ./data/clouddriver.yml, but got:\n" + string(got))
	}

}

func parseClouddriverConfig(fn string) (*client.ClouddriverConfig, error) {
	dat, err := ioutil.ReadFile(fn)

	h := client.ClouddriverConfig{}
	err = yaml.Unmarshal([]byte(dat), &h)
	if err != nil {
		return nil, err
	}

	return &h, nil
}
