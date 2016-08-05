package openid

import (
	"net/http"

	"github.com/ory-am/fosite/token/jwt"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"github.com/ory-am/fosite"
	"github.com/ory-am/fosite/handler/oauth2"
)

type OpenIDConnectImplicitHandler struct {
	AuthorizeImplicitGrantTypeHandler *oauth2.AuthorizeImplicitGrantTypeHandler
	*IDTokenHandleHelper
	ScopeStrategy fosite.ScopeStrategy

	RS256JWTStrategy *jwt.RS256JWTStrategy
}

func (c *OpenIDConnectImplicitHandler) HandleAuthorizeEndpointRequest(ctx context.Context, req *http.Request, ar fosite.AuthorizeRequester, resp fosite.AuthorizeResponder) error {
	if !(ar.GetRequestedScopes().Has("openid") && (ar.GetResponseTypes().Has("token", "id_token") || ar.GetResponseTypes().Exact("id_token"))) {
		return nil
	}

	if !ar.GetClient().GetGrantTypes().Has("implicit") {
		return errors.Wrap(fosite.ErrInvalidGrant, "")
	}

	if ar.GetResponseTypes().Exact("id_token") && !ar.GetClient().GetResponseTypes().Has("id_token") {
		return errors.Wrap(fosite.ErrInvalidGrant, "")
	} else if ar.GetResponseTypes().Matches("token", "id_token") && !ar.GetClient().GetResponseTypes().Has("token", "id_token") {
		return errors.Wrap(fosite.ErrInvalidGrant, "")
	}

	client := ar.GetClient()
	for _, scope := range ar.GetRequestedScopes() {
		if !c.ScopeStrategy(client.GetScopes(), scope) {
			return errors.Wrap(fosite.ErrInvalidScope, "")
		}
	}

	sess, ok := ar.GetSession().(Session)
	if !ok {
		return ErrInvalidSession
	}

	claims := sess.IDTokenClaims()
	if ar.GetResponseTypes().Has("token") {
		if err := c.AuthorizeImplicitGrantTypeHandler.IssueImplicitAccessToken(ctx, req, ar, resp); err != nil {
			return errors.Wrap(err, err.Error())
		}

		ar.SetResponseTypeHandled("token")
		hash, err := c.RS256JWTStrategy.Hash([]byte(resp.GetFragment().Get("access_token")))
		if err != nil {
			return err
		}

		claims.AccessTokenHash = hash[:c.RS256JWTStrategy.GetSigningMethodLength()/2]
	}

	if err := c.IssueImplicitIDToken(ctx, req, ar, resp); err != nil {
		return errors.Wrap(err, err.Error())
	}

	// there is no need to check for https, because implicit flow does not require https
	// https://tools.ietf.org/html/rfc6819#section-4.4.2

	ar.SetResponseTypeHandled("id_token")
	return nil
}