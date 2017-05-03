package oauth2

import (
	"github.com/ory/fosite"
	"context"
)

type AuthorizeCodeGrantStorage interface {
	AuthorizeCodeStorage

	PersistAuthorizeCodeGrantSession(ctx context.Context, authorizeCode, accessSignature, refreshSignature string, request fosite.Requester) error
}
