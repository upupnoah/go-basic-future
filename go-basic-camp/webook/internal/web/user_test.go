package web

import (
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestEncrypt(t *testing.T) {
	password := []byte("hello#world123#123")
	bcrypted, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		t.Fatal(err)
	}
	err = bcrypt.CompareHashAndPassword(bcrypted, password)
	assert.NoError(t, err)
}

func TestJWT(t *testing.T) {
	token := jwt.New(jwt.SigningMethodRS256)
	s, _ := token.SignedString([]byte("N14EsDR03ubrCCYzHQvIPleU4rli8VzA"))
	t.Logf("token: %s", s)
}
