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

	"github.com/spinnaker/kleat/api/client/config"
	"github.com/spinnaker/kleat/api/client/metricstores"
	"github.com/spinnaker/kleat/internal/convert"
)

var monitoringTests = configTest{
	generator: func(h *config.Hal) proto.Message { return convert.HalToMonitoring(h) },
	tests: []testCase{
		{
			"Empty hal config",
			&config.Hal{},
			&config.Monitoring{
				Monitor: &config.Monitoring_Monitor{},
			},
		},
		{
			"Datadog enabled",
			&config.Hal{
				MetricStores: &metricstores.MetricStores{
					Datadog: &metricstores.Datadog{
						Enabled: wrapperspb.Bool(true),
						ApiKey:  "my-api-key",
						AppKey:  "my-app-keh",
						Tags:    []string{"a:b", "c:d"},
					},
				},
			},
			&config.Monitoring{
				Datadog: &metricstores.Datadog{
					Enabled: wrapperspb.Bool(true),
					ApiKey:  "my-api-key",
					AppKey:  "my-app-keh",
					Tags:    []string{"a:b", "c:d"},
				},
				Monitor: &config.Monitoring_Monitor{
					MetricStore: []config.MetricStoreType{
						config.MetricStoreType_datadog,
					},
				},
			},
		},
		{
			"New Relic enabled",
			&config.Hal{
				MetricStores: &metricstores.MetricStores{
					Newrelic: &metricstores.Newrelic{
						Enabled:   wrapperspb.Bool(true),
						InsertKey: "my-key",
						Host:      "https://my-newrelic-host",
						Tags:      []string{"a:b", "c:d"},
					},
				},
			},
			&config.Monitoring{
				Newrelic: &metricstores.Newrelic{
					Enabled:   wrapperspb.Bool(true),
					InsertKey: "my-key",
					Host:      "https://my-newrelic-host",
					Tags:      []string{"a:b", "c:d"},
				},
				Monitor: &config.Monitoring_Monitor{
					MetricStore: []config.MetricStoreType{
						config.MetricStoreType_newrelic,
					},
				},
			},
		},
		{
			"Prometheus enabled",
			&config.Hal{
				MetricStores: &metricstores.MetricStores{
					Prometheus: &metricstores.Prometheus{
						Enabled:     wrapperspb.Bool(true),
						PushGateway: "https://my-prometheus-gateway",
					},
				},
			},
			&config.Monitoring{
				Prometheus: &metricstores.Prometheus{
					Enabled:     wrapperspb.Bool(true),
					PushGateway: "https://my-prometheus-gateway",
				},
				Monitor: &config.Monitoring_Monitor{
					MetricStore: []config.MetricStoreType{
						config.MetricStoreType_prometheus,
					},
				},
			},
		},
		{
			"Stackdriver enabled",
			&config.Hal{
				MetricStores: &metricstores.MetricStores{
					Stackdriver: &metricstores.Stackdriver{
						Enabled:         wrapperspb.Bool(true),
						CredentialsPath: "/var/secrets/google",
						Project:         "my-project",
						Zone:            "my-region",
					},
				},
			},
			&config.Monitoring{
				Stackdriver: &metricstores.Stackdriver{
					Enabled:         wrapperspb.Bool(true),
					CredentialsPath: "/var/secrets/google",
					Project:         "my-project",
					Zone:            "my-region",
				},
				Monitor: &config.Monitoring_Monitor{
					MetricStore: []config.MetricStoreType{
						config.MetricStoreType_stackdriver,
					},
				},
			},
		},
		{
			"Period configured",
			&config.Hal{
				MetricStores: &metricstores.MetricStores{
					Period: 360,
				},
			},
			&config.Monitoring{
				Monitor: &config.Monitoring_Monitor{
					Period: 360,
				},
			},
		},
	},
}

func TestHalToMonitoring(t *testing.T) {
	runConfigTest(t, monitoringTests)
}
