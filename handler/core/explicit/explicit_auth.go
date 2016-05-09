package explicit

import (
	"net/http"
	"time"

	"strings"

	"github.com/go-errors/errors"
	. "github.com/ory-am/fosite"
	"github.com/ory-am/fosite/handler/core"
	"golang.org/x/net/context"
)

const authCodeDefaultLifespan = time.Hour / 2

// AuthorizeExplicitGrantTypeHandler is a response handler for the Authorize Code grant using the explicit grant type
// as defined in https://tools.ietf.org/html/rfc6749#section-4.1
type AuthorizeExplicitGrantTypeHandler struct {
	AccessTokenStrategy   core.AccessTokenStrategy
	RefreshTokenStrategy  core.RefreshTokenStrategy
	AuthorizeCodeStrategy core.AuthorizeCodeStrategy

	// AuthorizeCodeGrantStorage is used to persist session data across requests.
	AuthorizeCodeGrantStorage AuthorizeCodeGrantStorage

	// AuthCodeLifespan defines the lifetime of an authorize code.
	AuthCodeLifespan time.Duration

	// AccessTokenLifespan defines the lifetime of an access token.
	AccessTokenLifespan time.Duration
}

func (c *AuthorizeExplicitGrantTypeHandler) HandleAuthorizeEndpointRequest(ctx context.Context, req *http.Request, ar AuthorizeRequester, resp AuthorizeResponder) error {
	// This let's us define multiple response types, for example open id connect's id_token
	if !ar.GetResponseTypes().Exact("code") {
		return nil
	}

	if !ar.GetClient().GetResponseTypes().Has("code") {
		return errors.New(ErrInvalidGrant)
	}

	return c.IssueAuthorizeCode(ctx, req, ar, resp)
}

func (c *AuthorizeExplicitGrantTypeHandler) IssueAuthorizeCode(ctx context.Context, req *http.Request, ar AuthorizeRequester, resp AuthorizeResponder) error {
	code, signature, err := c.AuthorizeCodeStrategy.GenerateAuthorizeCode(ctx, ar)
	if err != nil {
		return errors.New(ErrServerError)
	}

	if err := c.AuthorizeCodeGrantStorage.CreateAuthorizeCodeSession(ctx, signature, ar); err != nil {
		return errors.New(ErrServerError)
	}

	resp.AddQuery("code", code)
	resp.AddQuery("state", ar.GetState())
	resp.AddQuery("scope", strings.Join(ar.GetGrantedScopes(), " "))
	ar.SetResponseTypeHandled("code")
	return nil
}
