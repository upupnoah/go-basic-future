package web

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestBcrypt(t *testing.T) {
	b, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	if err != nil {
		t.Fatal(err)
	}
	err2 := bcrypt.CompareHashAndPassword(b, []byte("123456"))
	if err2 != nil {
		t.Fatal(err2)
	}
}
