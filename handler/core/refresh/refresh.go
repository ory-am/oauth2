package refresh

import (
	"net/http"
	"time"

	"github.com/go-errors/errors"
	"github.com/ory-am/common/pkg"
	"github.com/ory-am/fosite"
	"github.com/ory-am/fosite/handler/core"
	"golang.org/x/net/context"
)

type RefreshTokenGrantHandler struct {
	AccessTokenStrategy core.AccessTokenStrategy

	RefreshTokenStrategy core.RefreshTokenStrategy

	// RefreshTokenGrantStorage is used to persist session data across requests.
	RefreshTokenGrantStorage RefreshTokenGrantStorage

	// AccessTokenLifespan defines the lifetime of an access token.
	AccessTokenLifespan time.Duration
}

// HandleTokenEndpointRequest implements https://tools.ietf.org/html/rfc6749#section-6
func (c *RefreshTokenGrantHandler) HandleTokenEndpointRequest(ctx context.Context, req *http.Request, request fosite.AccessRequester) error {
	// grant_type REQUIRED.
	// Value MUST be set to "client_credentials".
	if !request.GetGrantTypes().Exact("refresh_token") {
		return errors.New(fosite.ErrUnknownRequest)
	}

	if !request.GetClient().GetGrantTypes().Has("refresh_token") {
		return errors.New(fosite.ErrInvalidGrant)
	}

	// The authorization server MUST ... validate the refresh token.
	signature, err := c.RefreshTokenStrategy.ValidateRefreshToken(ctx, request, req.PostForm.Get("refresh_token"))
	if err != nil {
		return errors.New(fosite.ErrInvalidRequest)
	}

	accessRequest, err := c.RefreshTokenGrantStorage.GetRefreshTokenSession(ctx, signature, nil)
	if err == pkg.ErrNotFound {
		return errors.New(fosite.ErrInvalidRequest)
	} else if err != nil {
		return errors.New(fosite.ErrServerError)
	}

	request.SetScopes(accessRequest.GetScopes())
	for _, scope := range accessRequest.GetGrantedScopes() {
		request.GrantScope(scope)
	}

	// The authorization server MUST ... and ensure that the refresh token was issued to the authenticated client
	if accessRequest.GetClient().GetID() != request.GetClient().GetID() {
		return errors.New(fosite.ErrInvalidRequest)
	}
	return nil
}

// PopulateTokenEndpointResponse implements https://tools.ietf.org/html/rfc6749#section-6
func (c *RefreshTokenGrantHandler) PopulateTokenEndpointResponse(ctx context.Context, req *http.Request, requester fosite.AccessRequester, responder fosite.AccessResponder) error {
	if !requester.GetGrantTypes().Exact("refresh_token") {
		return errors.New(fosite.ErrUnknownRequest)
	}

	signature, err := c.RefreshTokenStrategy.ValidateRefreshToken(ctx, requester, req.PostForm.Get("refresh_token"))
	if err != nil {
		return errors.New(fosite.ErrInvalidRequest)
	}

	accessToken, accessSignature, err := c.AccessTokenStrategy.GenerateAccessToken(ctx, requester)
	if err != nil {
		return errors.New(fosite.ErrServerError)
	}

	refreshToken, refreshSignature, err := c.RefreshTokenStrategy.GenerateRefreshToken(ctx, requester)
	if err != nil {
		return errors.New(fosite.ErrServerError)
	}

	if err := c.RefreshTokenGrantStorage.PersistRefreshTokenGrantSession(ctx, signature, accessSignature, refreshSignature, requester); err != nil {
		return errors.New(fosite.ErrServerError)
	}

	responder.SetAccessToken(accessToken)
	responder.SetTokenType("bearer")
	responder.SetExpiresIn(c.AccessTokenLifespan / time.Second)
	responder.SetScopes(requester.GetGrantedScopes())
	responder.SetExtra("refresh_token", refreshToken)
	return nil
}
