package creds

import (
	"crypto/rand"
	"encoding/base64"
	"strings"
)

const PasswordLength = 16

func NewPassword() string {
	bytes := make([]byte, PasswordLength)
	if _, err := rand.Read(bytes); err != nil {
		panic("unable to read rand bytes: " + err.Error())
	}
	return strings.Trim(base64.RawURLEncoding.EncodeToString(bytes), "-_")
}
