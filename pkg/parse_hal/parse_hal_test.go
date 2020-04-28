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

func TestEmptyHalConfigToClouddriver(t *testing.T) {
	h := &config.Hal{}
	gotC := HalToClouddriver(h)
	wantC := &config.Clouddriver{}
	if !reflect.DeepEqual(gotC, wantC) {
		t.Errorf("Expected empty hal config to generate empty clouddriver config, got %v", gotC)
	}
}

func TestEmptyProvidersToClouddriver(t *testing.T) {
	h := &config.Hal{
		Providers: &config.Hal_Providers{},
	}
	gotC := HalToClouddriver(h)
	wantC := &config.Clouddriver{}
	if !reflect.DeepEqual(gotC, wantC) {
		t.Errorf("Expected empty hal config to generate empty clouddriver config, got %v", gotC)
	}
}

func TestEmptyKubernetesProviderToClouddriverConfig(t *testing.T) {
	h := &config.Hal{
		Providers: &config.Hal_Providers{
			Kubernetes: &cloudprovider.Kubernetes{},
		},
	}
	gotC := HalToClouddriver(h)
	wantC := &config.Clouddriver{
		Kubernetes: &cloudprovider.Kubernetes{},
	}
	if !reflect.DeepEqual(gotC, wantC) {
		t.Errorf("Expected empty Kubernetes config in hal config to pass through to clouddriver config, got %+v", gotC)
	}
}

func TestKubernetesAccountToClouddriver(t *testing.T) {
	k := &cloudprovider.Kubernetes{
		Enabled: true,
		Accounts: []*cloudprovider.KubernetesAccount{
			{
				Name:           "my-account",
				Kinds:          []string{"deployment"},
				OmitNamespaces: []string{"kube-system"},
			},
		},
		PrimaryAccount: "my-account",
	}
	h := &config.Hal{
		Providers: &config.Hal_Providers{
			Kubernetes: k,
		},
	}
	gotC := HalToClouddriver(h)
	wantC := &config.Clouddriver{
		Kubernetes: k,
	}
	if !reflect.DeepEqual(gotC, wantC) {
		t.Errorf("Expected kubernetes account to be in clouddriver config, got %+v", gotC)
	}
}

func TestEmptyArtifactsToHalConfig(t *testing.T) {
	h := &config.Hal{
		Artifacts: &artifact.Artifacts{},
	}
	gotC := HalToClouddriver(h)
	wantC := &config.Clouddriver{
		Artifacts: &artifact.Artifacts{},
	}
	if !reflect.DeepEqual(gotC, wantC) {
		t.Errorf("Expected empty artifact providers to be passed to clouddriver config, got %+v", gotC)
	}
}

func TestEmptyGcsArtifactConfigToHalConfig(t *testing.T) {
	a := &artifact.Artifacts{
		Gcs: &artifact.Gcs{},
	}
	h := &config.Hal{
		Artifacts: a,
	}
	gotC := HalToClouddriver(h)
	wantC := &config.Clouddriver{
		Artifacts: a,
	}
	if !reflect.DeepEqual(gotC, wantC) {
		t.Errorf("Expected empty GCS artifact config to be passed to clouddriver config, got %+v", gotC)
	}
}

func TestGcsArtifactAccountToHalConfig(t *testing.T) {
	a := &artifact.Artifacts{
		Gcs: &artifact.Gcs{
			Enabled: true,
			Accounts: []*artifact.GcsAccount{
				{
					Name:     "my-account",
					JsonPath: "/var/secrets/my-key.json",
				},
			},
		},
	}
	h := &config.Hal{
		Artifacts: a,
	}
	gotC := HalToClouddriver(h)
	wantC := &config.Clouddriver{
		Artifacts: a,
	}
	if !reflect.DeepEqual(gotC, wantC) {
		t.Errorf("Expected GCS artifact config to be passed to clouddriver config, got %+v", gotC)
	}
}

func TestEmptyHalConfigToEcho(t *testing.T) {
	h := &config.Hal{}
	gotE := HalToEcho(h)
	wantE := &config.Echo{}
	if !reflect.DeepEqual(gotE, wantE) {
		t.Errorf("Expected empty hal config to generate empty echo config, got %v", gotE)
	}
}

func TestEmptyNotificationsToEchoConfig(t *testing.T) {
	h := &config.Hal{
		Notifications: &notification.Notifications{},
	}
	gotE := HalToEcho(h)
	wantE := &config.Echo{}
	if !reflect.DeepEqual(gotE, wantE) {
		t.Errorf("Expected empty hal config to generate empty echo config, got %v", gotE)
	}
}

