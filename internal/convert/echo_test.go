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
	"github.com/spinnaker/kleat/api/client/ci"
	"github.com/spinnaker/kleat/api/client/config"
	"github.com/spinnaker/kleat/api/client/notification"
	"github.com/spinnaker/kleat/api/client/pubsub"
	"github.com/spinnaker/kleat/internal/convert"
)

// deploymentMethod is the deployment method that will be auto-populated by kleat
// regardless of what is present in the input config.
var deploymentMethod = &client.DeploymentMethod{Type: "kleat", Version: "unknown"}

// defaultStats is a wrapper around deploymentMethod, and represents the stats config
// that should be generated when the stats are not explicitly configured in the input
var defaultStats = &config.Echo_Stats{
	DeploymentMethod: deploymentMethod,
}

var echoTests = configTest{
	generator: func(h *config.Hal) proto.Message { return convert.HalToEcho(h) },
	tests: []testCase{
		{
			"Empty hal config",
			&config.Hal{},
			&config.Echo{Stats: defaultStats},
		},
		{
			"Empty notifications",
			&config.Hal{
				Notifications: &notification.Notifications{},
			},
			&config.Echo{Stats: defaultStats},
		},
		{
			"Slack notification",
			&config.Hal{
				Notifications: &notification.Notifications{
					Slack: &notification.Slack{
						Enabled: wrapperspb.Bool(true),
						BotName: "my-bot",
						Token:   "my-token",
						BaseUrl: "https://slack.test/",
					},
				},
			},
			&config.Echo{
				Slack: &notification.Slack{
					Enabled: wrapperspb.Bool(true),
					BotName: "my-bot",
					Token:   "my-token",
					BaseUrl: "https://slack.test/",
				},
				Stats: defaultStats,
			},
		},
		{
			"Empty pub/sub",
			&config.Hal{
				Pubsub: &pubsub.Pubsub{},
			},
			&config.Echo{
				Pubsub: &pubsub.Pubsub{},
				Stats:  defaultStats,
			},
		},
		{
			"Empty Google pub/sub",
			&config.Hal{
				Pubsub: &pubsub.Pubsub{
					Google: &pubsub.Google{},
				},
			},
			&config.Echo{
				Pubsub: &pubsub.Pubsub{
					Google: &pubsub.Google{},
				},
				Stats: defaultStats,
			},
		},
		{
			"Google pub/sub",
			&config.Hal{
				Pubsub: &pubsub.Pubsub{
					Google: &pubsub.Google{
						Subscriptions: []*pubsub.GoogleSubscriber{
							{
								Name:             "my-account",
								Project:          "my-project",
								SubscriptionName: "my-subscription",
								JsonPath:         "/var/secrets/my-account.json",
								MessageFormat:    pubsub.MessageFormat_GCS,
							},
						},
						Publishers: []*pubsub.GooglePublisher{
							{
								Name:      "my-account",
								Project:   "my-project",
								TopicName: "my-topic",
							},
						},
					},
				},
			},
			&config.Echo{
				Pubsub: &pubsub.Pubsub{
					Google: &pubsub.Google{
						Subscriptions: []*pubsub.GoogleSubscriber{
							{
								Name:             "my-account",
								Project:          "my-project",
								SubscriptionName: "my-subscription",
								JsonPath:         "/var/secrets/my-account.json",
								MessageFormat:    pubsub.MessageFormat_GCS,
							},
						},
						Publishers: []*pubsub.GooglePublisher{
							{
								Name:      "my-account",
								Project:   "my-project",
								TopicName: "my-topic",
							},
						},
					},
				},
				Stats: defaultStats,
			},
		},
		{
			"Empty CI",
			&config.Hal{
				Ci: &ci.Ci{},
			},
			&config.Echo{Stats: defaultStats},
		},
		{
			"Empty GCB account",
			&config.Hal{
				Ci: &ci.Ci{
					Gcb: &ci.GoogleCloudBuild{},
				},
			},
			&config.Echo{
				Gcb:   &ci.GoogleCloudBuild{},
				Stats: defaultStats,
			},
		},
		{
			"GCB account",
			&config.Hal{
				Ci: &ci.Ci{
					Gcb: &ci.GoogleCloudBuild{
						Enabled: wrapperspb.Bool(true),
						Accounts: []*ci.GoogleCloudBuildAccount{
							{
								Name:             "my-account",
								Project:          "my-project",
								SubscriptionName: "my-subscription",
							},
						},
					},
				},
			},
			&config.Echo{
				Gcb: &ci.GoogleCloudBuild{
					Enabled: wrapperspb.Bool(true),
					Accounts: []*ci.GoogleCloudBuildAccount{
						{
							Name:             "my-account",
							Project:          "my-project",
							SubscriptionName: "my-subscription",
						},
					},
				},
				Stats: defaultStats,
			},
		},
		{
			"Empty stats",
			&config.Hal{
				Stats: &client.Stats{},
			},
			&config.Echo{Stats: defaultStats},
		},
		{
			"Stats enabled",
			&config.Hal{
				Stats: &client.Stats{Enabled: wrapperspb.Bool(true)},
			},
			&config.Echo{
				Stats: &config.Echo_Stats{
					Enabled:          wrapperspb.Bool(true),
					DeploymentMethod: deploymentMethod,
				},
			},
		},
		{
			"Stats disabled",
			&config.Hal{
				Stats: &client.Stats{Enabled: wrapperspb.Bool(false)},
			},
			&config.Echo{
				Stats: &config.Echo_Stats{
					Enabled:          wrapperspb.Bool(false),
					DeploymentMethod: deploymentMethod,
				},
			},
		},
		{
			"Stats spinnaker version",
			&config.Hal{
				Version: "1.20.4",
			},
			&config.Echo{
				Stats: &config.Echo_Stats{
					DeploymentMethod: deploymentMethod,
					SpinnakerVersion: "1.20.4",
				},
			},
		},
		{
			"Stats instance id",
			&config.Hal{
				Stats: &client.Stats{
					InstanceId: "01EB4ZQP5ZH3FP6FC9ZB9X5ECW",
				},
			},
			&config.Echo{
				Stats: &config.Echo_Stats{
					DeploymentMethod: deploymentMethod,
					InstanceId:       "01EB4ZQP5ZH3FP6FC9ZB9X5ECW",
				},
			},
		},
		{
			"Stats endpoint",
			&config.Hal{
				Stats: &client.Stats{
					Endpoint: "https://stats.mydomain.test/",
				},
			},
			&config.Echo{
				Stats: &config.Echo_Stats{
					DeploymentMethod: deploymentMethod,
					Endpoint:         "https://stats.mydomain.test/",
				},
			},
		},
		{
			"Timezone set",
			&config.Hal{
				Timezone: "America/Chicago",
			},
			&config.Echo{
				Scheduler: &config.Echo_Scheduler{
					Cron: &config.Echo_Scheduler_Cron{
						Timezone: "America/Chicago",
					},
				},
				Stats: defaultStats,
			},
		},
		{
			"Enable ManagedDelivery",
			&config.Hal{
				ManagedDelivery: &config.ManagedDelivery{
					Enabled: wrapperspb.Bool(true),
				},
			},
			&config.Echo{
				Services: &config.Echo_Services{
					Keel: &config.ServiceSettings{
						Enabled: wrapperspb.Bool(true),
					},
				},
				Stats: defaultStats,
			},
		},
	},
}

func TestHalToEcho(t *testing.T) {
	runConfigTest(t, echoTests)
}
