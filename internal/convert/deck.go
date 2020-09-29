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

package convert

import (
	"fmt"

	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/spinnaker/kleat/api/client/cloudprovider"
	"github.com/spinnaker/kleat/api/client/config"
)

// HalToDeck generates the deck config for the supplied config.Hal h.
func HalToDeck(h *config.Hal) *config.Deck {
	gateURL := getGateURL(h)
	return &config.Deck{
		GateUrl:         gateURL,
		AuthEnabled:     h.GetSecurity().GetAuthn().GetEnabled(),
		AuthEndpoint:    gateURL + "/auth/user",
		BakeryDetailUrl: gateURL + "/bakery/logs/{{context.region}}/{{context.status.resourceId}}",
		Canary:          getDeckCanaryConfig(h),
		DefaultTimeZone: h.GetTimezone(),
		Feature:         getDeckFeaturesConfig(h),
		Notifications:   getDeckNotificationsConfig(h),
		Providers:       getDeckProvidersConfig(h),
		Version:         h.GetVersion(),
		ManagedDelivery: getManagedDelivery(h),
	}
}

func getGateURL(h *config.Hal) string {
	if override := h.GetSecurity().GetApiSecurity().GetOverrideBaseUrl(); override != "" {
		return override
	}
	scheme := "http"
	if h.GetSecurity().GetApiSecurity().GetSsl().GetEnabled().GetValue() {
		scheme = "https"
	}
	return fmt.Sprintf("%s://localhost:8084", scheme)
}

func getManagedDelivery(h *config.Hal) *config.Deck_ManagedDelivery {
	if !h.GetManagedDelivery().GetEnabled().GetValue() {
		return nil
	}

	return &config.Deck_ManagedDelivery{
		ManifestBasePath: ".spinnaker",
	}
}

func getDeckCanaryConfig(h *config.Hal) *config.Deck_Canary {
	return &config.Deck_Canary{
		DefaultJudge:       h.GetCanary().GetDefaultJudge(),
		FeatureDisabled:    wrapperspb.Bool(!h.GetCanary().GetEnabled().GetValue()),
		MetricsAccountName: h.GetCanary().GetDefaultMetricsAccount(),
		MetricStore:        h.GetCanary().GetDefaultMetricsStore(),
		ShowAllConfigs:     h.GetCanary().GetShowAllConfigsEnabled(),
		StorageAccountName: h.GetCanary().GetStorageAccountName(),
		TemplatesEnabled:   h.GetCanary().GetTemplatesEnabled(),
	}
}

func getDeckFeaturesConfig(h *config.Hal) *config.Deck_Features {
	return &config.Deck_Features{
		PipelineTemplates:            h.GetFeatures().GetPipelineTemplates(),
		ChaosMonkey:                  h.GetFeatures().GetChaos(),
		FiatEnabled:                  h.GetSecurity().GetAuthz().GetEnabled(),
		ManagedPipelineTemplatesV2UI: h.GetFeatures().GetManagedPipelineTemplatesV2UI(),

		ManagedDelivery:  h.GetManagedDelivery().GetEnabled(),
		ManagedResources: h.GetManagedDelivery().GetEnabled(),
	}
}

func getDeckNotificationsConfig(h *config.Hal) *config.Deck_Notifications {
	return &config.Deck_Notifications{
		Bearychat:    h.GetNotifications().GetBearychat(),
		Email:        h.GetNotifications().GetEmail(),
		GithubStatus: h.GetNotifications().GetGithubStatus(),
		GoogleChat:   h.GetNotifications().GetGooglechat(),
		Pubsub:       h.GetNotifications().GetPubsub(),
		Slack:        h.GetNotifications().GetSlack(),
		Sms:          h.GetNotifications().GetTwilio(),
	}
}

