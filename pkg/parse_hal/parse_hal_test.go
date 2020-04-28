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
package parse_hal

import (
	"reflect"
	"testing"

	"github.com/spinnaker/kleat/api/client"
	"github.com/spinnaker/kleat/api/client/artifact"
	"github.com/spinnaker/kleat/api/client/ci"
	"github.com/spinnaker/kleat/api/client/cloudprovider"
	"github.com/spinnaker/kleat/api/client/config"
	"github.com/spinnaker/kleat/api/client/notification"
	"github.com/spinnaker/kleat/api/client/pubsub"
)

var halToClouddriverTests = []struct {
	n     string
	h     *config.Hal
	wantC *config.Clouddriver
}{
	{
		"Empty hal config",
		&config.Hal{},
		&config.Clouddriver{},
	},
	{
		"Empty providers",
		&config.Hal{
			Providers: &cloudprovider.Providers{},
		},
		&config.Clouddriver{},
	},
	{
		"Empty Kubernetes provider",
		&config.Hal{
			Providers: &cloudprovider.Providers{
				Kubernetes: &cloudprovider.Kubernetes{},
			},
		},
		&config.Clouddriver{
			Kubernetes: &cloudprovider.Kubernetes{},
		},
	},
	{
		"Kubernetes account",
		&config.Hal{
			Providers: &cloudprovider.Providers{
				Kubernetes: &cloudprovider.Kubernetes{
					Enabled: true,
					Accounts: []*cloudprovider.KubernetesAccount{
						{
							Name:           "my-account",
							Kinds:          []string{"deployment"},
							OmitNamespaces: []string{"kube-system"},
						},
					},
					PrimaryAccount: "my-account",
				},
			},
		},
		&config.Clouddriver{
			Kubernetes: &cloudprovider.Kubernetes{
				Enabled: true,
				Accounts: []*cloudprovider.KubernetesAccount{
					{
						Name:           "my-account",
						Kinds:          []string{"deployment"},
						OmitNamespaces: []string{"kube-system"},
					},
				},
				PrimaryAccount: "my-account",
			},
		},
	},
	{
		"Empty artifacts",
		&config.Hal{
			Artifacts: &artifact.Artifacts{},
		},
		&config.Clouddriver{
			Artifacts: &artifact.Artifacts{},
		},
	},
	{
		"Empty GCS artifacts",
		&config.Hal{
			Artifacts: &artifact.Artifacts{
				Gcs: &artifact.Gcs{},
			},
		},
		&config.Clouddriver{
			Artifacts: &artifact.Artifacts{
				Gcs: &artifact.Gcs{},
			},
		},
	},
	{
		"GCS artifact account",
		&config.Hal{
			Artifacts: &artifact.Artifacts{
				Gcs: &artifact.Gcs{
					Enabled: true,
					Accounts: []*artifact.GcsAccount{
						{
							Name:     "my-account",
							JsonPath: "/var/secrets/my-key.json",
						},
					},
				},
			},
		},
		&config.Clouddriver{
			Artifacts: &artifact.Artifacts{
				Gcs: &artifact.Gcs{
					Enabled: true,
					Accounts: []*artifact.GcsAccount{
						{
							Name:     "my-account",
							JsonPath: "/var/secrets/my-key.json",
						},
					},
				},
			},
		},
	},
}

func TestHalToClouddriver(t *testing.T) {
	for _, tt := range halToClouddriverTests {
		t.Run(tt.n, func(t *testing.T) {
			gotC := HalToClouddriver(tt.h)
			if !reflect.DeepEqual(gotC, tt.wantC) {
				t.Errorf("Expected hal config to generate %v, got %v", tt.wantC, gotC)
			}
		})
	}
}

var halToEchoTests = []struct {
	n     string
	h     *config.Hal
	wantE *config.Echo
}{
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
			Stats: &client.Stats{Enabled: true},
		},
		&config.Echo{
			Stats: &client.Stats{Enabled: true},
		},
	},
	{
		"Stats disabled",
		&config.Hal{
			Stats: &client.Stats{Enabled: false},
		},
		&config.Echo{
			Stats: &client.Stats{Enabled: false},
		},
	},
}

func TestHalToEcho(t *testing.T) {
	for _, tt := range halToEchoTests {
		t.Run(tt.n, func(t *testing.T) {
			gotC := HalToEcho(tt.h)
			if !reflect.DeepEqual(gotC, tt.wantE) {
				t.Errorf("Expected hal config to generate %v, got %v", tt.wantE, gotC)
			}
		})
	}
}