func TestSlackNotificationToEchoConfig(t *testing.T) {
	slack := &notification.Slack{
		Enabled: true,
		BotName: "my-bot",
		Token:   "my-token",
		BaseUrl: "https://slack.test/",
	}
	h := &config.Hal{
		Notifications: &notification.Notifications{
			Slack: slack,
		},
	}
	gotE := HalToEcho(h)
	wantE := &config.Echo{
		Slack: slack,
	}
	if !reflect.DeepEqual(gotE, wantE) {
		t.Errorf("Expected slack notifications to be passed through to echo config, got %v", gotE)
	}
}

func TestEmptyPubsubsToEchoConfig(t *testing.T) {
	h := &config.Hal{
		Pubsub: &pubsub.Pubsub{},
	}
	gotE := HalToEcho(h)
	wantE := &config.Echo{
		Pubsub: &pubsub.Pubsub{},
	}
	if !reflect.DeepEqual(gotE, wantE) {
		t.Errorf("Expected empty pubsubs to be passed through to echo config, got %v", gotE)
	}
}

func TestEmptyGooglePubsubToEchoConfig(t *testing.T) {
	pubsub := &pubsub.Pubsub{
		Google: &pubsub.Google{},
	}
	h := &config.Hal{
		Pubsub: pubsub,
	}
	gotE := HalToEcho(h)
	wantE := &config.Echo{
		Pubsub: pubsub,
	}
	if !reflect.DeepEqual(gotE, wantE) {
		t.Errorf("Expected empty Google pubsub config to be passed through to echo config, got %v", gotE)
	}
}

func TestGooglePubsubToEchoConfig(t *testing.T) {
	pubsub := &pubsub.Pubsub{
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
	}
	h := &config.Hal{
		Pubsub: pubsub,
	}
	gotE := HalToEcho(h)
	wantE := &config.Echo{
		Pubsub: pubsub,
	}
	if !reflect.DeepEqual(gotE, wantE) {
		t.Errorf("Expected Google pubsub config to be passed through to echo config, got %v", gotE)
	}
}

func TestEmptyCiConfigToEcho(t *testing.T) {
	h := &config.Hal{
		Ci: &ci.Ci{},
	}
	gotE := HalToEcho(h)
	wantE := &config.Echo{}
	if !reflect.DeepEqual(gotE, wantE) {
		t.Errorf("Expected empty CI config to lead to empty echo config, got %v", gotE)
	}
}

func TestEmptyGcbConfigAccountToEcho(t *testing.T) {
	gcb := &ci.GoogleCloudBuild{}
	cis := &ci.Ci{
		Gcb: gcb,
	}
	h := &config.Hal{
		Ci: cis,
	}
	gotE := HalToEcho(h)
	wantE := &config.Echo{
		Gcb: gcb,
	}
	if !reflect.DeepEqual(gotE, wantE) {
		t.Errorf("Expected empty Google pubsub config to be passed through to echo config, got %v", gotE)
	}
}

func TestGcbAccountToEcho(t *testing.T) {
	gcb := &ci.GoogleCloudBuild{
		Enabled: true,
		Accounts: []*ci.GoogleCloudBuildAccount{
			{
				Name:             "my-account",
				Project:          "my-project",
				SubscriptionName: "my-subscription",
			},
		},
	}
	cis := &ci.Ci{
		Gcb: gcb,
	}
	h := &config.Hal{
		Ci: cis,
	}
	gotE := HalToEcho(h)
	wantE := &config.Echo{
		Gcb: gcb,
	}
	if !reflect.DeepEqual(gotE, wantE) {
		t.Errorf("Expected Google Cloud Build account to be passed through to echo config, got %v", gotE)
	}
}

func TestEmptyStatsToEcho(t *testing.T) {
	h := &config.Hal{
		Stats: &client.Stats{},
	}
	gotE := HalToEcho(h)
	wantE := &config.Echo{
		Stats: &client.Stats{},
	}
	if !reflect.DeepEqual(gotE, wantE) {
		t.Errorf("Expected empty stats config to be passed through to echo config, got %v", gotE)
	}
}

func TestStatsEnabledToEcho(t *testing.T) {
	stats := &client.Stats{Enabled: true}
	h := &config.Hal{
		Stats: stats,
	}
	gotE := HalToEcho(h)
	wantE := &config.Echo{
		Stats: stats,
	}
	if !reflect.DeepEqual(gotE, wantE) {
		t.Errorf("Expected enabled stats config to be passed through to echo config, got %v", gotE)
	}
}

func TestStatsDisabledToEcho(t *testing.T) {
	stats := &client.Stats{Enabled: false}
	h := &config.Hal{
		Stats: stats,
	}
	gotE := HalToEcho(h)
	wantE := &config.Echo{
		Stats: stats,
	}
	if !reflect.DeepEqual(gotE, wantE) {
		t.Errorf("Expected disabled stats config to be passed through to echo config, got %v", gotE)
	}
}
