package strategy

import (
	"testing"

	"github.com/ory-am/fosite"
	"github.com/ory-am/fosite/token/jwt"
	"github.com/stretchr/testify/assert"
	"time"
)

var j = &DefaultIDTokenStrategy{
	RS256JWTStrategy: &jwt.RS256JWTStrategy{
		PrivateKey: []byte(jwt.TestCertificates[0][1]),
		PublicKey:  []byte(jwt.TestCertificates[1][1]),
	},
}

func TestGenerateIDToken(t *testing.T) {
	var req *fosite.AccessRequest
	for k, c := range []struct {
		setup     func()
		expectErr bool
	}{
		{
			setup: func() {
				req = fosite.NewAccessRequest(&IDTokenSession{
					Claims:  &jwt.IDTokenClaims{
						Subject: "peter",
					},
					Headers: &jwt.Header{},
				})
				req.Form.Set("nonce", "some-secure-nonce-state")
			},
			expectErr: false,
		},
		{
			setup: func() {
				req = fosite.NewAccessRequest(&IDTokenSession{
					Claims:  &jwt.IDTokenClaims{
						Subject: "peter",
						AuthTime: time.Now(),
					},
					Headers: &jwt.Header{},
				})
				req.Form.Set("nonce", "some-secure-nonce-state")
				req.Form.Set("max_age", "1234")
			},
			expectErr: false,
		},
		{
			setup: func() {
				req = fosite.NewAccessRequest(&IDTokenSession{
					Claims:  &jwt.IDTokenClaims{
						Subject: "peter",
						ExpiresAt: time.Now().Add(-time.Hour),
					},
					Headers: &jwt.Header{},
				})
				req.Form.Set("nonce", "some-secure-nonce-state")
			},
			expectErr: true,
		},
		{
			setup: func() {
				req = fosite.NewAccessRequest(&IDTokenSession{
					Claims:  &jwt.IDTokenClaims{
						Subject: "peter",
					},
					Headers: &jwt.Header{},
				})
				req.Form.Set("nonce", "some-secure-nonce-state")
				req.Form.Set("max_age", "1234")
			},
			expectErr: true,
		},
		{
			setup: func() {
				req = fosite.NewAccessRequest(&IDTokenSession{
					Claims:  &jwt.IDTokenClaims{},
					Headers: &jwt.Header{},
				})
				req.Form.Set("nonce", "some-secure-nonce-state")
			},
			expectErr: true,
		},
		{
			setup: func() {
				req = fosite.NewAccessRequest(&IDTokenSession{
					Claims:  &jwt.IDTokenClaims{
						Subject: "peter",
					},
					Headers: &jwt.Header{},
				})
			},
			expectErr: true,
		},
	} {
		c.setup()
		token, err := j.GenerateIDToken(nil, nil, req)
		assert.Equal(t, c.expectErr, err != nil, "%d: %s", k, err)
		if !c.expectErr {
			assert.NotEmpty(t, token)
		}
	}
}
