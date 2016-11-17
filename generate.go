package main

import (
	"fmt"

	"github.com/rosenhouse/cf-filler/creds"
	"github.com/rosenhouse/cf-filler/vars"
)

type DeploymentVars map[string]interface{}

func (o DeploymentVars) GeneratePasswords(keynames ...string) {
	for _, name := range keynames {
		o[name] = creds.NewPassword()
	}
}

func (o DeploymentVars) GeneratePasswordArray(passArray *vars.PasswordArray) {
	var passwords []string
	for i := 0; i < passArray.Count; i++ {
		passwords = append(passwords, creds.NewPassword())
	}
	o[passArray.VarName] = passwords
}

func (o DeploymentVars) GenerateBasicKeyPair(plainKeyPair *vars.BasicKeyPair) error {
	private, public, err := creds.NewRSAKeyPair()
	if err != nil {
		return fmt.Errorf("create RSA key pair: %s", err)
	}

	o[plainKeyPair.VarName_PublicKey] = public
	o[plainKeyPair.VarName_PrivateKey] = private
	return nil
}

func (o DeploymentVars) GenerateCerts(desiredCertSet *vars.CertSet) error {
	ca, err := creds.NewCA(desiredCertSet.CA.CommonName)
	if err != nil {
		return fmt.Errorf("init ca: %s", err)
	}

	if len(desiredCertSet.CA.VarName_CA) > 0 {
		o[desiredCertSet.CA.VarName_CA] = ca.CertPEM
	}

	for _, ckp := range desiredCertSet.CertKeyPairs {
		private, cert, err := ca.NewCertKeyPair(ckp.CommonName, ckp.Domains)
		if err != nil {
			return err
		}
		o[ckp.VarName_Cert] = cert
		o[ckp.VarName_Key] = private
	}

	return nil
}

func (o DeploymentVars) GenerateSSHKeyAndFingerprint(keyAndFingerprint *vars.SSHKeyAndFingerprint) error {
	sshPrivateKey, sshKeyFingerprint, err := creds.NewSSHKeyAndFingerprint()
	if err != nil {
		return fmt.Errorf("generate ssh key and fingerprint: %s", err)
	}

	o[keyAndFingerprint.VarName_PrivateKey] = sshPrivateKey
	o[keyAndFingerprint.VarName_Fingerprint] = sshKeyFingerprint
	return nil
}
