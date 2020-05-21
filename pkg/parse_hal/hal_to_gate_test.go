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
package parse_hal_test

import (
	"testing"

	"github.com/spinnaker/kleat/api/client/canary"

	"github.com/spinnaker/kleat/api/client/config"
	"github.com/spinnaker/kleat/api/client/security"
	"github.com/spinnaker/kleat/api/client/security/authn"
	"github.com/spinnaker/kleat/pkg/parse_hal"
	"google.golang.org/protobuf/proto"
)

var halToGateTests = []struct {
	name string
	hal  *config.Hal
	want *config.Gate
}{
	{
		"Empty hal config",
		&config.Hal{},
		&config.Gate{},
	},
	{
		"Server SSL config",
		&config.Hal{
			Security: &security.Security{
				ApiSecurity: &security.ApiSecurity{
					Ssl: &security.ApiSsl{
						Enabled:      true,
						KeyAlias:     "alias",
						KeyStore:     "/var/secrets/keystore.jks",
						KeyStoreType: "jks",
						ClientAuth:   security.ClientAuth_WANT,
					},
				},
			},
		},
		&config.Gate{
			Server: &config.ServerConfig{
				Ssl: &security.ApiSsl{
					Enabled:      true,
					KeyAlias:     "alias",
					KeyStore:     "/var/secrets/keystore.jks",
					KeyStoreType: "jks",
					ClientAuth:   security.ClientAuth_WANT,
				},
			},
		},
	},
	{
		"CORS config",
		&config.Hal{
			Security: &security.Security{
				ApiSecurity: &security.ApiSecurity{
					CorsAccessPattern: "https://spinnaker.test/",
				},
			},
		},
		&config.Gate{
			Cors: &config.Cors{
				AllowedOriginsPattern: "https://spinnaker.test/",
			},
		},
	},
	{
		"Basic auth",
		&config.Hal{
			Security: &security.Security{
				Authn: &authn.Authentication{
					Basic: &authn.Basic{
						Enabled: true,
						User: &authn.UsernamePassword{
							Username: "user",
							Password: "passw0rd",
						},
					},
				},
			},
		},
		&config.Gate{
			Security: &config.SpringSecurity{
				Basic: &authn.Basic{
					Enabled: true,
					User: &authn.UsernamePassword{
						Username: "user",
						Password: "passw0rd",
					},
				},
			},
		},
	},
	{
		"Oauth2 enabled",
		&config.Hal{
			Security: &security.Security{
				Authn: &authn.Authentication{
					Oauth2: &authn.OAuth2{
						Enabled: true,
						Client: &authn.OAuth2Client{
							ClientId:     "my-client",
							ClientSecret: "my-secret",
						},
						Provider: authn.OAuth2_GITHUB,
					},
				},
			},
		},
		&config.Gate{
			Security: &config.SpringSecurity{
				Oauth2: &authn.OAuth2{
					Enabled: true,
					Client: &authn.OAuth2Client{
						ClientId:     "my-client",
						ClientSecret: "my-secret",
					},
					Provider: authn.OAuth2_GITHUB,
				},
			},
		},
	},
	{
		// Because Gate does not look at the enabled field for Oauth2 and instead
		// enables Oauth2 any time there is a client id set, we should not write out
		// the config if it is disabled. We can change this and remove the test if
		// we update gate to look at the enabled flag.
		"Oauth2 disabled",
		&config.Hal{
			Security: &security.Security{
				Authn: &authn.Authentication{
					Oauth2: &authn.OAuth2{
						Enabled: false,
						Client: &authn.OAuth2Client{
							ClientId:     "my-client",
							ClientSecret: "my-secret",
						},
						Provider: authn.OAuth2_GITHUB,
					},
				},
			},
		},
		&config.Gate{},
	},
	{
		"SAML",
		&config.Hal{
			Security: &security.Security{
				Authn: &authn.Authentication{
					Saml: &authn.Saml{
						Enabled:     true,
						MetadataUrl: "https://my-saml-provider.test/",
					},
				},
			},
		},
		&config.Gate{
			Saml: &authn.Saml{
				Enabled:     true,
				MetadataUrl: "https://my-saml-provider.test/",
			},
		},
	},
	{
		"LDAP",
		&config.Hal{
			Security: &security.Security{
				Authn: &authn.Authentication{
					Ldap: &authn.Ldap{
						Enabled: true,
						Url:     "https://my-ldap-provider.test",
					},
				},
			},
		},
		&config.Gate{
			Ldap: &authn.Ldap{
				Enabled: true,
				Url:     "https://my-ldap-provider.test",
			},
		},
	},
	{
		"X509",
		&config.Hal{
			Security: &security.Security{
				Authn: &authn.Authentication{
					X509: &authn.X509{
						Enabled: true,
						RoleOid: "abc",
					},
				},
			},
		},
		&config.Gate{
			X509: &authn.X509{
				Enabled: true,
				RoleOid: "abc",
			},
		},
	},
	{
		"IAP",
		&config.Hal{
			Security: &security.Security{
				Authn: &authn.Authentication{
					Iap: &authn.Iap{
						Enabled:   true,
						JwtHeader: "my-header",
						IssuerId:  "abc",
					},
				},
			},
		},
		&config.Gate{
			Google: &config.Gate_GoogleConfig{
				Iap: &authn.Iap{
					Enabled:   true,
					JwtHeader: "my-header",
					IssuerId:  "abc",
				},
			},
		},
	},
	{
		"Canary enabled",
		&config.Hal{
			Canary: &canary.Canary{
				Enabled: true,
			},
		},
		&config.Gate{
			Services: &config.Gate_Services{
				Kayenta: &config.ServiceSettings{
					Enabled: true,
				},
			},
		},
	},
	{
		"Deck base URL",
		&config.Hal{
			Security: &security.Security{
				UiSecurity: &security.UiSecurity{
					OverrideBaseUrl: "https://spinnaker.test:9000",
				},
			},
		},
		&config.Gate{
			Services: &config.Gate_Services{
				Deck: &config.ServiceSettings{
					BaseUrl: "https://spinnaker.test:9000",
				},
			},
		},
	},
}

func TestHalToGate(t *testing.T) {
	for _, tt := range halToGateTests {
		t.Run(tt.name, func(t *testing.T) {
			got := parse_hal.HalToGate(tt.hal)
			if !proto.Equal(got, tt.want) {
				t.Errorf("Expected hal config to generate %v, got %v", tt.want, got)
			}
		})
	}
}
