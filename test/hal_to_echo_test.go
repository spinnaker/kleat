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
	"io/ioutil"
	"testing"

	"github.com/spinnaker/kleat/api/client/config"

	"github.com/spinnaker/kleat/pkg/parse_hal"
	"sigs.k8s.io/yaml"
)

func TestHalToEcho(t *testing.T) {
	h, err := parse_hal.ParseHalConfig("./data/halconfig.yml")
	if err != nil {
		t.Errorf(err.Error())
	}

	gotE := parse_hal.HalToEcho(h)

	wantE, err := parseEchoConfig("./data/echo.yml")
	if err != nil {
		t.Errorf(err.Error())
	}

	want, err := yaml.Marshal(wantE)
	if err != nil {
		t.Errorf(err.Error())
	}

	got, err := yaml.Marshal(gotE)
	if err != nil {
		t.Errorf(err.Error())
	}

	res := bytes.Compare(want, got)
	if res != 0 {
		t.Errorf("Expected generated Echo config to match contents of ./data/echo.yml, but got:\n" + string(got))
	}
}

func parseEchoConfig(fn string) (*config.Echo, error) {
	dat, err := ioutil.ReadFile(fn)

	h := config.Echo{}
	err = yaml.UnmarshalStrict([]byte(dat), &h)
	if err != nil {
		return nil, err
	}

	return &h, nil
}
