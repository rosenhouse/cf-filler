package main

import (
	"fmt"
	"strings"

	"github.com/rosenhouse/cf-filler/vars"
)

func CreateVars(systemDomain, mysqlHost string, recipe *vars.Recipe) (DeploymentVars, error) {
	o := DeploymentVars{}
	o["system_domain"] = systemDomain

	for varName, template := range recipe.Strings {
		o[varName] = strings.Replace(template, "((system_domain))", systemDomain, -1)
	}

	o.GeneratePasswords(recipe.Passwords...)

	for _, pa := range recipe.PasswordArrays {
		o.GeneratePasswordArray(pa)
	}

	for _, certSet := range recipe.CertSets {
		if err := o.GenerateCerts(certSet); err != nil {
			return o, fmt.Errorf("generate cert set: %s", err)
		}
	}

	for _, kp := range recipe.BasicKeyPairs {
		if err := o.GenerateBasicKeyPair(kp); err != nil {
			return o, fmt.Errorf("generate key pair: %s", err)
		}
	}

	for _, kaf := range recipe.SSHKeys {
		if err := o.GenerateSSHKeyAndFingerprint(kaf); err != nil {
			return o, fmt.Errorf("generate ssh creds: %s", err)
		}
	}

	o["diego_bbs_sql_db_connection_string"] = fmt.Sprintf("%s:%s@tcp(%s:3306)/diego",
		o["cf_mysql_mysql_seeded_databases_diego_username"],
		o["cf_mysql_mysql_seeded_databases_diego_password"],
		mysqlHost)

	return o, nil
}
