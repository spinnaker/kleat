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
	"github.com/spinnaker/kleat/api/client/canary"
	"github.com/spinnaker/kleat/api/client/config"
)

// HalToKayenta generates the kayanta config for the supplied config.Hal h.
func HalToKayenta(h *config.Hal) *config.Kayenta {
	if !h.GetCanary().GetEnabled().GetValue() {
		return &config.Kayenta{}
	}

	return &config.Kayenta{
		Kayenta: &config.Kayenta_ServiceIntegrations{
			Google:      getGoogleCanaryConfig(h),
			Stackdriver: getStackdriverCanaryConfig(h),
			Gcs:         getGcsCanaryConfig(h),
			Prometheus:  h.GetCanary().GetServiceIntegrations().GetPrometheus(),
			Datadog:     h.GetCanary().GetServiceIntegrations().GetDatadog(),
			Aws:         getAwsCanaryConfig(h),
			S3:          getS3CanaryConfig(h),
			Signalfx:    h.GetCanary().GetServiceIntegrations().GetSignalfx(),
			Newrelic:    h.GetCanary().GetServiceIntegrations().GetNewrelic(),
		},
	}
}

func getGoogleCanaryConfig(h *config.Hal) *config.Kayenta_ServiceIntegrations_Google {
	if h.GetCanary().GetServiceIntegrations().GetGoogle() == nil {
		return nil
	}
	return &config.Kayenta_ServiceIntegrations_Google{
		Enabled:  h.GetCanary().GetServiceIntegrations().GetGoogle().GetEnabled(),
		Accounts: h.GetCanary().GetServiceIntegrations().GetGoogle().GetAccounts(),
	}
}

func getStackdriverCanaryConfig(h *config.Hal) *canary.Stackdriver {
	if h.GetCanary().GetServiceIntegrations().GetGoogle() == nil {
		return nil
	}
	return &canary.Stackdriver{
		Enabled:                   h.GetCanary().GetServiceIntegrations().GetGoogle().GetStackdriverEnabled(),
		MetadataCachingIntervalMS: h.GetCanary().GetServiceIntegrations().GetGoogle().GetMetadataCachingIntervalMS(),
	}
}

func getGcsCanaryConfig(h *config.Hal) *canary.Gcs {
	if h.GetCanary().GetServiceIntegrations().GetGoogle() == nil {
		return nil
	}
	return &canary.Gcs{
		Enabled: h.GetCanary().GetServiceIntegrations().GetGoogle().GetGcsEnabled(),
	}
}

func getAwsCanaryConfig(h *config.Hal) *config.Kayenta_ServiceIntegrations_Aws {
	if h.GetCanary().GetServiceIntegrations().GetAws() == nil {
		return nil
	}
	return &config.Kayenta_ServiceIntegrations_Aws{
		Enabled:  h.GetCanary().GetServiceIntegrations().GetAws().GetEnabled(),
		Accounts: h.GetCanary().GetServiceIntegrations().GetAws().GetAccounts(),
	}
}

func getS3CanaryConfig(h *config.Hal) *canary.S3 {
	if h.GetCanary().GetServiceIntegrations().GetAws() == nil {
		return nil
	}
	return &canary.S3{
		Enabled: h.GetCanary().GetServiceIntegrations().GetAws().GetS3Enabled(),
	}
}