// TODO: find more elegant go-native way to reduce duplicated logic for:
// 1. Finding account config for primaryAccount
// 2. Finding first configured region for a primaryAccount
func getDeckProvidersConfig(h *config.Hal) *config.Deck_Providers {
	providers := &config.Deck_Providers{}

	if h.GetProviders().GetAppengine().GetEnabled().GetValue() {
		providers.Appengine = &config.Deck_Providers_Appengine{
			Defaults: &config.Deck_Providers_Appengine_Defaults{
				Account: h.GetProviders().GetAppengine().GetPrimaryAccount(),
			},
		}
	}

	if h.GetProviders().GetAws().GetEnabled().GetValue() {
		defaultAccountName := h.GetProviders().GetAws().GetPrimaryAccount()
		var defaultAccount *cloudprovider.AwsAccount
		for _, a := range h.GetProviders().GetAws().GetAccounts() {
			if a.GetName() == defaultAccountName {
				defaultAccount = a
			}
		}

		defaultRegion := ""
		if defaultAccount != nil && len(defaultAccount.GetRegions()) > 0 {
			defaultRegion = defaultAccount.GetRegions()[0].GetName()
		}

		providers.Aws = &config.Deck_Providers_Aws{
			Defaults: &config.Deck_Providers_Aws_Defaults{
				Account: defaultAccountName,
				Region:  defaultRegion,
			},
		}
	}

	if h.GetProviders().GetAzure().GetEnabled().GetValue() {
		defaultAccountName := h.GetProviders().GetAzure().GetPrimaryAccount()
		var defaultAccount *cloudprovider.AzureAccount
		for _, a := range h.GetProviders().GetAzure().GetAccounts() {
			if a.GetName() == defaultAccountName {
				defaultAccount = a
			}
		}

		defaultRegion := ""
		if defaultAccount != nil && len(defaultAccount.GetRegions()) > 0 {
			defaultRegion = defaultAccount.GetRegions()[0]
		}

		providers.Azure = &config.Deck_Providers_Azure{
			Defaults: &config.Deck_Providers_Azure_Defaults{
				Account: defaultAccountName,
				Region:  defaultRegion,
			},
		}
	}

	if h.GetProviders().GetCloudfoundry().GetEnabled().GetValue() {
		providers.Cloudfoundry = &config.Deck_Providers_Cloudfoundry{
			Defaults: &config.Deck_Providers_Cloudfoundry_Defaults{
				Account: h.GetProviders().GetCloudfoundry().GetPrimaryAccount(),
			},
		}
	}

	if h.GetProviders().GetDcos().GetEnabled().GetValue() {
		providers.Dcos = &config.Deck_Providers_Dcos{
			Defaults: &config.Deck_Providers_Dcos_Defaults{
				Account: h.GetProviders().GetDcos().GetPrimaryAccount(),
			},
		}
	}

	if h.GetProviders().GetEcs().GetEnabled().GetValue() {
		providers.Ecs = &config.Deck_Providers_Ecs{
			Defaults: &config.Deck_Providers_Ecs_Defaults{
				Account: h.GetProviders().GetEcs().GetPrimaryAccount(),
			},
		}
	}

	if h.GetProviders().GetGoogle().GetEnabled().GetValue() {
		defaultAccountName := h.GetProviders().GetGoogle().GetPrimaryAccount()
		var defaultAccount *cloudprovider.GoogleComputeEngineAccount
		for _, a := range h.GetProviders().GetGoogle().GetAccounts() {
			if a.GetName() == defaultAccountName {
				defaultAccount = a
			}
		}

		defaultRegion := ""
		if defaultAccount != nil && len(defaultAccount.GetRegions()) > 0 {
			defaultRegion = defaultAccount.GetRegions()[0]
		}

		providers.Gce = &config.Deck_Providers_Gce{
			Defaults: &config.Deck_Providers_Gce_Defaults{
				Account: defaultAccountName,
				Region:  defaultRegion,
			},
		}
	}

	if h.GetProviders().GetHuaweicloud().GetEnabled().GetValue() {
		defaultAccountName := h.GetProviders().GetHuaweicloud().GetPrimaryAccount()
		var defaultAccount *cloudprovider.HuaweiCloudAccount
		for _, a := range h.GetProviders().GetHuaweicloud().GetAccounts() {
			if a.GetName() == defaultAccountName {
				defaultAccount = a
			}
		}

		defaultRegion := ""
		if defaultAccount != nil && len(defaultAccount.GetRegions()) > 0 {
			defaultRegion = defaultAccount.GetRegions()[0]
		}

		providers.Huaweicloud = &config.Deck_Providers_HuaweiCloud{
			Defaults: &config.Deck_Providers_HuaweiCloud_Defaults{
				Account: defaultAccountName,
				Region:  defaultRegion,
			},
		}
	}

	if h.GetProviders().GetKubernetes().GetEnabled().GetValue() {
		providers.Kubernetes = &config.Deck_Providers_Kubernetes{}
	}

	return providers
}
