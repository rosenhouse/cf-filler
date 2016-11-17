package creds

import (
	"fmt"

	"github.com/square/certstrap/pkix"
)

const CAExpiryYears = 10
const HostCertExpiryYears = 2

type CA struct {
	CommonName string
	CertPEM    string

	key  *pkix.Key
	cert *pkix.Certificate
}

func NewCA(commonName string) (*CA, error) {
	ca := &CA{CommonName: commonName}

	var err error
	ca.key, err = pkix.CreateRSAKey(KeyBits)
	if err != nil {
		return nil, fmt.Errorf("create ca key: %s", err)
	}

	ca.cert, err = pkix.CreateCertificateAuthority(ca.key, "", CAExpiryYears,
		"", "", "", "", ca.CommonName)
	if err != nil {
		return nil, fmt.Errorf("create ca cert: %s", err)
	}

	caCertPEMBytes, err := ca.cert.Export()
	if err != nil {
		return nil, fmt.Errorf("export pem: %s", err)
	}

	ca.CertPEM = string(caCertPEMBytes)
	return ca, nil
}

// NewCertKeyPair generates a new private key and certificate signed signed by the CA
// The key and cert are returned in PEM format
func (ca *CA) NewCertKeyPair(commonName string, domains []string) (string, string, error) {
	key, err := pkix.CreateRSAKey(KeyBits)
	if err != nil {
		return "", "", fmt.Errorf("create host key: %s", err)
	}
	csr, err := pkix.CreateCertificateSigningRequest(key, "", nil,
		domains, "", "", "", "", commonName)
	if err != nil {
		return "", "", fmt.Errorf("create host csr: %s", err)
	}

	cert, err := pkix.CreateCertificateHost(ca.cert, ca.key, csr, HostCertExpiryYears)
	if err != nil {
		return "", "", fmt.Errorf("sign host csr: %s", err)
	}

	privateKeyPEMBytes, err := key.ExportPrivate()
	if err != nil {
		return "", "", fmt.Errorf("export private key: %s", err)
	}

	certPEMBytes, err := cert.Export()
	if err != nil {
		return "", "", err
	}

	return string(privateKeyPEMBytes), string(certPEMBytes), nil
}
