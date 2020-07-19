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
	"google.golang.org/protobuf/proto"

	"github.com/spinnaker/kleat/api/client/config"
	"github.com/spinnaker/kleat/api/client/security/authn"
)

// HalToGate generates the gate config for the supplied config.Hal h.
func HalToGate(h *config.Hal) *config.Gate {
	return &config.Gate{
		Server:       getServerConfig(h),
		Cors:         getCorsConfig(h),
		Security:     getSpringConfig(h),
		Saml:         h.GetSecurity().GetAuthn().GetSaml(),
		Ldap:         h.GetSecurity().GetAuthn().GetLdap(),
		X509:         h.GetSecurity().GetAuthn().GetX509(),
		Google:       getGoogleConfig(h),
		Services:     getGateServices(h),
		Integrations: getGateIntegrations(h),
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

func getOauth2Config(h *config.Hal) *authn.OAuth2 {
	// TODO(ezimanyi): Consider changing the oauth2 to instead pass through the
	// config anytime oauth2 != nil. This requires changing the logic in gate
	// to properly check the enabled flag.
	// (In that case we can probably just inline and remove this function.)
	if oauth2 := h.GetSecurity().GetAuthn().GetOauth2(); oauth2.GetEnabled().GetValue() {
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

func getGateServices(h *config.Hal) *config.Gate_Services {
	result := &config.Gate_Services{
		Deck:    getDeckConfig(h),
		Kayenta: getKayentaConfig(h),
	}
	if proto.Equal(result, &config.Gate_Services{}) {
		return nil
	}
	return result
}

func getKayentaConfig(h *config.Hal) *config.ServiceSettings {
	if h.GetCanary().GetEnabled().GetValue() {
		return &config.ServiceSettings{
			Enabled: h.GetCanary().GetEnabled(),
		}
	}
	return nil
}

func getDeckConfig(h *config.Hal) *config.ServiceSettings {
	if h.GetSecurity().GetUiSecurity().GetOverrideBaseUrl() != "" {
		return &config.ServiceSettings{
			BaseUrl: h.GetSecurity().GetUiSecurity().GetOverrideBaseUrl(),
		}
	}
	return nil
}

func getGateIntegrations(h *config.Hal) *config.Gate_Integrations {
	if h.GetFeatures().GetGremlin().GetValue() {
		return &config.Gate_Integrations{
			Gremlin: &config.Gate_Integrations_Gremlin{
				Enabled: h.GetFeatures().GetGremlin(),
			},
		}
	}
	return nil
}
