package main

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"

	"golang.org/x/crypto/ssh"
)

func encodePEM(keyBytes []byte, keyType string) string {
	block := &pem.Block{
		Type:  keyType,
		Bytes: keyBytes,
	}

	return string(pem.EncodeToMemory(block))
}

const (
	pemHeaderPrivateKey = "RSA PRIVATE KEY"
	pemHeaderPublicKey  = "PUBLIC KEY"
)

// PrivateKeyToPEM serializes an RSA Private key into PEM format.
func PrivateKeyToPEM(privateKey *rsa.PrivateKey) string {
	keyBytes := x509.MarshalPKCS1PrivateKey(privateKey)

	return encodePEM(keyBytes, pemHeaderPrivateKey)
}

// PublicKeyToPEM serializes an RSA Public key into PEM format.
func PublicKeyToPEM(publicKey *rsa.PublicKey) (string, error) {
	keyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return "", err
	}

	return encodePEM(keyBytes, pemHeaderPublicKey), nil
}

func GenerateRSAKeyPair() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	private, err := rsa.GenerateKey(rand.Reader, KeyBits)
	if err != nil {
		return nil, nil, err
	}
	public := private.Public().(*rsa.PublicKey)
	return private, public, nil
}

// hexadecimal md5 hash grouped by 2 characters separated by colons
func FingerprintMD5(key ssh.PublicKey) string {
	hash := md5.Sum(key.Marshal())
	out := ""
	for i := 0; i < len(hash); i++ {
		if i > 0 {
			out += ":"
		}
		out += fmt.Sprintf("%02x", hash[i]) // don't forget the leading zeroes
	}
	return out
}

func GenerateSSHKeyAndFingerprint() (string, string, error) {
	priv, pub, err := GenerateRSAKeyPair()
	if err != nil {
		return "", "", fmt.Errorf("generate rsa key pair: %s", err)
	}

	sshPubKey, err := ssh.NewPublicKey(pub)
	if err != nil {
		return "", "", err
	}

	return PrivateKeyToPEM(priv), FingerprintMD5(sshPubKey), nil
}
