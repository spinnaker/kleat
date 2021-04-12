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
	"github.com/spinnaker/kleat/api/client/notification"
	"github.com/spinnaker/kleat/api/client/security"
	"github.com/spinnaker/kleat/api/client/security/authn"
	"github.com/spinnaker/kleat/internal/convert"
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
					FeatureDisabled: wrapperspb.Bool(true),
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
							Enabled: wrapperspb.Bool(true),
						},
					},
				},
			},
			&config.Deck{
				GateUrl:         "https://localhost:8084",
				AuthEndpoint:    "https://localhost:8084/auth/user",
				BakeryDetailUrl: "https://localhost:8084/bakery/logs/{{context.region}}/{{context.status.resourceId}}",
				Canary: &config.Deck_Canary{
					FeatureDisabled: wrapperspb.Bool(true),
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
					Authn: &authn.Authentication{Enabled: wrapperspb.Bool(true)},
				},
			},
			&config.Deck{
				AuthEnabled:     wrapperspb.Bool(true),
				GateUrl:         "http://localhost:8084",
				AuthEndpoint:    "http://localhost:8084/auth/user",
				BakeryDetailUrl: "http://localhost:8084/bakery/logs/{{context.region}}/{{context.status.resourceId}}",
				Canary: &config.Deck_Canary{
					FeatureDisabled: wrapperspb.Bool(true),
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
					Enabled:               wrapperspb.Bool(true),
					DefaultMetricsAccount: "my-metrics-account",
					DefaultMetricsStore:   "datadog",
					ShowAllConfigsEnabled: wrapperspb.Bool(true),
					TemplatesEnabled:      wrapperspb.Bool(true),
					StorageAccountName:    "my-storage-account",
				},
			},
			&config.Deck{
				GateUrl:         "http://localhost:8084",
				AuthEndpoint:    "http://localhost:8084/auth/user",
				BakeryDetailUrl: "http://localhost:8084/bakery/logs/{{context.region}}/{{context.status.resourceId}}",
				Canary: &config.Deck_Canary{
					FeatureDisabled:    wrapperspb.Bool(false),
					MetricsAccountName: "my-metrics-account",
					MetricStore:        "datadog",
					ShowAllConfigs:     wrapperspb.Bool(true),
					StorageAccountName: "my-storage-account",
					TemplatesEnabled:   wrapperspb.Bool(true),
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
					FeatureDisabled: wrapperspb.Bool(true),
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
					PipelineTemplates:            wrapperspb.Bool(true),
					Chaos:                        wrapperspb.Bool(true),
					ManagedPipelineTemplatesV2UI: wrapperspb.Bool(true),
				},
			},
			&config.Deck{
				GateUrl:         "http://localhost:8084",
				AuthEndpoint:    "http://localhost:8084/auth/user",
				BakeryDetailUrl: "http://localhost:8084/bakery/logs/{{context.region}}/{{context.status.resourceId}}",
				Canary: &config.Deck_Canary{
					FeatureDisabled: wrapperspb.Bool(true),
				},
				Feature: &config.Deck_Features{
					PipelineTemplates:            wrapperspb.Bool(true),
					ManagedPipelineTemplatesV2UI: wrapperspb.Bool(true),
					ChaosMonkey:                  wrapperspb.Bool(true),
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
						Enabled: wrapperspb.Bool(true),
					},
				},
			},
			&config.Deck{
				GateUrl:         "http://localhost:8084",
				AuthEndpoint:    "http://localhost:8084/auth/user",
				BakeryDetailUrl: "http://localhost:8084/bakery/logs/{{context.region}}/{{context.status.resourceId}}",
				Canary: &config.Deck_Canary{
					FeatureDisabled: wrapperspb.Bool(true),
				},
				Feature: &config.Deck_Features{},
				Notifications: &config.Deck_Notifications{
					Sms: &notification.Twilio{
						Enabled: wrapperspb.Bool(true),
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
						Enabled: wrapperspb.Bool(true),
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
					FeatureDisabled: wrapperspb.Bool(true),
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
			"Enabled provider without defaults",
			&config.Hal{
				Providers: &cloudprovider.Providers{
					Kubernetes: &cloudprovider.Kubernetes{
						Enabled: wrapperspb.Bool(true),
					},
				},
			},
			&config.Deck{
				GateUrl:         "http://localhost:8084",
				AuthEndpoint:    "http://localhost:8084/auth/user",
				BakeryDetailUrl: "http://localhost:8084/bakery/logs/{{context.region}}/{{context.status.resourceId}}",
				Canary: &config.Deck_Canary{
					FeatureDisabled: wrapperspb.Bool(true),
				},
				Feature:       &config.Deck_Features{},
				Notifications: &config.Deck_Notifications{},
				Providers: &config.Deck_Providers{
					Kubernetes: &config.Deck_Providers_Kubernetes{},
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
					FeatureDisabled: wrapperspb.Bool(true),
				},
				Feature:       &config.Deck_Features{},
				Notifications: &config.Deck_Notifications{},
				Providers:     &config.Deck_Providers{},
			},
		},
		{
			"Enable ManagedDelivery",
			&config.Hal{
				ManagedDelivery: &config.ManagedDelivery{
					Enabled: wrapperspb.Bool(true),
				},
			},
			&config.Deck{
				GateUrl:         "http://localhost:8084",
				AuthEndpoint:    "http://localhost:8084/auth/user",
				BakeryDetailUrl: "http://localhost:8084/bakery/logs/{{context.region}}/{{context.status.resourceId}}",
				Canary: &config.Deck_Canary{
					FeatureDisabled: wrapperspb.Bool(true),
				},
				Feature: &config.Deck_Features{
					ManagedDelivery:  wrapperspb.Bool(true),
					ManagedResources: wrapperspb.Bool(true),
				},
				Notifications: &config.Deck_Notifications{},
				Providers:     &config.Deck_Providers{},
			},
		},
	},
}

func TestHalToDeck(t *testing.T) {
	runConfigTest(t, deckTests)
}
