package utils

import (
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestEncrypt(t *testing.T) {
	password := "123456"
	got1 := Encrypt(password)
	err := bcrypt.CompareHashAndPassword([]byte(got1), []byte(password))
	if err != nil {
		t.Fatal(err)
	}
}
