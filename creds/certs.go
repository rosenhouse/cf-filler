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

	ca.CertPEM, err = asString(ca.cert)
	if err != nil {
		return nil, err
	}

	return ca, nil
}

type certKeyPair struct {
	CommonName string
	Domains    []string

	key  *pkix.Key
	cert *pkix.Certificate
}

func (ca *CA) NewCertKeyPair(commonName string, domains []string) (string, string, error) {
	ckp := &certKeyPair{
		CommonName: commonName,
		Domains:    domains,
	}
	var err error

	ckp.key, err = pkix.CreateRSAKey(KeyBits)
	if err != nil {
		return "", "", fmt.Errorf("create host key: %s", err)
	}
	csr, err := pkix.CreateCertificateSigningRequest(ckp.key, "", nil,
		ckp.Domains, "", "", "", "", ckp.CommonName)
	if err != nil {
		return "", "", fmt.Errorf("create host csr: %s", err)
	}

	ckp.cert, err = pkix.CreateCertificateHost(ca.cert, ca.key, csr, HostCertExpiryYears)
	if err != nil {
		return "", "", fmt.Errorf("sign host csr: %s", err)
	}

	privateBytes, err := ckp.key.ExportPrivate()
	if err != nil {
		return "", "", fmt.Errorf("export private key: %s", err)
	}

	cert, err := asString(ckp.cert)
	if err != nil {
		return "", "", err
	}

	return string(privateBytes), cert, nil
}

type exportable interface {
	Export() ([]byte, error)
}

func asString(e exportable) (string, error) {
	pemBytes, err := e.Export()
	if err != nil {
		return "", fmt.Errorf("export pem: %s", err)
	}

	return string(pemBytes), nil
}
