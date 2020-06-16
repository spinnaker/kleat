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

package convert_test

import (
	"testing"

	"github.com/spinnaker/kleat/api/client/canary"
	"github.com/spinnaker/kleat/internal/convert"
	"github.com/spinnaker/kleat/internal/wrappers"

	"github.com/spinnaker/kleat/api/client"
	"github.com/spinnaker/kleat/api/client/cloudprovider"
	"github.com/spinnaker/kleat/api/client/config"
	"github.com/spinnaker/kleat/api/client/security"
	"google.golang.org/protobuf/proto"
)

var orcaTests = configTest{
	generator: func(h *config.Hal) proto.Message { return convert.HalToOrca(h) },
	tests: []testCase{
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
						Enabled:        wrappers.False(),
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
						Enabled:        wrappers.True(),
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
				Features: &client.Features{PipelineTemplates: wrappers.False()},
			},
			&config.Orca{},
		},
		{
			"Pipeline templates enabled",
			&config.Hal{
				Features: &client.Features{PipelineTemplates: wrappers.True()},
			},
			&config.Orca{
				PipelineTemplates: &config.Orca_PipelineTemplates{Enabled: wrappers.True()},
			},
		},
		{
			"Webhook trust configured",
			&config.Hal{
				Webhook: &security.WebhookConfig{
					Trust: &security.TrustStore{
						Enabled:            wrappers.True(),
						TrustStore:         "/var/secrets/store.jks",
						TrustStorePassword: "passw0rd",
					},
				},
			},
			&config.Orca{
				Webhook: &security.WebhookConfig{
					Trust: &security.TrustStore{
						Enabled:            wrappers.True(),
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
					Enabled: wrappers.True(),
				},
			},
			&config.Orca{
				Services: &config.Orca_Services{
					Kayenta: &config.ServiceSettings{
						Enabled: wrappers.True(),
					},
				},
			},
		},
		{
			"Timezone set",
			&config.Hal{
				Timezone: "America/Chicago",
			},
			&config.Orca{
				Tasks: &config.Orca_Tasks{
					ExecutionWindow: &config.Orca_Tasks_ExecutionWindow{
						Timezone: "America/Chicago",
					},
				},
			},
		},
	},
}

func TestHalToOrca(t *testing.T) {
	runConfigTest(t, orcaTests)
}
