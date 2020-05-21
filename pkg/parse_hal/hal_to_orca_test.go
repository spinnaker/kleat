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
package parse_hal_test

import (
	"testing"

	"github.com/spinnaker/kleat/api/client/canary"

	"github.com/spinnaker/kleat/api/client"
	"github.com/spinnaker/kleat/api/client/cloudprovider"
	"github.com/spinnaker/kleat/api/client/config"
	"github.com/spinnaker/kleat/api/client/security"
	"github.com/spinnaker/kleat/pkg/parse_hal"
	"google.golang.org/protobuf/proto"
)

var halToOrcaTests = []struct {
	name string
	hal  *config.Hal
	want *config.Orca
}{
	{
		"Empty hal config",
		&config.Hal{},
		&config.Orca{},
	},
	{
		"AWS disabled",
		&config.Hal{
			Providers: &cloudprovider.Providers{
				Aws: &cloudprovider.Aws{
					Enabled:        false,
					PrimaryAccount: "my-account",
				},
			},
		},
		&config.Orca{},
	},
	{
		"AWS enabled",
		&config.Hal{
			Providers: &cloudprovider.Providers{
				Aws: &cloudprovider.Aws{
					Enabled:        true,
					PrimaryAccount: "my-account",
				},
			},
		},
		&config.Orca{
			Default: &config.Orca_Defaults{
				Bake: &config.Orca_Defaults_BakeDefaults{
					Account: "my-account",
				},
			},
		},
	},
	{
		// TODO: This should probably explicitly set templates to disabled rather
		// than return an empty config.
		"Pipeline templates disabled",
		&config.Hal{
			Features: &client.Features{PipelineTemplates: false},
		},
		&config.Orca{},
	},
	{
		"Pipeline templates enabled",
		&config.Hal{
			Features: &client.Features{PipelineTemplates: true},
		},
		&config.Orca{
			PipelineTemplates: &config.Orca_PipelineTemplates{Enabled: true},
		},
	},
	{
		"Webhook trust configured",
		&config.Hal{
			Webhook: &security.WebhookConfig{
				Trust: &security.TrustStore{
					Enabled:            true,
					TrustStore:         "/var/secrets/store.jks",
					TrustStorePassword: "passw0rd",
				},
			},
		},
		&config.Orca{
			Webhook: &security.WebhookConfig{
				Trust: &security.TrustStore{
					Enabled:            true,
					TrustStore:         "/var/secrets/store.jks",
					TrustStorePassword: "passw0rd",
				},
			},
		},
	},
	{
		"Canary enabled",
		&config.Hal{
			Canary: &canary.Canary{
				Enabled: true,
			},
		},
		&config.Orca{
			Services: &config.Orca_Services{
				Kayenta: &config.ServiceEnabled{
					Enabled: true,
				},
			},
		},
	},
}

func TestHalToOrca(t *testing.T) {
	for _, tt := range halToOrcaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := parse_hal.HalToOrca(tt.hal)
			if !proto.Equal(got, tt.want) {
				t.Errorf("Expected hal config to generate %v, got %v", tt.want, got)
			}
		})
	}
}
