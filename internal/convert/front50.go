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
	"github.com/spinnaker/kleat/api/client/config"
)

// HalToFront50 generates the front50 config for the supplied config.Hal h.
func HalToFront50(h *config.Hal) *config.Front50 {
	return &config.Front50{
		Spinnaker: &config.Front50_Spinnaker{
			Gcs:      h.GetPersistentStorage().GetGcs(),
			Azs:      h.GetPersistentStorage().GetAzs(),
			Oracle:   h.GetPersistentStorage().GetOracle(),
			S3:       h.GetPersistentStorage().GetS3(),
			Delivery: getDelivery(h),
		},
	}
}

func getDelivery(h *config.Hal) *config.Front50_Delivery {
	if !h.GetManagedDelivery().GetEnabled().GetValue() {
		return nil
	}

	return &config.Front50_Delivery{
		Enabled: h.GetManagedDelivery().GetEnabled(),
	}
}
