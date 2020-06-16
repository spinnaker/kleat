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

	"github.com/spinnaker/kleat/api/client/config"
	"google.golang.org/protobuf/proto"
)

var kayentaTests = configTest{
	generator: func(h *config.Hal) proto.Message { return convert.HalToKayenta(h) },
	tests: []testCase{
		{
			"Empty hal config",
			&config.Hal{},
			&config.Kayenta{},
		},
		{
			"Canary disabled",
			&config.Hal{
				Canary: &canary.Canary{
					Enabled: wrappers.False(),
				},
			},
			&config.Kayenta{},
		},
		{
			"Stackdriver enabled",
			&config.Hal{
				Canary: &canary.Canary{
					Enabled: wrappers.True(),
					ServiceIntegrations: &canary.Canary_ServiceIntegrations{
						Google: &canary.Google{
							Enabled: wrappers.True(),
							Accounts: []*canary.GoogleAccount{
								{
									Name:    "my-google-account",
									Project: "my-google-project",
									SupportedTypes: []canary.SupportedType{
										canary.SupportedType_METRICS_STORE,
									},
								},
							},
							GcsEnabled:                wrappers.False(),
							StackdriverEnabled:        wrappers.True(),
							MetadataCachingIntervalMS: 1000,
						},
					},
				},
			},
			&config.Kayenta{
				Kayenta: &config.Kayenta_ServiceIntegrations{
					Google: &config.Kayenta_ServiceIntegrations_Google{
						Enabled: wrappers.True(),
						Accounts: []*canary.GoogleAccount{
							{
								Name:    "my-google-account",
								Project: "my-google-project",
								SupportedTypes: []canary.SupportedType{
									canary.SupportedType_METRICS_STORE,
								},
							},
						},
					},
					Stackdriver: &canary.Stackdriver{
						Enabled:                   wrappers.True(),
						MetadataCachingIntervalMS: 1000,
					},
					Gcs: &canary.Gcs{
						Enabled: wrappers.False(),
					},
				},
			},
		},
		{
			"GCS enabled",
			&config.Hal{
				Canary: &canary.Canary{
					Enabled: wrappers.True(),
					ServiceIntegrations: &canary.Canary_ServiceIntegrations{
						Google: &canary.Google{
							Enabled: wrappers.True(),
							Accounts: []*canary.GoogleAccount{
								{
									Name:    "my-google-account",
									Project: "my-google-project",
									SupportedTypes: []canary.SupportedType{
										canary.SupportedType_CONFIGURATION_STORE,
										canary.SupportedType_OBJECT_STORE,
									},
								},
							},
							GcsEnabled:         wrappers.True(),
							StackdriverEnabled: wrappers.False(),
						},
					},
				},
			},
			&config.Kayenta{
				Kayenta: &config.Kayenta_ServiceIntegrations{
					Google: &config.Kayenta_ServiceIntegrations_Google{
						Enabled: wrappers.True(),
						Accounts: []*canary.GoogleAccount{
							{
								Name:    "my-google-account",
								Project: "my-google-project",
								SupportedTypes: []canary.SupportedType{
									canary.SupportedType_CONFIGURATION_STORE,
									canary.SupportedType_OBJECT_STORE,
								},
							},
						},
					},
					Stackdriver: &canary.Stackdriver{
						Enabled: wrappers.False(),
					},
					Gcs: &canary.Gcs{
						Enabled: wrappers.True(),
					},
				},
			},
		},
		{
			"Prometheus enabled",
			&config.Hal{
				Canary: &canary.Canary{
					Enabled: wrappers.True(),
					ServiceIntegrations: &canary.Canary_ServiceIntegrations{
						Prometheus: &canary.Prometheus{
							Enabled: wrappers.True(),
							Accounts: []*canary.PrometheusAccount{
								{
									Name: "my-prometheus-account",
									Endpoint: &canary.PrometheusAccount_Endpoint{
										BaseUrl: "https://my-prometheus-server",
									},
									UsernamePasswordFile: "/var/secrets/prometheus",
								},
							},
							MetadataCachingIntervalMS: 1000,
						},
					},
				},
			},
			&config.Kayenta{
				Kayenta: &config.Kayenta_ServiceIntegrations{
					Prometheus: &canary.Prometheus{
						Enabled: wrappers.True(),
						Accounts: []*canary.PrometheusAccount{
							{
								Name: "my-prometheus-account",
								Endpoint: &canary.PrometheusAccount_Endpoint{
									BaseUrl: "https://my-prometheus-server",
								},
								UsernamePasswordFile: "/var/secrets/prometheus",
							},
						},
						MetadataCachingIntervalMS: 1000,
					},
				},
			},
		},
		{
			"Datadog enabled",
			&config.Hal{
				Canary: &canary.Canary{
					Enabled: wrappers.True(),
					ServiceIntegrations: &canary.Canary_ServiceIntegrations{
						Datadog: &canary.Datadog{
							Enabled: wrappers.True(),
							Accounts: []*canary.DatadogAccount{
								{
									Name: "my-datadog-account",
									Endpoint: &canary.DatadogAccount_Endpoint{
										BaseUrl: "https://my-datadog-server",
									},
									ApiKey:         "my-api-key",
									ApplicationKey: "my-app-key",
								},
							},
						},
					},
				},
			},
			&config.Kayenta{
				Kayenta: &config.Kayenta_ServiceIntegrations{
					Datadog: &canary.Datadog{
						Enabled: wrappers.True(),
						Accounts: []*canary.DatadogAccount{
							{
								Name: "my-datadog-account",
								Endpoint: &canary.DatadogAccount_Endpoint{
									BaseUrl: "https://my-datadog-server",
								},
								ApiKey:         "my-api-key",
								ApplicationKey: "my-app-key",
							},
						},
					},
				},
			},
		},
		{
			"S3 enabled",
			&config.Hal{
				Canary: &canary.Canary{
					Enabled: wrappers.True(),
					ServiceIntegrations: &canary.Canary_ServiceIntegrations{
						Aws: &canary.Aws{
							Enabled: wrappers.True(),
							Accounts: []*canary.AwsAccount{
								{
									Name: "my-aws-account",
									SupportedTypes: []canary.SupportedType{
										canary.SupportedType_CONFIGURATION_STORE,
										canary.SupportedType_OBJECT_STORE,
									},
								},
							},
							S3Enabled: wrappers.True(),
						},
					},
				},
			},
			&config.Kayenta{
				Kayenta: &config.Kayenta_ServiceIntegrations{
					Aws: &config.Kayenta_ServiceIntegrations_Aws{
						Enabled: wrappers.True(),
						Accounts: []*canary.AwsAccount{
							{
								Name: "my-aws-account",
								SupportedTypes: []canary.SupportedType{
									canary.SupportedType_CONFIGURATION_STORE,
									canary.SupportedType_OBJECT_STORE,
								},
							},
						},
					},
					S3: &canary.S3{
						Enabled: wrappers.True(),
					},
				},
			},
		},
		{
			"SignalFx enabled",
			&config.Hal{
				Canary: &canary.Canary{
					Enabled: wrappers.True(),
					ServiceIntegrations: &canary.Canary_ServiceIntegrations{
						Signalfx: &canary.SignalFx{
							Enabled: wrappers.True(),
							Accounts: []*canary.SignalFxAccount{
								{
									Name: "my-signalfx-account",
									Endpoint: &canary.SignalFxAccount_Endpoint{
										BaseUrl: "https://my-signalfx-server",
									},
									AccessToken: "my-token",
								},
							},
						},
					},
				},
			},
			&config.Kayenta{
				Kayenta: &config.Kayenta_ServiceIntegrations{
					Signalfx: &canary.SignalFx{
						Enabled: wrappers.True(),
						Accounts: []*canary.SignalFxAccount{
							{
								Name: "my-signalfx-account",
								Endpoint: &canary.SignalFxAccount_Endpoint{
									BaseUrl: "https://my-signalfx-server",
								},
								AccessToken: "my-token",
							},
						},
					},
				},
			},
		},
		{
			"New Relic enabled",
			&config.Hal{
				Canary: &canary.Canary{
					Enabled: wrappers.True(),
					ServiceIntegrations: &canary.Canary_ServiceIntegrations{
						Newrelic: &canary.NewRelic{
							Enabled: wrappers.True(),
							Accounts: []*canary.NewRelicAccount{
								{
									Name: "my-newrelic-account",
									Endpoint: &canary.NewRelicAccount_Endpoint{
										BaseUrl: "https://my-signalfx-server",
									},
								},
							},
						},
					},
				},
			},
			&config.Kayenta{
				Kayenta: &config.Kayenta_ServiceIntegrations{
					Newrelic: &canary.NewRelic{
						Enabled: wrappers.True(),
						Accounts: []*canary.NewRelicAccount{
							{
								Name: "my-newrelic-account",
								Endpoint: &canary.NewRelicAccount_Endpoint{
									BaseUrl: "https://my-signalfx-server",
								},
							},
						},
					},
				},
			},
		},
	},
}

func TestHalToKayenta(t *testing.T) {
	runConfigTest(t, kayentaTests)
}
