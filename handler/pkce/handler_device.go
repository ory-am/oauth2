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

package pkce

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"github.com/ory/x/errorsx"

	"github.com/pkg/errors"

	"github.com/ory/fosite"
	"github.com/ory/fosite/handler/oauth2"
)

type HandlerDevice struct {
	CoreStorage        oauth2.CoreStorage
	DeviceCodeStrategy oauth2.DeviceCodeStrategy
	UserCodeStrategy   oauth2.UserCodeStrategy

	// If set to true, clients must use PKCE.
	Force bool

	// If set to true, public clients must use PKCE.
	ForceForPublicClients bool

	// Whether or not to allow the plain challenge method (S256 should be used whenever possible, plain is really discouraged).
	EnablePlainChallengeMethod bool

	AuthorizeCodeStrategy oauth2.AuthorizeCodeStrategy
	Storage               PKCERequestStorage
}

func (c *HandlerDevice) HandleDeviceAuthorizeEndpointRequest(ctx context.Context, ar fosite.Requester, resp fosite.DeviceAuthorizeResponder) error {

	if !ar.GetClient().GetGrantTypes().Has("urn:ietf:params:oauth:grant-type:device_code") {
		return nil
	}

	challenge := ar.GetRequestForm().Get("code_challenge")
	method := ar.GetRequestForm().Get("code_challenge_method")
	client := ar.GetClient()
	userCode := resp.GetUserCode()

	userCodeSignature := c.UserCodeStrategy.UserCodeSignature(userCode)

	session, err := c.CoreStorage.GetUserCodeSession(ctx, userCodeSignature, fosite.NewRequest().Session)
	if err != nil {
		return err
	}

	if err := c.validate(challenge, method, client); err != nil {
		return err
	}

	if err := c.Storage.CreatePKCERequestSession(ctx, session.GetID(), ar.Sanitize([]string{
		"code_challenge",
		"code_challenge_method",
	})); err != nil {
		return errorsx.WithStack(fosite.ErrServerError.WithWrap(err).WithDebug(err.Error()))
	}

	fmt.Println(resp)

	return nil
}

func (c *HandlerDevice) validate(challenge, method string, client fosite.Client) error {
	if challenge == "" {
		// If the server requires Proof Key for Code Exchange (PKCE) by OAuth
		// clients and the client does not send the "code_challenge" in
		// the request, the authorization endpoint MUST return the authorization
		// error response with the "error" value set to "invalid_request".  The
		// "error_description" or the response of "error_uri" SHOULD explain the
		// nature of error, e.g., code challenge required.

		fmt.Println(c.ForceForPublicClients)
		fmt.Println(client.IsPublic())

		if c.Force {
			return errorsx.WithStack(fosite.ErrInvalidRequest.
				WithHint("Clients must include a code_challenge when performing the authorize code flow, but it is missing.").
				WithDebug("The server is configured in a way that enforces PKCE for clients."))
		}
		if c.ForceForPublicClients && client.IsPublic() {
			return errorsx.WithStack(fosite.ErrInvalidRequest.
				WithHint("This client must include a code_challenge when performing the authorize code flow, but it is missing.").
				WithDebug("The server is configured in a way that enforces PKCE for this client."))
		}
		return nil
	}

	// If the server supporting PKCE does not support the requested
	// transformation, the authorization endpoint MUST return the
	// authorization error response with "error" value set to
	// "invalid_request".  The "error_description" or the response of
	// "error_uri" SHOULD explain the nature of error, e.g., transform
	// algorithm not supported.
	switch method {
	case "S256":
		break
	case "plain":
		fallthrough
	case "":
		if !c.EnablePlainChallengeMethod {
			return errorsx.WithStack(fosite.ErrInvalidRequest.
				WithHint("Clients must use code_challenge_method=S256, plain is not allowed.").
				WithDebug("The server is configured in a way that enforces PKCE S256 as challenge method for clients."))
		}
	default:
		return errorsx.WithStack(fosite.ErrInvalidRequest.
			WithHint("The code_challenge_method is not supported, use S256 instead."))
	}
	return nil
}

