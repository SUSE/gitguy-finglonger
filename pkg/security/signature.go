package security

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"strings"
)

var (
	secret string
)

// IsValidSignature check the signature against the secret to only allow request from github
func IsValidSignature(body []byte, signature string) bool {
	if !strings.HasPrefix(signature, "sha1=") {
		return false
	}
	mac := hmac.New(sha1.New, []byte(secret))
	mac.Write(body)
	actual := mac.Sum(nil)

	expected, err := hex.DecodeString(signature[5:])
	if err != nil {
		return false
	}
	return hmac.Equal(actual, expected)
}
