package fosite

import (
	"context"
	"errors"
	"net/http"

	"github.com/ory/fosite/i18n"
	"github.com/ory/x/errorsx"
)

// NewPushedAuthorizeRequest validates the request and produces an AuthorizeRequester object that can be stored
func (f *Fosite) NewPushedAuthorizeRequest(ctx context.Context, r *http.Request) (AuthorizeRequester, error) {
	request := NewAuthorizeRequest()
	request.Request.Lang = i18n.GetLangFromRequest(f.MessageCatalog, r)

	if r.Method != "" && r.Method != "POST" {
		return request, errorsx.WithStack(ErrInvalidRequest.WithHintf("HTTP method is '%s', expected 'POST'.", r.Method))
	}

	if err := r.ParseMultipartForm(1 << 20); err != nil && err != http.ErrNotMultipart {
		return request, errorsx.WithStack(ErrInvalidRequest.WithHint("Unable to parse HTTP body, make sure to send a properly formatted form request body.").WithWrap(err).WithDebug(err.Error()))
	}
	request.Form = r.Form
	request.State = request.Form.Get("state")

	// Authenticate the client in the same way as at the token endpoint
	// (Section 2.3 of [RFC6749]).
	client, err := f.AuthenticateClient(ctx, r, r.Form)
	if err != nil {
		var rfcerr *RFC6749Error
		if errors.As(err, &rfcerr) && rfcerr.ErrorField != ErrInvalidClient.ErrorField {
			return request, errorsx.WithStack(ErrInvalidClient.WithHint("The requested OAuth 2.0 Client could not be authenticated.").WithWrap(err).WithDebug(err.Error()))
		}

		return request, err
	}
	request.Client = client

	// Reject the request if the "request_uri" authorization request
	// parameter is provided.
	if r.Form.Get("request_uri") != "" {
		return request, errorsx.WithStack(ErrInvalidRequest.WithHint("The request must not contain 'request_uri'."))
	}

	// For private_key_jwt or basic auth client authentication, "client_id" may not inside the form
	// However this is required by NewAuthorizeRequest implementation
	if len(r.Form.Get("client_id")) == 0 {
		r.Form.Set("client_id", client.GetID())
	}

	// Validate as if this is a new authorize request
	fr, err := f.newAuthorizeRequest(ctx, r, true)
	if err != nil {
		return fr, err
	}

	if fr.GetRequestedScopes().Has("openid") && r.Form.Get("redirect_uri") == "" {
		return fr, errorsx.WithStack(ErrInvalidRequest.WithHint("Redirect URI information is required."))
	}

	return fr, nil
}