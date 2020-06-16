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

import "github.com/spinnaker/kleat/api/client/config"

// HalToDeckEnv generates the deck environment config for the supplied config.Hal h.
func HalToDeckEnv(h *config.Hal) *config.DeckEnv {
	if !h.GetSecurity().GetUiSecurity().GetSsl().GetEnabled().GetValue() {
		return &config.DeckEnv{}
	}
	return &config.DeckEnv{
		DeckCert:   h.GetSecurity().GetUiSecurity().GetSsl().GetSslCertificateFile(),
		DeckKey:    h.GetSecurity().GetUiSecurity().GetSsl().GetSslCertificateKeyFile(),
		Passphrase: h.GetSecurity().GetUiSecurity().GetSsl().GetSslCertificatePassphrase(),
	}
}