func (c *HandlerDevice) HandleTokenEndpointRequest(ctx context.Context, request fosite.AccessRequester) error {
	if !c.CanHandleTokenEndpointRequest(request) {
		return errorsx.WithStack(fosite.ErrUnknownRequest)
	}

	// code_verifier
	// REQUIRED.  Code verifier
	//
	// The "code_challenge_method" is bound to the Authorization Code when
	// the Authorization Code is issued.  That is the method that the token
	// endpoint MUST use to verify the "code_verifier".
	verifier := request.GetRequestForm().Get("code_verifier")

	code := request.GetRequestForm().Get("device_code")
	if code == "" {
		return errorsx.WithStack(errorsx.WithStack(fosite.ErrUnknownRequest.WithHint("device_code missing form body")))
	}
	codeSignature := c.DeviceCodeStrategy.DeviceCodeSignature(code)

	authorizeRequest, err := c.Storage.GetPKCERequestSession(ctx, codeSignature, request.GetSession())
	if errors.Is(err, fosite.ErrNotFound) {
		return errorsx.WithStack(fosite.ErrInvalidGrant.WithHint("Unable to find initial PKCE data tied to this request").WithWrap(err).WithDebug(err.Error()))
	} else if err != nil {
		return errorsx.WithStack(fosite.ErrServerError.WithWrap(err).WithDebug(err.Error()))
	}

	if err := c.Storage.DeletePKCERequestSession(ctx, codeSignature); err != nil {
		return errorsx.WithStack(fosite.ErrServerError.WithWrap(err).WithDebug(err.Error()))
	}

	challenge := authorizeRequest.GetRequestForm().Get("code_challenge")
	method := authorizeRequest.GetRequestForm().Get("code_challenge_method")
	client := authorizeRequest.GetClient()
	if err := c.validate(challenge, method, client); err != nil {
		return err
	}

	if !c.Force && challenge == "" && verifier == "" {
		return nil
	}

	// NOTE: The code verifier SHOULD have enough entropy to make it
	// 	impractical to guess the value.  It is RECOMMENDED that the output of
	// 	a suitable random number generator be used to create a 32-octet
	// 	sequence.  The octet sequence is then base64url-encoded to produce a
	// 	43-octet URL safe string to use as the code verifier.

	// Validation
	if len(verifier) < 43 {
		return errorsx.WithStack(fosite.ErrInvalidGrant.
			WithHint("The PKCE code verifier must be at least 43 characters."))
	} else if len(verifier) > 128 {
		return errorsx.WithStack(fosite.ErrInvalidGrant.
			WithHint("The PKCE code verifier can not be longer than 128 characters."))
	} else if verifierWrongFormat.MatchString(verifier) {
		return errorsx.WithStack(fosite.ErrInvalidGrant.
			WithHint("The PKCE code verifier must only contain [a-Z], [0-9], '-', '.', '_', '~'."))
	}

	// Upon receipt of the request at the token endpoint, the server
	// verifies it by calculating the code challenge from the received
	// "code_verifier" and comparing it with the previously associated
	// "code_challenge", after first transforming it according to the
	// "code_challenge_method" method specified by the client.
	//
	// 	If the "code_challenge_method" from Section 4.3 was "S256", the
	// received "code_verifier" is hashed by SHA-256, base64url-encoded, and
	// then compared to the "code_challenge", i.e.:
	//
	// BASE64URL-ENCODE(SHA256(ASCII(code_verifier))) == code_challenge
	//
	// If the "code_challenge_method" from Section 4.3 was "plain", they are
	// compared directly, i.e.:
	//
	// code_verifier == code_challenge.
	//
	// 	If the values are equal, the token endpoint MUST continue processing
	// as normal (as defined by OAuth 2.0 [RFC6749]).  If the values are not
	// equal, an error response indicating "invalid_grant" as described in
	// Section 5.2 of [RFC6749] MUST be returned.
	switch method {
	case "S256":
		hash := sha256.New()
		if _, err := hash.Write([]byte(verifier)); err != nil {
			return errorsx.WithStack(fosite.ErrServerError.WithWrap(err).WithDebug(err.Error()))
		}

		if base64.RawURLEncoding.EncodeToString(hash.Sum([]byte{})) != challenge {
			return errorsx.WithStack(fosite.ErrInvalidGrant.
				WithHint("The PKCE code challenge did not match the code verifier."))
		}
		break
	case "plain":
		fallthrough
	default:
		if verifier != challenge {
			return errorsx.WithStack(fosite.ErrInvalidGrant.
				WithHint("The PKCE code challenge did not match the code verifier."))
		}
	}

	return nil
}

func (c *HandlerDevice) PopulateTokenEndpointResponse(ctx context.Context, requester fosite.AccessRequester, responder fosite.AccessResponder) error {
	return nil
}

func (c *HandlerDevice) CanSkipClientAuth(requester fosite.AccessRequester) bool {
	return false
}

func (c *HandlerDevice) CanHandleTokenEndpointRequest(requester fosite.AccessRequester) bool {
	// grant_type REQUIRED.
	// Value MUST be set to "authorization_code"
	return requester.GetGrantTypes().ExactOne("urn:ietf:params:oauth:grant-type:device_code")
}
