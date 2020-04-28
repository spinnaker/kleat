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
	"github.com/spinnaker/kleat/internal/protoyaml"
	"github.com/spinnaker/kleat/pkg/parse_hal"
)

var halToClouddriverTests = []struct {
	n string
	h string
	c string
}{
	{"Empty hal", "./data/hal/empty.yml", "./data/clouddriver/empty.yml"},
	{"Empty providers", "./data/hal/empty_providers.yml", "./data/clouddriver/empty.yml"},
	{"Empty Kubernetes provider", "./data/hal/empty_kubernetes_provider.yml", "./data/clouddriver/empty_kubernetes_provider.yml"},
	{"Kubernetes account", "./data/hal/kubernetes_account.yml", "./data/clouddriver/kubernetes_account.yml"},
	{"Empty artifacts", "./data/hal/empty_artifacts.yml", "./data/clouddriver/empty_artifacts.yml"},
	{"Empty GCS artifacts", "./data/hal/empty_gcs_artifacts.yml", "./data/clouddriver/empty_gcs_artifacts.yml"},
	{"GCS artifact account", "./data/hal/gcs_artifact_account.yml", "./data/clouddriver/gcs_artifact_account.yml"},
}

func TestHalToClouddriver(t *testing.T) {
	for _, tt := range halToClouddriverTests {
		t.Run(tt.n, func(t *testing.T) {
			h, err := parse_hal.ParseHalConfig(tt.h)
			if err != nil {
				t.Errorf(err.Error())
			}
			gotC := parse_hal.HalToClouddriver(h)
			wantC, err := parseClouddriverConfig(tt.c)
			if err != nil {
				t.Errorf(err.Error())
			}
			want, err := protoyaml.Marshal(wantC)
			if err != nil {
				t.Errorf(err.Error())
			}
			got, err := protoyaml.Marshal(gotC)
			if err != nil {
				t.Errorf(err.Error())
			}
			res := bytes.Compare(want, got)
			if res != 0 {
				t.Errorf("Expected hal config to generate %v, got %v", wantC, gotC)
			}
		})
	}
}

func parseClouddriverConfig(fn string) (*config.Clouddriver, error) {
	dat, err := ioutil.ReadFile(fn)
	h := config.Clouddriver{}
	err = protoyaml.UnmarshalStrict([]byte(dat), &h)
	if err != nil {
		return nil, err
	}
	return &h, nil
}
