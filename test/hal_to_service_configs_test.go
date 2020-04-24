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
	"reflect"
	"testing"

	"github.com/spinnaker/kleat/api/client/config"

	"github.com/spinnaker/kleat/pkg/parse_hal"
)

func TestHalToServiceConfigs(t *testing.T) {
	h, err := parse_hal.ParseHalConfig("./data/halconfig.yml")
	if err != nil {
		t.Errorf(err.Error())
	}

	gotS, err := parse_hal.HalToServiceConfigs(h)
	if err != nil {
		t.Errorf(err.Error())
	}

	c, err := parse_hal.HalToClouddriver(h)
	if err != nil {
		t.Errorf(err.Error())
	}
	e, err := parse_hal.HalToEcho(h)
	if err != nil {
		t.Errorf(err.Error())
	}
	f, err := parse_hal.HalToFront50(h)
	if err != nil {
		t.Errorf(err.Error())
	}

	wantS := &config.Services{
		Clouddriver: c,
		Echo:        e,
		Front50:     f,
	}

	if !reflect.DeepEqual(gotS, wantS) {
		t.Errorf("Expected HalToServiceConfigs to generate map of service configs, got %v", gotS)
	}
}
