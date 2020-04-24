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
)

func TestEmptyHalConfigToClouddriver(t *testing.T) {
	h := &client.HalConfig{}
	gotC, err := HalToClouddriver(h)
	if err != nil {
		t.Errorf("Error writing clouddriver config %s", err)
	}
	wantC := &client.ClouddriverConfig{}
	if !reflect.DeepEqual(gotC, wantC) {
		t.Errorf("Expected empty hal config to generate empty clouddriver config, got %v", gotC)
	}
}

func TestEmptyProvidersToClouddriver(t *testing.T) {
	h := &client.HalConfig{
		Providers: &client.HalConfig_Providers{},
	}
	gotC, err := HalToClouddriver(h)
	if err != nil {
		t.Errorf("Error writing clouddriver config %s", err)
	}
	wantC := &client.ClouddriverConfig{}
	if !reflect.DeepEqual(gotC, wantC) {
		t.Errorf("Expected empty hal config to generate empty clouddriver config, got %v", gotC)
	}
}

func TestEmptyKubernetesProviderToClouddriverConfig(t *testing.T) {
	h := &client.HalConfig{
		Providers: &client.HalConfig_Providers{
			Kubernetes: &client.KubernetesProvider{},
		},
	}
	gotC, err := HalToClouddriver(h)
	if err != nil {
		t.Errorf("Error writing clouddriver config %s", err)
	}
	wantC := &client.ClouddriverConfig{
		Kubernetes: &client.KubernetesProvider{},
	}
	if !reflect.DeepEqual(gotC, wantC) {
		t.Errorf("Expected empty Kubernetes config in hal config to pass through to clouddriver config, got %+v", gotC)
	}
}

func TestKubernetesAccountToClouddriver(t *testing.T) {
	k := &client.KubernetesProvider{
		Enabled: true,
		Accounts: []*client.KubernetesAccount{
			{
				Name:           "my-account",
				Kinds:          []string{"deployment"},
				OmitNamespaces: []string{"kube-system"},
			},
		},
		PrimaryAccount: "my-account",
	}
	h := &client.HalConfig{
		Providers: &client.HalConfig_Providers{
			Kubernetes: k,
		},
	}
	gotC, err := HalToClouddriver(h)
	if err != nil {
		t.Errorf("Error writing clouddriver config %s", err)
	}
	wantC := &client.ClouddriverConfig{
		Kubernetes: k,
	}
	if !reflect.DeepEqual(gotC, wantC) {
		t.Errorf("Expected kubernetes account to be in clouddriver config, got %+v", gotC)
	}
}

func TestEmptyArtifactsToHalConfig(t *testing.T) {
	h := &client.HalConfig{
		Artifacts: &client.ArtifactProviders{},
	}
	gotC, err := HalToClouddriver(h)
	if err != nil {
		t.Errorf("Error writing clouddriver config %s", err)
	}
	wantC := &client.ClouddriverConfig{
		Artifacts: &client.ArtifactProviders{},
	}
	if !reflect.DeepEqual(gotC, wantC) {
		t.Errorf("Expected empty artifact providers to be passed to clouddriver config, got %+v", gotC)
	}
}

func TestEmptyGcsArtifactConfigToHalConfig(t *testing.T) {
	a := &client.ArtifactProviders{
		Gcs: &client.GcsArtifactProvider{},
	}
	h := &client.HalConfig{
		Artifacts: a,
	}
	gotC, err := HalToClouddriver(h)
	if err != nil {
		t.Errorf("Error writing clouddriver config %s", err)
	}
	wantC := &client.ClouddriverConfig{
		Artifacts: a,
	}
	if !reflect.DeepEqual(gotC, wantC) {
		t.Errorf("Expected empty GCS artifact config to be passed to clouddriver config, got %+v", gotC)
	}
}

func TestGcsArtifactAccountToHalConfig(t *testing.T) {
	a := &client.ArtifactProviders{
		Gcs: &client.GcsArtifactProvider{
			Enabled: true,
			Accounts: []*client.GcsArtifactAccount{
				{
					Name:     "my-account",
					JsonPath: "/var/secrets/my-key.json",
				},
			},
		},
	}
	h := &client.HalConfig{
		Artifacts: a,
	}
	gotC, err := HalToClouddriver(h)
	if err != nil {
		t.Errorf("Error writing clouddriver config %s", err)
	}
	wantC := &client.ClouddriverConfig{
		Artifacts: a,
	}
	if !reflect.DeepEqual(gotC, wantC) {
		t.Errorf("Expected GCS artifact config to be passed to clouddriver config, got %+v", gotC)
	}
}

