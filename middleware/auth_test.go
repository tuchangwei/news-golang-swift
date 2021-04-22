package middleware

import (
	"server/utils/result"
	"testing"
)

var email = "1@1.com"
func TestGenerateAndVerifyToken(t *testing.T) {
	tokenStr, err := GenerateToken(email)
	if err != nil {
		t.Fatal("Generate token error:", err)
	}
	code, msg, e := verifyToken(tokenStr)
	if code != result.Success {
		t.Fatal(*msg)
	}
	if email != *e {
		t.Fatal("verify email failed")
	}
}
