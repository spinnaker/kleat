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

	"github.com/spinnaker/kleat/api/client"
	"github.com/spinnaker/kleat/api/client/ci"
	"github.com/spinnaker/kleat/api/client/config"
	"github.com/spinnaker/kleat/api/client/notification"
	"github.com/spinnaker/kleat/api/client/pubsub"
	"github.com/spinnaker/kleat/internal/convert"
	"github.com/spinnaker/kleat/internal/wrappers"
	"google.golang.org/protobuf/proto"
)

var defaultStats *client.Stats = nil

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
						Enabled: wrappers.True(),
						BotName: "my-bot",
						Token:   "my-token",
						BaseUrl: "https://slack.test/",
					},
				},
			},
			&config.Echo{
				Slack: &notification.Slack{
					Enabled: wrappers.True(),
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
						Enabled: wrappers.True(),
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
					Enabled: wrappers.True(),
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
			&config.Echo{
				Stats: &client.Stats{},
			},
		},
		{
			"Stats enabled",
			&config.Hal{
				Stats: &client.Stats{Enabled: wrappers.True()},
			},
			&config.Echo{
				Stats: &client.Stats{Enabled: wrappers.True()},
			},
		},
		{
			"Stats disabled",
			&config.Hal{
				Stats: &client.Stats{Enabled: wrappers.False()},
			},
			&config.Echo{
				Stats: &client.Stats{Enabled: wrappers.False()},
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
	},
}

func TestHalToEcho(t *testing.T) {
	runConfigTest(t, echoTests)
}
