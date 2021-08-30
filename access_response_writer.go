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

package fosite

import (
	"context"

	"github.com/ory/fosite/i18n"
	"github.com/ory/x/errorsx"
	"golang.org/x/text/language"

	"github.com/pkg/errors"
)

func (f *Fosite) NewAccessResponse(ctx context.Context, requester AccessRequester) (AccessResponder, error) {
	var err error
	var tk TokenEndpointHandler

	response := NewAccessResponse()

	ctx = context.WithValue(ctx, AccessRequestContextKey, requester)
	ctx = context.WithValue(ctx, AccessResponseContextKey, response)

	for _, tk = range f.TokenEndpointHandlers {
		if err = tk.PopulateTokenEndpointResponse(ctx, requester, response); err == nil {
			// do nothing
		} else if errors.Is(err, ErrUnknownRequest) {
			// do nothing
		} else if err != nil {
			return nil, err
		}
	}

	if response.GetAccessToken() == "" || response.GetTokenType() == "" {
		lang := language.English
		g11nContext, ok := requester.(G11NContext)
		if ok {
			lang = g11nContext.GetLang()
		}

		return nil, errorsx.WithStack(ErrServerError.WithLocalizer(f.MessageCatalog, lang).WithHintID(i18n.ErrHintInternalError).WithDebug("Access token or token type not set by TokenEndpointHandlers."))
	}

	return response, nil
}
