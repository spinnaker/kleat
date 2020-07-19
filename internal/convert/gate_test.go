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

package convert_test

import (
	"testing"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/spinnaker/kleat/api/client"
	"github.com/spinnaker/kleat/api/client/canary"
	"github.com/spinnaker/kleat/api/client/config"
	"github.com/spinnaker/kleat/api/client/security"
	"github.com/spinnaker/kleat/api/client/security/authn"
	"github.com/spinnaker/kleat/internal/convert"
)

var gateTests = configTest{
	generator: func(h *config.Hal) proto.Message { return convert.HalToGate(h) },
	tests: []testCase{
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
							Enabled:      wrapperspb.Bool(true),
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
						Enabled:      wrapperspb.Bool(true),
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
							Enabled: wrapperspb.Bool(true),
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
						Enabled: wrapperspb.Bool(true),
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
							Enabled: wrapperspb.Bool(true),
							Client: &authn.OAuth2Client{
								ClientId:     "my-client",
								ClientSecret: "my-secret",
							},
						},
					},
				},
			},
			&config.Gate{
				Security: &config.SpringSecurity{
					Oauth2: &authn.OAuth2{
						Enabled: wrapperspb.Bool(true),
						Client: &authn.OAuth2Client{
							ClientId:     "my-client",
							ClientSecret: "my-secret",
						},
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
							Enabled: wrapperspb.Bool(false),
							Client: &authn.OAuth2Client{
								ClientId:     "my-client",
								ClientSecret: "my-secret",
							},
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
							Enabled:     wrapperspb.Bool(true),
							MetadataUrl: "https://my-saml-provider.test/",
						},
					},
				},
			},
			&config.Gate{
				Saml: &authn.Saml{
					Enabled:     wrapperspb.Bool(true),
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
							Enabled: wrapperspb.Bool(true),
							Url:     "https://my-ldap-provider.test",
						},
					},
				},
			},
			&config.Gate{
				Ldap: &authn.Ldap{
					Enabled: wrapperspb.Bool(true),
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
							Enabled: wrapperspb.Bool(true),
							RoleOid: "abc",
						},
					},
				},
			},
			&config.Gate{
				X509: &authn.X509{
					Enabled: wrapperspb.Bool(true),
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
							Enabled:   wrapperspb.Bool(true),
							JwtHeader: "my-header",
							IssuerId:  "abc",
						},
					},
				},
			},
			&config.Gate{
				Google: &config.Gate_GoogleConfig{
					Iap: &authn.Iap{
						Enabled:   wrapperspb.Bool(true),
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
					Enabled: wrapperspb.Bool(true),
				},
			},
			&config.Gate{
				Services: &config.Gate_Services{
					Kayenta: &config.ServiceSettings{
						Enabled: wrapperspb.Bool(true),
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
		{
			"Gremlin enabled",
			&config.Hal{
				Features: &client.Features{Gremlin: wrapperspb.Bool(true)},
			},
			&config.Gate{
				Integrations: &config.Gate_Integrations{
					Gremlin: &config.Gate_Integrations_Gremlin{Enabled: wrapperspb.Bool(true)},
				},
			},
		},
	},
}

func TestHalToGate(t *testing.T) {
	runConfigTest(t, gateTests)
}
