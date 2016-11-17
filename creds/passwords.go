package creds

import (
	"crypto/rand"
	"encoding/base64"
	"strings"
)

// Generate passwords of this length
const PasswordLength = 22

// NewPassword returns a cryptographically secure pseudorandom string
// suitable for use as a password, shared secret, etc.
//
// The returned string will match the regex [a-zA-Z0-9]*
func NewPassword() string {
	bytes := make([]byte, PasswordLength)
	if _, err := rand.Read(bytes); err != nil {
		panic("unable to read rand bytes: " + err.Error())
	}
	return cut(cut(base64.RawURLEncoding.EncodeToString(bytes), "-"), "_")
}

func cut(orig, toCut string) string {
	return strings.Replace(orig, toCut, "", -1)
}
