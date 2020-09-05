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
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
)

func (f *Fosite) WriteAuthorizeError(rw http.ResponseWriter, ar AuthorizeRequester, err error) {
	rw.Header().Set("Cache-Control", "no-store")
	rw.Header().Set("Pragma", "no-cache")

	rfcerr := ErrorToRFC6749Error(err)
	if !f.SendDebugMessagesToClients {
		rfcerr = rfcerr.Sanitize()
	}

	if !ar.IsRedirectURIValid() {
		js, err := json.MarshalIndent(rfcerr, "", "\t")
		if err != nil {
			if !f.SendDebugMessagesToClients {
				http.Error(rw, fmt.Sprintf(`{\n\t"error": "server_error",\n\t"error_description": "%s"\n}`, err.Error()), http.StatusInternalServerError)
			} else {
				http.Error(rw, `{\n\t"error": "server_error"\n}`, http.StatusInternalServerError)
			}
			return
		}

		rw.Header().Set("Content-Type", "application/json;charset=UTF-8")
		rw.WriteHeader(rfcerr.Code)
		rw.Write(js)
		return
	}

	redirectURI := ar.GetRedirectURI()

	// The endpoint URI MUST NOT include a fragment component.
	redirectURI.Fragment = ""

	query := rfcerr.ToValues()
	query.Add("state", ar.GetState())

	var redirectURIString string
	if !(len(ar.GetResponseTypes()) == 0 || ar.GetResponseTypes().ExactOne("code")) && errors.Cause(err) != ErrUnsupportedResponseType {
		redirectURIString = redirectURI.String() + "#" + query.Encode()
	} else {
		for key, values := range redirectURI.Query() {
			for _, value := range values {
				query.Add(key, value)
			}
		}
		redirectURI.RawQuery = query.Encode()
		redirectURIString = redirectURI.String()
	}

	rw.Header().Add("Location", redirectURIString)
	rw.WriteHeader(http.StatusFound)
}
