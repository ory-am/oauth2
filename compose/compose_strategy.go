// Copyright © 2017 Aeneas Rekkas <aeneas+oss@aeneas.io>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package compose

import (
	"crypto/rsa"

	"github.com/ory/fosite/handler/oauth2"
	"github.com/ory/fosite/handler/openid"
	"github.com/ory/fosite/token/hmac"
	"github.com/ory/fosite/token/jwt"
)

type CommonStrategy struct {
	oauth2.CoreStrategy
	openid.OpenIDConnectTokenStrategy
}

func NewOAuth2HMACStrategy(config *Config, secret []byte) *oauth2.HMACSHAStrategy {
	return &oauth2.HMACSHAStrategy{
		Enigma: &hmac.HMACStrategy{
			GlobalSecret: secret,
		},
		AccessTokenLifespan:   config.GetAccessTokenLifespan(),
		AuthorizeCodeLifespan: config.GetAuthorizeCodeLifespan(),
	}
}

func NewOAuth2JWTStrategy(key *rsa.PrivateKey, strategy *oauth2.HMACSHAStrategy) *oauth2.RS256JWTStrategy {
	return &oauth2.RS256JWTStrategy{
		RS256JWTStrategy: &jwt.RS256JWTStrategy{
			PrivateKey: key,
		},
		HMACSHAStrategy: strategy,
	}
}

func NewOpenIDConnectStrategy(key *rsa.PrivateKey) *openid.DefaultStrategy {
	return &openid.DefaultStrategy{
		RS256JWTStrategy: &jwt.RS256JWTStrategy{
			PrivateKey: key,
		},
	}
}
