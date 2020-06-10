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

	"github.com/spinnaker/kleat/api/client/cloudprovider"
	"github.com/spinnaker/kleat/api/client/notification"
	"github.com/spinnaker/kleat/internal/convert"
	"google.golang.org/protobuf/proto"

	"github.com/spinnaker/kleat/api/client"

	"github.com/spinnaker/kleat/api/client/canary"

	"github.com/spinnaker/kleat/api/client/security/authn"

	"github.com/spinnaker/kleat/api/client/security"

	"github.com/spinnaker/kleat/api/client/config"
)

var deckTests = configTest{
	generator: func(h *config.Hal) proto.Message { return convert.HalToDeck(h) },
	tests: []testCase{
		{
			"Empty hal config",
			&config.Hal{},
			&config.Deck{
				GateUrl:         "http://localhost:8084",
				AuthEndpoint:    "http://localhost:8084/auth/user",
				BakeryDetailUrl: "http://localhost:8084/bakery/logs/{{context.region}}/{{context.status.resourceId}}",
				Canary: &config.Deck_Canary{
					FeatureDisabled: true,
				},
				Feature:       &config.Deck_Features{},
				Notifications: &config.Deck_Notifications{},
				Providers:     &config.Deck_Providers{},
			},
		},
		{
			"SSL enabled",
			&config.Hal{
				Security: &security.Security{
					ApiSecurity: &security.ApiSecurity{
						Ssl: &security.ApiSsl{
							Enabled: true,
						},
					},
				},
			},
			&config.Deck{
				GateUrl:         "https://localhost:8084",
				AuthEndpoint:    "https://localhost:8084/auth/user",
				BakeryDetailUrl: "https://localhost:8084/bakery/logs/{{context.region}}/{{context.status.resourceId}}",
				Canary: &config.Deck_Canary{
					FeatureDisabled: true,
				},
				Feature:       &config.Deck_Features{},
				Notifications: &config.Deck_Notifications{},
				Providers:     &config.Deck_Providers{},
			},
		},
		{
			"Authn enabled",
			&config.Hal{
				Security: &security.Security{
					Authn: &authn.Authentication{Enabled: true},
				},
			},
			&config.Deck{
				AuthEnabled:     true,
				GateUrl:         "http://localhost:8084",
				AuthEndpoint:    "http://localhost:8084/auth/user",
				BakeryDetailUrl: "http://localhost:8084/bakery/logs/{{context.region}}/{{context.status.resourceId}}",
				Canary: &config.Deck_Canary{
					FeatureDisabled: true,
				},
				Feature:       &config.Deck_Features{},
				Notifications: &config.Deck_Notifications{},
				Providers:     &config.Deck_Providers{},
			},
		},
		{
			"Canary configured",
			&config.Hal{
				Canary: &canary.Canary{
					Enabled:               true,
					DefaultMetricsAccount: "my-metrics-account",
					DefaultMetricsStore:   "datadog",
					ShowAllConfigsEnabled: true,
					TemplatesEnabled:      true,
					StorageAccountName:    "my-storage-account",
				},
			},
			&config.Deck{
				GateUrl:         "http://localhost:8084",
				AuthEndpoint:    "http://localhost:8084/auth/user",
				BakeryDetailUrl: "http://localhost:8084/bakery/logs/{{context.region}}/{{context.status.resourceId}}",
				Canary: &config.Deck_Canary{
					MetricsAccountName: "my-metrics-account",
					MetricStore:        "datadog",
					ShowAllConfigs:     true,
					StorageAccountName: "my-storage-account",
					TemplatesEnabled:   true,
				},
				Feature:       &config.Deck_Features{},
				Notifications: &config.Deck_Notifications{},
				Providers:     &config.Deck_Providers{},
			},
		},
		{
			"Timezone configured",
			&config.Hal{
				Timezone: "America/Chicago",
			},
			&config.Deck{
				DefaultTimeZone: "America/Chicago",
				GateUrl:         "http://localhost:8084",
				AuthEndpoint:    "http://localhost:8084/auth/user",
				BakeryDetailUrl: "http://localhost:8084/bakery/logs/{{context.region}}/{{context.status.resourceId}}",
				Canary: &config.Deck_Canary{
					FeatureDisabled: true,
				},
				Feature:       &config.Deck_Features{},
				Notifications: &config.Deck_Notifications{},
				Providers:     &config.Deck_Providers{},
			},
		},
		{
			"Features configured",
			&config.Hal{
				Features: &client.Features{
					PipelineTemplates:            true,
					MineCanary:                   true,
					Chaos:                        true,
					ManagedPipelineTemplatesV2UI: true,
				},
			},
			&config.Deck{
				GateUrl:         "http://localhost:8084",
				AuthEndpoint:    "http://localhost:8084/auth/user",
				BakeryDetailUrl: "http://localhost:8084/bakery/logs/{{context.region}}/{{context.status.resourceId}}",
				// TODO: this test demonstrates less-than-ideal result of having
				// multiple canary flags that users must enable: both canary.enabled
				// and features.mineCanary must be set to true in order for the
				// canary UI and Kayenta to work properly.
				Canary: &config.Deck_Canary{
					FeatureDisabled: true,
				},
				Feature: &config.Deck_Features{
					PipelineTemplates:            true,
					ManagedPipelineTemplatesV2UI: true,
					ChaosMonkey:                  true,
					Canary:                       true,
				},
				Notifications: &config.Deck_Notifications{},
				Providers:     &config.Deck_Providers{},
			},
		},
		{
			"Notifications configured",
			&config.Hal{
				Notifications: &notification.Notifications{
					Twilio: &notification.Twilio{
						Enabled: true,
					},
				},
			},
			&config.Deck{
				GateUrl:         "http://localhost:8084",
				AuthEndpoint:    "http://localhost:8084/auth/user",
				BakeryDetailUrl: "http://localhost:8084/bakery/logs/{{context.region}}/{{context.status.resourceId}}",
				Canary: &config.Deck_Canary{
					FeatureDisabled: true,
				},
				Feature: &config.Deck_Features{},
				Notifications: &config.Deck_Notifications{
					Sms: &notification.Twilio{
						Enabled: true,
					},
				},
				Providers: &config.Deck_Providers{},
			},
		},
		{
			"Provider defaults",
			&config.Hal{
				Providers: &cloudprovider.Providers{
					Google: &cloudprovider.GoogleComputeEngine{
						Enabled: true,
						Accounts: []*cloudprovider.GoogleComputeEngineAccount{
							{
								Name:    "my-gce-account",
								Regions: []string{"us-east1", "us-east2"},
							},
						},
						PrimaryAccount: "my-gce-account",
					},
				},
			},
			&config.Deck{
				GateUrl:         "http://localhost:8084",
				AuthEndpoint:    "http://localhost:8084/auth/user",
				BakeryDetailUrl: "http://localhost:8084/bakery/logs/{{context.region}}/{{context.status.resourceId}}",
				Canary: &config.Deck_Canary{
					FeatureDisabled: true,
				},
				Feature:       &config.Deck_Features{},
				Notifications: &config.Deck_Notifications{},
				Providers: &config.Deck_Providers{
					Gce: &config.Deck_Providers_Gce{
						Defaults: &config.Deck_Providers_Gce_Defaults{
							Account: "my-gce-account",
							Region:  "us-east1",
						},
					},
				},
			},
		},
		{
			"Override gate URL",
			&config.Hal{
				Security: &security.Security{
					ApiSecurity: &security.ApiSecurity{
						OverrideBaseUrl: "https://spinnaker.test:8084",
					},
				},
			},
			&config.Deck{
				GateUrl:         "https://spinnaker.test:8084",
				AuthEndpoint:    "https://spinnaker.test:8084/auth/user",
				BakeryDetailUrl: "https://spinnaker.test:8084/bakery/logs/{{context.region}}/{{context.status.resourceId}}",
				Canary: &config.Deck_Canary{
					FeatureDisabled: true,
				},
				Feature:       &config.Deck_Features{},
				Notifications: &config.Deck_Notifications{},
				Providers:     &config.Deck_Providers{},
			},
		},
	},
}

func TestHalToDeck(t *testing.T) {
	runConfigTest(t, deckTests)
}
