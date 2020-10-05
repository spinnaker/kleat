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
	"github.com/spinnaker/kleat/api/client/storage"
	"github.com/spinnaker/kleat/internal/convert"
)

var keelTests = configTest{
	generator: func(h *config.Hal) proto.Message { return convert.HalToKeel(h) },
	tests: []testCase{
		{
			"Empty hal config",
			&config.Hal{},
			&config.Keel{},
		},
		{
			"Enabling ManagedDelivery",
			&config.Hal{
				ManagedDelivery: &config.ManagedDelivery{
					Enabled: wrapperspb.Bool(true),
				},
			},
			&config.Keel{},
		},
		{
			"Enabling Eureka",
			&config.Hal{
				ManagedDelivery: &config.ManagedDelivery{
					Enabled: wrapperspb.Bool(true),
					Eureka: &config.Eureka{
						Enabled: wrapperspb.Bool(true),
					},
				},
			},
			&config.Keel{
				Eureka: &config.Eureka{
					Enabled: wrapperspb.Bool(true),
				},
			},
		},
		{
			"Enabling Keel Config",
			&config.Hal{
				ManagedDelivery: &config.ManagedDelivery{
					Enabled: wrapperspb.Bool(true),
					Keel: &config.KeelConfig{
						ResourceCheck: &config.KeelConfig_ResourceCheck{
							MinAgeDuration: "2s",
						},
						ArtifactRefresh: &config.KeelConfig_ArtifactRefresh{
							Frequency: "PT1M",
						},
						Constraints: &config.KeelConfig_Constraints{
							ManualJudgement: &config.KeelConfig_Constraints_ManualJudgement{
								InteractiveNotifactions: &config.KeelConfig_Constraints_InteractiveNotifications{
									Enabled: wrapperspb.Bool(true),
								},
							},
						},
						Plugins: &config.KeelConfig_Plugins{
							Ec2: &config.KeelConfig_Plugins_EC2{
								Enabled: wrapperspb.Bool(true),
							},
							DeliveryConfig: &config.KeelConfig_Plugins_DeliveryConfig{
								Enabled: wrapperspb.Bool(true),
							},
							K8S: &config.KeelConfig_Plugins_Kubernetes{
								Enabled: wrapperspb.Bool(true),
							},
							Titus: &config.KeelConfig_Plugins_Titus{
								Enabled: wrapperspb.Bool(true),
							},
						},
					},
				},
			},
			&config.Keel{
				Keel: &config.KeelConfig{
					ResourceCheck: &config.KeelConfig_ResourceCheck{
						MinAgeDuration: "2s",
					},
					ArtifactRefresh: &config.KeelConfig_ArtifactRefresh{
						Frequency: "PT1M",
					},
					Constraints: &config.KeelConfig_Constraints{
						ManualJudgement: &config.KeelConfig_Constraints_ManualJudgement{
							InteractiveNotifactions: &config.KeelConfig_Constraints_InteractiveNotifications{
								Enabled: wrapperspb.Bool(true),
							},
						},
					},
					Plugins: &config.KeelConfig_Plugins{
						Ec2: &config.KeelConfig_Plugins_EC2{
							Enabled: wrapperspb.Bool(true),
						},
						DeliveryConfig: &config.KeelConfig_Plugins_DeliveryConfig{
							Enabled: wrapperspb.Bool(true),
						},
						K8S: &config.KeelConfig_Plugins_Kubernetes{
							Enabled: wrapperspb.Bool(true),
						},
						Titus: &config.KeelConfig_Plugins_Titus{
							Enabled: wrapperspb.Bool(true),
						},
					},
				},
			},
		},
		{
			"SQL Configuration",
			&config.Hal{
				PersistentStorage: &storage.PersistentStorage{
					Sql: &storage.SQL{
						Enabled: wrapperspb.Bool(true),
						ConnectionPools: &storage.SQL_ConnectionPools{
							Default: &storage.SQL_ConnectionPool{
								User:     "username",
								Password: "password",
								JdbcUrl:  "jdbc:mysql://somehost:3306/db",
							},
						},
					},
				},
				ManagedDelivery: &config.ManagedDelivery{
					Enabled: wrapperspb.Bool(true),
				},
			},
			&config.Keel{
				Sql: &storage.SQL{
					Enabled: wrapperspb.Bool(true),
					ConnectionPools: &storage.SQL_ConnectionPools{
						Default: &storage.SQL_ConnectionPool{
							User:     "username",
							Password: "password",
							JdbcUrl:  "jdbc:mysql://somehost:3306/db",
						},
					},
				},
			},
		},
	},
}

func TestHalToKeel(t *testing.T) {
	runConfigTest(t, keelTests)
}
