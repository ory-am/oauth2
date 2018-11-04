/*
 * Copyright © 2015-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * @author		Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @copyright 	2015-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @license 	Apache-2.0
 *
 */

package integration_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/ory/fosite"
	"github.com/ory/fosite/compose"
	"github.com/ory/fosite/handler/openid"
	"github.com/ory/fosite/internal"
	"github.com/ory/fosite/token/jwt"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	goauth "golang.org/x/oauth2"
)

func TestOIDCImplicitFlowPublicClientPKCE(t *testing.T) {
	session := &defaultSession{
		DefaultSession: &openid.DefaultSession{
			Claims: &jwt.IDTokenClaims{
				Subject: "peter",
			},
			Headers: &jwt.Headers{},
		},
	}
	f := compose.ComposeAllEnabled(new(compose.Config), fositeStore, []byte("some-secret-thats-random-some-secret-thats-random-"), internal.MustRSAKey())
	ts := mockServer(t, f, session)
	defer ts.Close()

	oauthClient := newOAuth2Client(ts)

	oauthClient.ClientSecret = ""
	oauthClient.ClientID = "public-client"
	oauthClient.Scopes = []string{"openid"}

	fositeStore.Clients["public-client"].(*fosite.DefaultClient).RedirectURIs[0] = ts.URL + "/callback"

	var state = "12345678901234567890"
	for k, c := range []struct {
		responseType  string
		description   string
		nonce         string
		setup         func()
		codeVerifier  string
		codeChallenge string
	}{
		{

			responseType:  "id_token%20code",
			nonce:         "1111111111111111",
			description:   "should pass id token (id_token code) with PKCE applied.",
			setup:         func() {},
			codeVerifier:  "e7343b9bee0847e3b589ccb60d124ff81adcba6067b84f79b092f86249111fdc",
			codeChallenge: "J11vOtKUitab04a_N0Ogm0dQBytTgl0fgHzYk4xUryo",
		},
	} {
		t.Run(fmt.Sprintf("case=%d/description=%s", k, c.description), func(t *testing.T) {
			c.setup()

			var callbackURL *url.URL
			authURL := strings.Replace(oauthClient.AuthCodeURL(state), "response_type=code", "response_type="+c.responseType, -1) +
				"&nonce=" + c.nonce + "&code_challenge_method=S256&code_challenge=" + c.codeChallenge
			client := &http.Client{
				CheckRedirect: func(req *http.Request, via []*http.Request) error {
					callbackURL = req.URL
					return errors.New("Dont follow redirects")
				},
			}
			resp, err := client.Get(authURL)
			require.Error(t, err)

			t.Logf("Response (%d): %s", k, callbackURL.String())
			fragment, err := url.ParseQuery(callbackURL.Fragment)
			require.NoError(t, err)

			code := fragment.Get("code")
			assert.NotEmpty(t, code)

			assert.NotEmpty(t, fragment.Get("id_token"))

			resp, err = http.PostForm(oauthClient.Endpoint.TokenURL, url.Values{
				"code":          {code},
				"grant_type":    {"authorization_code"},
				"client_id":     {"public-client"},
				"redirect_uri":  {ts.URL + "/callback"},
				"code_verifier": {c.codeVerifier},
			})
			require.NoError(t, err)
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			require.NoError(t, err)

			assert.Equal(t, resp.StatusCode, http.StatusOK)
			token := goauth.Token{}
			require.NoError(t, json.Unmarshal(body, &token))

			require.NotEmpty(t, token.AccessToken, "Got body: %s", string(body))

			httpClient := oauthClient.Client(goauth.NoContext, &token)
			resp, err = httpClient.Get(ts.URL + "/info")
			require.NoError(t, err)
			assert.Equal(t, http.StatusNoContent, resp.StatusCode)

			t.Logf("Passed test case (%d) %s", k, c.description)
		})
	}
}
