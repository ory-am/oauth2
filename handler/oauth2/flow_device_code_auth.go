package oauth2

import (
	"context"
	"time"

	"github.com/ory/fosite"
	"github.com/ory/x/errorsx"
)

type AuthorizeDeviceGrantTypeHandler struct {
	CoreStorage           CoreStorage
	DeviceCodeStrategy    DeviceCodeStrategy
	UserCodeStrategy      UserCodeStrategy
	AccessTokenStrategy   AccessTokenStrategy
	RefreshTokenStrategy  RefreshTokenStrategy
	AuthorizeCodeStrategy AuthorizeCodeStrategy
	RefreshTokenScopes    []string
	AccessTokenLifespan   time.Duration
	RefreshTokenLifespan  time.Duration
}

func (c *AuthorizeDeviceGrantTypeHandler) HandleAuthorizeEndpointRequest(ctx context.Context, ar fosite.AuthorizeRequester, resp fosite.AuthorizeResponder) error {

	if !ar.GetResponseTypes().ExactOne("device_code") {
		return nil
	}

	if !ar.GetClient().GetGrantTypes().Has("urn:ietf:params:oauth:grant-type:device_code") {
		return nil
	}

	resp.AddParameter("state", ar.GetState())

	userCode := ar.GetRequestForm().Get("user_code")
	userCodeSignature := c.UserCodeStrategy.DeviceCodeSignature(userCode)

	session, err := c.CoreStorage.GetUserCodeSession(ctx, userCodeSignature, ar.GetSession())
	if err != nil {
		return err
	}

	if session.GetClient().GetID() != ar.GetClient().GetID() {
		return errorsx.WithStack(fosite.ErrInvalidGrant.WithHint("The OAuth 2.0 Client ID from this request does not match the one from the authorize request."))
	}

	expires := session.GetSession().GetExpiresAt(fosite.UserCode)
	if time.Now().UTC().After(expires) {
		return errorsx.WithStack(fosite.ErrTokenExpired)
	}

	// session.GetID() is the HMAC signature of the device code generated in the inital request
	err = c.CoreStorage.CreateDeviceCodeSession(ctx, session.GetID(), ar)
	if err != nil {
		return errorsx.WithStack(err)
	}

	ar.SetResponseTypeHandled("device_code")
	return nil
}