func TestEmptyHalConfigToEcho(t *testing.T) {
	h := &client.HalConfig{}
	gotE, err := HalToEcho(h)
	if err != nil {
		t.Errorf("Error writing echo config %s", err)
	}
	wantE := &client.EchoConfig{}
	if !reflect.DeepEqual(gotE, wantE) {
		t.Errorf("Expected empty hal config to generate empty echo config, got %v", gotE)
	}
}

func TestEmptyNotificationsToEchoConfig(t *testing.T) {
	h := &client.HalConfig{
		Notifications: &client.HalConfig_Notifications{},
	}
	gotE, err := HalToEcho(h)
	if err != nil {
		t.Errorf("Error writing echo config %s", err)
	}
	wantE := &client.EchoConfig{}
	if !reflect.DeepEqual(gotE, wantE) {
		t.Errorf("Expected empty hal config to generate empty echo config, got %v", gotE)
	}
}

func TestSlackNotificationToEchoConfig(t *testing.T) {
	slack := &client.SlackNotification{
		Enabled: true,
		BotName: "my-bot",
		Token:   "my-token",
		BaseUrl: "https://slack.test/",
	}
	h := &client.HalConfig{
		Notifications: &client.HalConfig_Notifications{
			Slack: slack,
		},
	}
	gotE, err := HalToEcho(h)
	if err != nil {
		t.Errorf("Error writing echo config %s", err)
	}
	wantE := &client.EchoConfig{
		Slack: slack,
	}
	if !reflect.DeepEqual(gotE, wantE) {
		t.Errorf("Expected slack notifications to be passed through to echo config, got %v", gotE)
	}
}

func TestEmptyPubsubsToEchoConfig(t *testing.T) {
	h := &client.HalConfig{
		Pubsub: &client.PubsubProviders{},
	}
	gotE, err := HalToEcho(h)
	if err != nil {
		t.Errorf("Error writing echo config %s", err)
	}
	wantE := &client.EchoConfig{
		Pubsub: &client.PubsubProviders{},
	}
	if !reflect.DeepEqual(gotE, wantE) {
		t.Errorf("Expected empty pubsubs to be passed through to echo config, got %v", gotE)
	}
}

func TestEmptyGooglePubsubToEchoConfig(t *testing.T) {
	pubsub := &client.PubsubProviders{
		Google: &client.GooglePubsub{},
	}
	h := &client.HalConfig{
		Pubsub: pubsub,
	}
	gotE, err := HalToEcho(h)
	if err != nil {
		t.Errorf("Error writing echo config %s", err)
	}
	wantE := &client.EchoConfig{
		Pubsub: pubsub,
	}
	if !reflect.DeepEqual(gotE, wantE) {
		t.Errorf("Expected empty Google pubsub config to be passed through to echo config, got %v", gotE)
	}
}

func TestGooglePubsubToEchoConfig(t *testing.T) {
	pubsub := &client.PubsubProviders{
		Google: &client.GooglePubsub{
			Subscriptions: []*client.GooglePubsubSubscriber{
				{
					Name:             "my-account",
					Project:          "my-project",
					SubscriptionName: "my-subscription",
					JsonPath:         "/var/secrets/my-account.json",
					MessageFormat:    "GCS",
				},
			},
			Publishers: []*client.GooglePubsubPublisher{
				{
					Name:      "my-account",
					Project:   "my-project",
					TopicName: "my-topic",
				},
			},
		},
	}
	h := &client.HalConfig{
		Pubsub: pubsub,
	}
	gotE, err := HalToEcho(h)
	if err != nil {
		t.Errorf("Error writing echo config %s", err)
	}
	wantE := &client.EchoConfig{
		Pubsub: pubsub,
	}
	if !reflect.DeepEqual(gotE, wantE) {
		t.Errorf("Expected Google pubsub config to be passed through to echo config, got %v", gotE)
	}
}
