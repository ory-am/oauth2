/*
 * Copyright © 2017-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
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
 * @Copyright 	2017-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @license 	Apache-2.0
 *
 */

package fosite

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/dgraph-io/ristretto"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/pkg/errors"

	"github.com/ory/fosite/internal/gen"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	jose "gopkg.in/square/go-jose.v2"
)

func initServerWithKey(t *testing.T) *httptest.Server {
	var set *jose.JSONWebKeySet
	var h http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		require.NoError(t, json.NewEncoder(w).Encode(set))
	}
	ts := httptest.NewServer(h)

	set = &jose.JSONWebKeySet{
		Keys: []jose.JSONWebKey{
			{
				KeyID: "bar",
				Use:   "sig",
				Key:   &gen.MustRSAKey().PublicKey,
			},
		},
	}

	t.Cleanup(ts.Close)
	return ts
}

var errRoundTrip = errors.New("roundtrip error")

type failingTripper struct{}

func (r *failingTripper) RoundTrip(*http.Request) (*http.Response, error) {
	return nil, errRoundTrip
}

func TestDefaultJWKSFetcherStrategy(t *testing.T) {
	ctx := context.Background()
	var h http.HandlerFunc

	s := NewDefaultJWKSFetcherStrategy()
	t.Run("case=fetching", func(t *testing.T) {
		var set *jose.JSONWebKeySet
		h = func(w http.ResponseWriter, r *http.Request) {
			require.NoError(t, json.NewEncoder(w).Encode(set))
		}
		ts := httptest.NewServer(h)
		defer ts.Close()

		set = &jose.JSONWebKeySet{
			Keys: []jose.JSONWebKey{
				{
					KeyID: "foo",
					Use:   "sig",
					Key:   &gen.MustRSAKey().PublicKey,
				},
			},
		}

		keys, err := s.Resolve(ctx, ts.URL, false)
		require.NoError(t, err)
		assert.True(t, len(keys.Key("foo")) == 1)

		set = &jose.JSONWebKeySet{
			Keys: []jose.JSONWebKey{
				{
					KeyID: "bar",
					Use:   "sig",
					Key:   &gen.MustRSAKey().PublicKey,
				},
			},
		}

		keys, err = s.Resolve(ctx, ts.URL, false)
		require.NoError(t, err)
		assert.Len(t, keys.Keys, 1, "%+v", keys)
		assert.True(t, len(keys.Key("foo")) == 1)
		assert.True(t, len(keys.Key("bar")) == 0)

		keys, err = s.Resolve(ctx, ts.URL, true)
		require.NoError(t, err)
		assert.True(t, len(keys.Key("foo")) == 0)
		assert.True(t, len(keys.Key("bar")) == 1)
	})

	t.Run("JWKSFetcherWithCache", func(t *testing.T) {
		ts := initServerWithKey(t)

		cache, _ := ristretto.NewCache(&ristretto.Config{NumCounters: 10 * 1000, MaxCost: 1000, BufferItems: 64})
		location := ts.URL
		expected := &jose.JSONWebKeySet{}
		require.True(t, cache.Set(defaultJWKSFetcherStrategyCachePrefix+location, expected, 1))
		cache.Wait()

		s := NewDefaultJWKSFetcherStrategy(JWKSFetcherWithCache(cache))
		actual, err := s.Resolve(ctx, location, false)
		require.NoError(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("JWKSFetcherWithTTL", func(t *testing.T) {
		ts := initServerWithKey(t)

		s := NewDefaultJWKSFetcherStrategy(JKWKSFetcherWithDefaultTTL(time.Nanosecond))
		_, err := s.Resolve(ctx, ts.URL, false)
		require.NoError(t, err)
		s.(*DefaultJWKSFetcherStrategy).cache.Wait()

		_, ok := s.(*DefaultJWKSFetcherStrategy).cache.Get(defaultJWKSFetcherStrategyCachePrefix + ts.URL)
		assert.Falsef(t, ok, "expected cache to be empty")
	})

	t.Run("JWKSFetcherWithHTTPClient", func(t *testing.T) {
		assert.Panics(t, func() {
			s := NewDefaultJWKSFetcherStrategy(JWKSFetcherWithHTTPClient(&retryablehttp.Client{HTTPClient: &http.Client{Transport: new(failingTripper)}}))
			_, _ = s.Resolve(ctx, "https://google.com", false)
		})
	})

	t.Run("case=error_network", func(t *testing.T) {
		s := NewDefaultJWKSFetcherStrategy()
		h = func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(400)
		}
		ts := httptest.NewServer(h)
		defer ts.Close()

		_, err := s.Resolve(context.Background(), ts.URL, true)
		require.Error(t, err)

		_, err = s.Resolve(context.Background(), "$%/19", true)
		require.Error(t, err)
	})

	t.Run("case=error_encoding", func(t *testing.T) {
		s := NewDefaultJWKSFetcherStrategy()
		h = func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("[]"))
		}
		ts := httptest.NewServer(h)
		defer ts.Close()

		_, err := s.Resolve(context.Background(), ts.URL, true)
		require.Error(t, err)
	})
}
