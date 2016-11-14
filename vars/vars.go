// Package vars provides types to describe variables in cf-deployment
package vars

// PlainKeyPair defines a basic RSA public & private key pair
type PlainKeyPair struct {
	VarName_PublicKey  string
	VarName_PrivateKey string
}

// CertSet defines a collection of certificates and keys
// where all the certificates are signed by a common Certificate Authority
type CertSet struct {
	CA           *CA
	CertKeyPairs []*CertKeyPair
}

// CA defines a Certificate Authority
type CA struct {
	VarName_CA string
	CommonName string
}

// CertKeyPair defines a certificate and corresponding private key
type CertKeyPair struct {
	VarName_Cert string
	VarName_Key  string
	CommonName   string
	Domains      []string
}
