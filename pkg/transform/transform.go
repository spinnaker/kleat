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

// Package transform supports transforming between various formats of Spinnaker
// configuration.
package transform

import (
	"google.golang.org/protobuf/proto"

	"github.com/spinnaker/kleat/api/client/config"
	"github.com/spinnaker/kleat/internal/convert"
	"github.com/spinnaker/kleat/internal/serializer"
)

// HalToServiceConfigs returns the microservice configurations corresponding
// to a supplied *config.Hal.
func HalToServiceConfigs(h *config.Hal) *config.Services {
	return &config.Services{
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
		Monitoring:  convert.HalToMonitoring(h),
		Keel:        convert.HalToKeel(h),
	}
}

// GenerateConfigFiles generates the config files corresponding to a supplied
// *config.Services.
func GenerateConfigFiles(s *config.Services) (*config.ConfigFiles, error) {
	var files = []struct {
		fileName   string
		message    proto.Message
		serializer serializer.Serializer
	}{
		{
			"clouddriver.yml",
			s.GetClouddriver(),
			serializer.Yaml,
		},
		{
			"echo.yml",
			s.GetEcho(),
			serializer.Yaml,
		},
		{
			"front50.yml",
			s.GetFront50(),
			serializer.Yaml,
		},
		{
			"gate.yml",
			s.GetGate(),
			serializer.Yaml,
		},
		{
			"fiat.yml",
			s.GetFiat(),
			serializer.Yaml,
		},
		{
			"igor.yml",
			s.GetIgor(),
			serializer.Yaml,
		},
		{
			"kayenta.yml",
			s.GetKayenta(),
			serializer.Yaml,
		},
		{
			"orca.yml",
			s.GetOrca(),
			serializer.Yaml,
		},
		{
			"rosco.yml",
			s.GetRosco(),
			serializer.Yaml,
		},
		{
			"settings.js",
			s.GetDeck(),
			serializer.SettingsJs,
		},
		{
			"deck.env",
			s.GetDeckEnv(),
			serializer.EnvFile,
		},
		{
			"spinnaker-monitoring.yml",
			s.GetMonitoring(),
			serializer.Yaml,
		},
		{
			"keel.yml",
			s.GetKeel(),
			serializer.Yaml,
		},
	}

	cf := make([]*config.ConfigFile, len(files))
	for i, file := range files {
		contents, err := file.serializer.Serialize(file.message)
		if err != nil {
			return nil, err
		}
		cf[i] = &config.ConfigFile{
			Name:     file.fileName,
			Contents: contents,
		}
	}

	return &config.ConfigFiles{ConfigFile: cf}, nil
}
