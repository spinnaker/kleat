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

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/spinnaker/kleat/api/client"
	"github.com/spinnaker/kleat/api/client/canary"
	"github.com/spinnaker/kleat/api/client/cloudprovider"
	"github.com/spinnaker/kleat/api/client/config"
	"github.com/spinnaker/kleat/api/client/security"
	"github.com/spinnaker/kleat/internal/convert"
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
						Enabled:        wrapperspb.Bool(false),
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
						Enabled:        wrapperspb.Bool(true),
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
				Features: &client.Features{PipelineTemplates: wrapperspb.Bool(false)},
			},
			&config.Orca{
				PipelineTemplates: &config.Orca_PipelineTemplates{
					Enabled: wrapperspb.Bool(false),
				},
			},
		},
		{
			"Pipeline templates enabled",
			&config.Hal{
				Features: &client.Features{PipelineTemplates: wrapperspb.Bool(true)},
			},
			&config.Orca{
				PipelineTemplates: &config.Orca_PipelineTemplates{Enabled: wrapperspb.Bool(true)},
			},
		},
		{
			"Webhook trust configured",
			&config.Hal{
				Webhook: &security.WebhookConfig{
					Trust: &security.TrustStore{
						Enabled:            wrapperspb.Bool(true),
						TrustStore:         "/var/secrets/store.jks",
						TrustStorePassword: "passw0rd",
					},
				},
			},
			&config.Orca{
				Webhook: &security.WebhookConfig{
					Trust: &security.TrustStore{
						Enabled:            wrapperspb.Bool(true),
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
					Enabled: wrapperspb.Bool(true),
				},
			},
			&config.Orca{
				Services: &config.Orca_Services{
					Kayenta: &config.ServiceSettings{
						Enabled: wrapperspb.Bool(true),
					},
				},
			},
		},
		{
			"ManagedDelivery enabled",
			&config.Hal{
				ManagedDelivery: &config.ManagedDelivery{
					Enabled: wrapperspb.Bool(true),
				},
			},
			&config.Orca{
				Services: &config.Orca_Services{
					Keel: &config.ServiceSettings{
						Enabled: wrapperspb.Bool(true),
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
