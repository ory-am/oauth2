package strategy

import (
	"github.com/ory-am/fosite/enigma/jwt"
	"github.com/ory-am/fosite"
	"github.com/stretchr/testify/assert"
	"testing"
)

var j = &JWTStrategy{
	Enigma: &jwt.Enigma{
		PrivateKey: []byte(jwt.TestCertificates[0][1]),
		PublicKey:  []byte(jwt.TestCertificates[1][1]),
	},
}

func TestGenerateIDToken(t *testing.T) {
	req := fosite.NewAccessRequest(&IDTokenSession{
		JWTClaims: &jwt.Claims{},
		JWTHeader: &jwt.Header{},
	})
	token, err := j.GenerateIDToken(nil, nil, req)
	assert.Nil(t, err)
	assert.NotEmpty(t, token)
}