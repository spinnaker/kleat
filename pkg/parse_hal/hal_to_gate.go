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

import "github.com/spinnaker/kleat/api/client/config"

func HalToGate(h *config.Hal) *config.Gate {
	return &config.Gate{
		Server: &config.ServerConfig{
			Ssl: h.GetSecurity().GetApiSecurity().GetSsl(),
			// TODO(ezimanyi): Halyard is setting the port and address here by looking
			// them up on the service settings map; the other services just set this
			// in their default config. We can probably do the same for Gate so that
			// kleat doesn't need to involve itself in service discovery. If we do
			// need to set these values here, we'll need to add them to the Server
			// proto and set them here.
		},
		Cors:     &config.Cors{AllowedOriginsPattern: h.GetSecurity().GetApiSecurity().GetCorsAccessPattern()},
		Security: &config.SpringSecurity{Oauth2: h.GetSecurity().GetAuthn().GetOauth2()},
		Saml:     h.GetSecurity().GetAuthn().GetSaml(),
		Ldap:     h.GetSecurity().GetAuthn().GetLdap(),
		X509:     h.GetSecurity().GetAuthn().GetX509(),
		Google:   &config.Gate_GoogleConfig{Iap: h.GetSecurity().GetAuthn().GetIap()},
	}
}
