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
	"google.golang.org/protobuf/proto"

	"github.com/spinnaker/kleat/api/client/config"
	"github.com/spinnaker/kleat/internal/convert"
	"github.com/spinnaker/kleat/internal/fileio"
	"github.com/spinnaker/kleat/internal/prototest"
	"github.com/spinnaker/kleat/internal/protoyaml"
	"github.com/spinnaker/kleat/pkg/transform"
)

const dataDir = "../../testdata"

func TestHalToServiceConfigs(t *testing.T) {
	h, err := fileio.ParseHalConfig(filepath.Join(dataDir, "halconfig.yml"))
	if err != nil {
		t.Fatal(err)
	}

	gotS := transform.HalToServiceConfigs(h)
	wantS := &config.Services{
		Clouddriver: convert.HalToClouddriver(h),
		Deck:        convert.HalToDeck(h),
		DeckEnv:     convert.HalToDeckEnv(h),
		Echo:        convert.HalToEcho(h),
		Fiat:        convert.HalToFiat(h),
		Front50:     convert.HalToFront50(h),
		Gate:        convert.HalToGate(h),
		Igor:        convert.HalToIgor(h),
		Kayenta:     convert.HalToKayenta(h),
		Monitoring:  convert.HalToMonitoring(h),
		Orca:        convert.HalToOrca(h),
		Rosco:       convert.HalToRosco(h),
		Keel:        convert.HalToKeel(h),
	}

	if !proto.Equal(gotS, wantS) {
		t.Errorf("Expected HalToServiceConfigs to generate map of service configs, got %v", gotS)
	}
}

func TestHalToServiceIntegration(t *testing.T) {
	hal, err := fileio.ParseHalConfig(filepath.Join(dataDir, "halconfig.yml"))
	if err != nil {
		t.Fatal(err)
	}

	services := transform.HalToServiceConfigs(hal)

	var halToServiceTests = []struct {
		file string
		got  proto.Message
	}{
		{
			"clouddriver.yml",
			services.GetClouddriver(),
		},
		{
			"deck.yml",
			services.GetDeck(),
		},
		{
			"deck-env.yml",
			services.GetDeckEnv(),
		},
		{
			"echo.yml",
			services.GetEcho(),
		},
		{
			"fiat.yml",
			services.GetFiat(),
		},
		{
			"front50.yml",
			services.GetFront50(),
		},
		{
			"gate.yml",
			services.GetGate(),
		},
		{
			"igor.yml",
			services.GetIgor(),
		},
		{
			"kayenta.yml",
			services.GetKayenta(),
		},
		{
			"spinnaker-monitoring.yml",
			services.GetMonitoring(),
		},
		{
			"orca.yml",
			services.GetOrca(),
		},
		{
			"rosco.yml",
			services.GetRosco(),
		},
		{
			"keel.yml",
			services.GetKeel(),
		},
	}

	for _, tt := range halToServiceTests {
		t.Run(tt.file, func(t *testing.T) {
			data, err := ioutil.ReadFile(filepath.Join(dataDir, tt.file))
			if err != nil {
				t.Fatal(err)
			}

			// Clone tt.got so that we unmarshal into a message of the same type;
			// the unmarshalling will overwrite any data in want.
			want := proto.Clone(tt.got)
			err = protoyaml.UnmarshalStrict(data, want)
			if err != nil {
				t.Fatal(err)
			}

			if equal, d := prototest.Equal(tt.got, want); !equal {
				t.Errorf("Incorrect service config generated:\n%s", d)
			}
		})
	}
}

func TestHalToServiceYAML(t *testing.T) {
	hal, err := fileio.ParseHalConfig(filepath.Join(dataDir, "halconfig.yml"))
	if err != nil {
		t.Fatal(err)
	}

	services := transform.HalToServiceConfigs(hal)
	configs, err := transform.GenerateConfigFiles(services)
	if err != nil {
		t.Fatal(err)
	}

	files := []string{
		"clouddriver.yml",
		"echo.yml",
		"fiat.yml",
		"front50.yml",
		"gate.yml",
		"igor.yml",
		"kayenta.yml",
		"orca.yml",
		"rosco.yml",
		"keel.yml",
		"spinnaker-monitoring.yml",
	}

	for _, file := range files {
		t.Run(file, func(t *testing.T) {
			gotConfig := getConfig(configs, file)
			if err != nil {
				t.Fatal(err)
			}

			wantConfig, err := ioutil.ReadFile(filepath.Join(dataDir, file))
			if err != nil {
				t.Fatal(err)
			}

			if !bytes.Equal(gotConfig, wantConfig) {
				d := diff.Diff(string(gotConfig), string(wantConfig))
				t.Errorf("Generated %s differs from expected:\n%s", file, d)
			}
		})
	}
}

func getConfig(c *config.ConfigFiles, name string) []byte {
	for _, file := range c.GetConfigFile() {
		if file.GetName() == name {
			return file.GetContents()
		}
	}
	return nil
}
