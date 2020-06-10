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
package transform_test

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/kylelemons/godebug/diff"
	"github.com/spinnaker/kleat/internal/convert"

	"github.com/spinnaker/kleat/api/client/config"
	"github.com/spinnaker/kleat/internal/protoyaml"
	"github.com/spinnaker/kleat/pkg/transform"
	"google.golang.org/protobuf/proto"
)

const dataDir = "../../testdata"

func TestHalToServiceConfigs(t *testing.T) {
	data, err := ioutil.ReadFile(filepath.Join(dataDir, "halconfig.yml"))
	if err != nil {
		t.Fatal(err)
	}
	h := &config.Hal{}
	if err := protoyaml.Unmarshal(data, h); err != nil {
		t.Fatal(err)
	}

	gotS := transform.HalToServiceConfigs(h)
	wantS := &config.Services{
		Clouddriver: convert.HalToClouddriver(h),
		Echo:        convert.HalToEcho(h),
		Front50:     convert.HalToFront50(h),
		Orca:        convert.HalToOrca(h),
		Gate:        convert.HalToGate(h),
		Fiat:        convert.HalToFiat(h),
		Kayenta:     convert.HalToKayenta(h),
		Rosco:       convert.HalToRosco(h),
		Deck:        convert.HalToDeck(h),
		DeckEnv:     convert.HalToDeckEnv(h),
		Igor:        convert.HalToIgor(h),
	}

	if !proto.Equal(gotS, wantS) {
		t.Errorf("Expected HalToServiceConfigs to generate map of service configs, got %v", gotS)
	}
}

func TestHalToServiceYAML(t *testing.T) {
	data, err := ioutil.ReadFile(filepath.Join(dataDir, "halconfig.yml"))
	if err != nil {
		t.Fatal(err)
	}

	hal := &config.Hal{}
	if err := protoyaml.Unmarshal(data, hal); err != nil {
		t.Fatal(err)
	}

	services := transform.HalToServiceConfigs(hal)

	var halToServiceTests = []struct {
		file      string
		gotConfig proto.Message
	}{
		{
			"clouddriver.yml",
			services.GetClouddriver(),
		},
		{
			"echo.yml",
			services.GetEcho(),
		},
		{
			"front50.yml",
			services.GetFront50(),
		},
		{
			"orca.yml",
			services.GetOrca(),
		},
		{
			"gate.yml",
			services.GetGate(),
		},
		{
			"fiat.yml",
			services.GetFiat(),
		},
		{
			"kayenta.yml",
			services.GetKayenta(),
		},
		{
			"rosco.yml",
			services.GetRosco(),
		},
		{
			"deck.yml",
			services.GetDeck(),
		},
		{
			"igor.yml",
			services.GetIgor(),
		},
	}

	for _, tt := range halToServiceTests {
		t.Run(tt.file, func(t *testing.T) {
			got, err := protoyaml.Marshal(tt.gotConfig)
			if err != nil {
				t.Fatal(err)
			}

			data, err := ioutil.ReadFile(filepath.Join(dataDir, tt.file))
			if err != nil {
				t.Fatal(err)
			}

			// Note that we're re-using tt.gotConfig here and will unmarshal the expected
			// YAML into it; this is acceptable as we have already converted tt.gotConfig to
			// YAML and have the YAML representation stored in got.
			// (We can't just create a new pointer as we don't have access to the concrete
			// type from within this function, so we'll re-use gotConfig rather than require
			// the test to pass in an empty config of the concrete type.)
			want, err := canonicalYaml(data, tt.gotConfig)
			if err != nil {
				t.Fatal(err)
			}
			if !bytes.Equal(got, want) {
				d := diff.Diff(string(got), string(want))
				t.Errorf("Generated %s differs from expected:\n%s", tt.file, d)
			}
		})
	}
}

// canonicalYaml converts the YAML representation of a proto.Message into its
// canonical form (ie, the form produced by serializing the object it represents)
// by unmarshaling then marshaling the provided YAML.
// The motivation is that the tests want to compare the generated YAML to
// expected YAML, but we don't want to be sensitive to key order in the provided
// YAML.
func canonicalYaml(data []byte, m proto.Message) ([]byte, error) {
	if err := protoyaml.UnmarshalStrict(data, m); err != nil {
		return nil, err
	}
	return protoyaml.Marshal(m)
}
