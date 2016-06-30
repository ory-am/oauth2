package client

import (
	"net/http"

	"github.com/ory-am/fosite"
	"github.com/ory-am/fosite/handler/core"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

type ClientCredentialsGrantHandler struct {
	*core.HandleHelper
}

// ValidateTokenEndpointRequest implements https://tools.ietf.org/html/rfc6749#section-4.4.2
func (c *ClientCredentialsGrantHandler) HandleTokenEndpointRequest(ctx context.Context, r *http.Request, request fosite.AccessRequester) (context.Context, error) {
	// grant_type REQUIRED.
	// Value MUST be set to "client_credentials".
	if !request.GetGrantTypes().Exact("client_credentials") {
		return ctx, errors.Wrap(fosite.ErrUnknownRequest, "")
	}

	client := request.GetClient()
	for _, scope := range request.GetScopes() {
		if client.GetGrantedScopes().Grant(scope) {
			request.GrantScope(scope)
		}
	}

	// The client MUST authenticate with the authorization server as described in Section 3.2.1.
	// This requirement is already fulfilled because fosite requries all token requests to be authenticated as described
	// in https://tools.ietf.org/html/rfc6749#section-3.2.1

	// There's nothing else to do. All other security considerations are for the client side.
	return ctx, nil
}

// PopulateTokenEndpointResponse implements https://tools.ietf.org/html/rfc6749#section-4.4.3
func (c *ClientCredentialsGrantHandler) PopulateTokenEndpointResponse(ctx context.Context, r *http.Request, request fosite.AccessRequester, response fosite.AccessResponder) (context.Context, error) {
	if !request.GetGrantTypes().Exact("client_credentials") {
		return ctx, errors.Wrap(fosite.ErrUnknownRequest, "")
	}

	if !request.GetClient().GetGrantTypes().Has("client_credentials") {
		return ctx, errors.Wrap(fosite.ErrInvalidGrant, "")
	}

	return c.IssueAccessToken(ctx, r, request, response)
}
