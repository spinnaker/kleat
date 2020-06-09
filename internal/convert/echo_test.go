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

	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/spinnaker/kleat/api/client"
	"github.com/spinnaker/kleat/api/client/ci"
	"github.com/spinnaker/kleat/api/client/config"
	"github.com/spinnaker/kleat/api/client/notification"
	"github.com/spinnaker/kleat/api/client/pubsub"
	"github.com/spinnaker/kleat/internal/convert"
	"google.golang.org/protobuf/proto"
)

var echoTests = configTest{
	generator: func(h *config.Hal) proto.Message { return convert.HalToEcho(h) },
	tests: []testCase{
		{
			"Empty hal config",
			&config.Hal{},
			&config.Echo{},
		},
		{
			"Empty notifications",
			&config.Hal{
				Notifications: &notification.Notifications{},
			},
			&config.Echo{},
		},
		{
			"Slack notification",
			&config.Hal{
				Notifications: &notification.Notifications{
					Slack: &notification.Slack{
						Enabled: true,
						BotName: "my-bot",
						Token:   "my-token",
						BaseUrl: "https://slack.test/",
					},
				},
			},
			&config.Echo{
				Slack: &notification.Slack{
					Enabled: true,
					BotName: "my-bot",
					Token:   "my-token",
					BaseUrl: "https://slack.test/",
				},
			},
		},
		{
			"Empty pub/sub",
			&config.Hal{
				Pubsub: &pubsub.Pubsub{},
			},
			&config.Echo{
				Pubsub: &pubsub.Pubsub{},
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
			},
		},
		{
			"Empty CI",
			&config.Hal{
				Ci: &ci.Ci{},
			},
			&config.Echo{},
		},
		{
			"Empty GCB account",
			&config.Hal{
				Ci: &ci.Ci{
					Gcb: &ci.GoogleCloudBuild{},
				},
			},
			&config.Echo{
				Gcb: &ci.GoogleCloudBuild{},
			},
		},
		{
			"GCB account",
			&config.Hal{
				Ci: &ci.Ci{
					Gcb: &ci.GoogleCloudBuild{
						Enabled: true,
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
					Enabled: true,
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
				Stats: &client.Stats{Enabled: &wrappers.BoolValue{Value: true}},
			},
			&config.Echo{
				Stats: &client.Stats{Enabled: &wrappers.BoolValue{Value: true}},
			},
		},
		{
			"Stats disabled",
			&config.Hal{
				Stats: &client.Stats{Enabled: &wrappers.BoolValue{Value: false}},
			},
			&config.Echo{
				Stats: &client.Stats{Enabled: &wrappers.BoolValue{Value: false}},
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
			},
		},
	},
}

func TestHalToEcho(t *testing.T) {
	for _, tt := range echoTests.tests {
		t.Run(tt.name, func(t *testing.T) {
			got := echoTests.generator(tt.hal)
			if !proto.Equal(got, tt.want) {
				t.Errorf("Expected hal config to generate %v, got %v", tt.want, got)
			}
		})
	}
}
