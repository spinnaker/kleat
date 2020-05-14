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
	"github.com/spinnaker/kleat/api/client/config"
	"github.com/spinnaker/kleat/api/client/security"
	"google.golang.org/protobuf/proto"
)

func HalToGate(h *config.Hal) *config.Gate {
	return &config.Gate{
		Server:   getServerConfig(h),
		Cors:     getCorsConfig(h),
		Security: getSpringConfig(h),
		Saml:     h.GetSecurity().GetAuthn().GetSaml(),
		Ldap:     h.GetSecurity().GetAuthn().GetLdap(),
		X509:     h.GetSecurity().GetAuthn().GetX509(),
		Google:   getGoogleConfig(h),
	}
}

func getServerConfig(h *config.Hal) *config.ServerConfig {
	if ssl := h.GetSecurity().GetApiSecurity().GetSsl(); ssl != nil {
		return &config.ServerConfig{Ssl: ssl}
	}
	return nil
}

func getSpringConfig(h *config.Hal) *config.SpringSecurity {
	result := &config.SpringSecurity{
		Oauth2: getOauth2Config(h),
		Basic:  h.GetSecurity().GetAuthn().GetBasic(),
	}

	if proto.Equal(result, &config.SpringSecurity{}) {
		return nil
	}
	return result
}

func getOauth2Config(h *config.Hal) *security.OAuth2 {
	// TODO(ezimanyi): Consider changing the oauth2 to instead pass through the
	// config anytime oauth2 != nil. This requires changing the logic in gate
	// to properly check the enabled flag.
	// (In that case we can probably just inline and remove this function.)
	if oauth2 := h.GetSecurity().GetAuthn().GetOauth2(); oauth2.GetEnabled() {
		return oauth2
	}
	return nil
}

func getCorsConfig(h *config.Hal) *config.Cors {
	if p := h.GetSecurity().GetApiSecurity().GetCorsAccessPattern(); p != "" {
		return &config.Cors{AllowedOriginsPattern: p}
	}
	return nil
}

func getGoogleConfig(h *config.Hal) *config.Gate_GoogleConfig {
	if iap := h.GetSecurity().GetAuthn().GetIap(); iap != nil {
		return &config.Gate_GoogleConfig{Iap: iap}
	}
	return nil
}
