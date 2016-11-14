package main

import (
	"fmt"
	"strings"

	"github.com/rosenhouse/cf-filler/vars"
)

func CreateVars(systemDomain, mysqlHost string) (DeploymentVars, error) {
	o := DeploymentVars{}
	o["system_domain"] = systemDomain

	for varName, template := range basicStrings {
		o[varName] = strings.Replace(template, "((system_domain))", systemDomain, -1)
	}

	o.GeneratePasswords(passwords...)

	for _, pa := range passwordArrays {
		o.GeneratePasswordArray(pa.VarName, pa.NumPasswords)
	}

	for setName, certSet := range certSets {
		if err := o.GenerateCerts(certSet); err != nil {
			return o, fmt.Errorf("generate cert set '%s': %s", setName, err)
		}
	}

	for _, kp := range basicKeyPairs {
		if err := o.GenerateBasicKeyPair(kp); err != nil {
			return o, fmt.Errorf("generate key pair: %s", err)
		}
	}

	for _, kaf := range sshKeys {
		err := o.GenerateSSHKeyAndFingerprint(kaf.VarName_PrivateKey, kaf.VarName_Fingerprint)
		if err != nil {
			return o, fmt.Errorf("generate ssh creds: %s", err)
		}
	}

	o["diego_bbs_sql_db_connection_string"] = fmt.Sprintf("%s:%s@tcp(%s:3306)/diego",
		o["cf_mysql_mysql_seeded_databases_diego_username"],
		o["cf_mysql_mysql_seeded_databases_diego_password"],
		mysqlHost)

	return o, nil
}

var sshKeys = []*vars.SSHKeyAndFingerprint{
	&vars.SSHKeyAndFingerprint{
		VarName_PrivateKey:  "diego_ssh_proxy_host_key",
		VarName_Fingerprint: "diego_ssh_proxy_host_key_fingerprint",
	},
}

var basicKeyPairs = []*vars.BasicKeyPair{
	&vars.BasicKeyPair{
		VarName_PrivateKey: "uaa_jwt_signing_key",
		VarName_PublicKey:  "uaa_jwt_verification_key",
	},
}

var certSets = map[string]*vars.CertSet{
	"etcd_servers": &vars.CertSet{
		CA: &vars.CA{
			VarName_CA: "etcd_ca_cert",
			CommonName: "etcdCA",
		},
		CertKeyPairs: []*vars.CertKeyPair{
			&vars.CertKeyPair{
				VarName_Cert: "etcd_server_cert",
				VarName_Key:  "etcd_server_key",
				CommonName:   "etcd.service.cf.internal",
				Domains: []string{
					"*.etcd.service.cf.internal",
					"etcd.service.cf.internal",
				},
			},
			&vars.CertKeyPair{
				VarName_Cert: "etcd_client_cert",
				VarName_Key:  "etcd_client_key",
				CommonName:   "clientName",
			},
		},
	},

	"etcd_peers": &vars.CertSet{
		CA: &vars.CA{
			VarName_CA: "etcd_peer_ca_cert",
			CommonName: "peerCA",
		},
		CertKeyPairs: []*vars.CertKeyPair{
			&vars.CertKeyPair{
				VarName_Cert: "etcd_peer_cert",
				VarName_Key:  "etcd_peer_key",
				CommonName:   "etcd.service.cf.internal",
				Domains: []string{
					"*.etcd.service.cf.internal",
					"etcd.service.cf.internal",
				},
			},
		},
	},

	"blobstore": &vars.CertSet{
		CA: &vars.CA{
			VarName_CA: "blobstore_tls_ca_cert",
			CommonName: "blobstore_ca",
		},
		CertKeyPairs: []*vars.CertKeyPair{
			&vars.CertKeyPair{
				VarName_Cert: "blobstore_tls_cert",
				VarName_Key:  "blobstore_tls_private_key",
				CommonName:   "blobstore.service.cf.internal",
			},
		},
	},

	"consul_agent": &vars.CertSet{
		CA: &vars.CA{
			VarName_CA: "consul_agent_ca_cert",
			CommonName: "consulCA",
		},
		CertKeyPairs: []*vars.CertKeyPair{
			&vars.CertKeyPair{
				VarName_Cert: "consul_agent_cert",
				VarName_Key:  "consul_agent_agent_key",
				CommonName:   "consul_agent",
			},
			&vars.CertKeyPair{
				VarName_Cert: "consul_agent_server_cert",
				VarName_Key:  "consul_agent_server_key",
				CommonName:   "server.dc1.cf.internal",
			},
		},
	},

	"diego_bbs": &vars.CertSet{
		CA: &vars.CA{
			VarName_CA: "diego_bbs_ca_cert",
			CommonName: "diegoCA",
		},
		CertKeyPairs: []*vars.CertKeyPair{
			&vars.CertKeyPair{
				VarName_Cert: "diego_bbs_client_cert",
				VarName_Key:  "diego_bbs_client_key",
				CommonName:   "bbs_client",
			},
			&vars.CertKeyPair{
				VarName_Cert: "diego_bbs_server_cert",
				VarName_Key:  "diego_bbs_server_key",
				CommonName:   "bbs.service.cf.internal",
			},
		},
	},

	"loggregator": &vars.CertSet{
		CA: &vars.CA{
			VarName_CA: "loggregator_tls_ca_cert",
			CommonName: "loggregatorCA",
		},
		CertKeyPairs: []*vars.CertKeyPair{
			&vars.CertKeyPair{
				VarName_Cert: "doppler_tls_server_cert",
				VarName_Key:  "doppler_tls_server_key",
				CommonName:   "doppler",
			},
			&vars.CertKeyPair{
				VarName_Cert: "metron_metron_agent_tls_client_cert",
				VarName_Key:  "metron_metron_agent_tls_client_key",
				CommonName:   "metron_agent",
			},
			&vars.CertKeyPair{
				VarName_Cert: "loggregator_tls_doppler_cert",
				VarName_Key:  "loggregator_tls_doppler_key",
				CommonName:   "doppler",
			},
			&vars.CertKeyPair{
				VarName_Cert: "loggregator_tls_tc_cert",
				VarName_Key:  "loggregator_tls_tc_key",
				CommonName:   "trafficcontroller",
			},
		},
	},

	"uaa": &vars.CertSet{
		CA: &vars.CA{CommonName: "uaaCA"},
		CertKeyPairs: []*vars.CertKeyPair{
			&vars.CertKeyPair{
				VarName_Cert: "uaa_sslCertificate",
				VarName_Key:  "uaa_sslPrivateKey",
				CommonName:   "uaa.service.cf.internal",
			},
		},
	},
}

var basicStrings = map[string]string{
	"app_domain": "((system_domain))",

	"uaa_uri":           "uaa.((system_domain))",
	"uaa_subdomain_uri": "*.uaa.((system_domain))",
	"uaa_url":           "https://uaa.((system_domain))",
	"uaa_token_url":     "https://uaa.((system_domain))/oauth/token",

	"login_uri":           "login.((system_domain))",
	"login_subdomain_uri": "*.login.((system_domain))",

	"api_uri": "api.((system_domain))",
	"api_url": "https://api.((system_domain))",

	"loggregator_uri":              "loggregator.((system_domain))",
	"doppler_uri":                  "doppler.((system_domain))",
	"doppler_subdomain_uri":        "*.doppler.((system_domain))",
	"metron_agent_deployment_name": "((system_domain))",

	"blobstore_uri":         "blobstore.((system_domain))",
	"blobstore_public_url":  "https://blobstore.((system_domain))",
	"blobstore_private_url": "https://blobstore.service.cf.internal:4443",

	"uaa_scim_users_admin_name":                           "admin",
	"blobstore_admin_users_username":                      "blobstore-user",
	"cc_staging_upload_user":                              "staging_user",
	"cf_mysql_mysql_galera_healthcheck_endpoint_username": "galera_healthcheck",
	"cf_mysql_mysql_seeded_databases_cc_username":         "cloud_controller",
	"cf_mysql_mysql_seeded_databases_diego_username":      "diego",
	"cf_mysql_mysql_seeded_databases_uaa_username":        "uaa",
	"nats_user":          "nats",
	"router_status_user": "router-status",
}

var passwords = []string{
	"blobstore_admin_users_password",
	"blobstore_secure_link_secret",
	"cc_bulk_api_password",
	"cc_db_encryption_key",
	"cc_internal_api_password",
	"cc_staging_upload_password",
	"cf_mysql_mysql_admin_password",
	"cf_mysql_mysql_cluster_health_password",
	"cf_mysql_mysql_galera_healthcheck_endpoint_password",
	"cf_mysql_mysql_galera_healthcheck_password",
	"cf_mysql_mysql_roadmin_password",
	"cf_mysql_mysql_seeded_databases_cc_password",
	"cf_mysql_mysql_seeded_databases_diego_password",
	"cf_mysql_mysql_seeded_databases_uaa_password",
	"nats_password",
	"router_status_password",
	"uaa_scim_users_admin_password",
	"dropsonde_shared_secret",
	"router_route_services_secret",
	"uaa_admin_client_secret",
	"uaa_clients_cc-routing_secret",
	"uaa_clients_cc-service-dashboards_secret",
	"uaa_clients_cloud_controller_username_lookup_secret",
	"uaa_clients_doppler_secret",
	"uaa_clients_gorouter_secret",
	"uaa_clients_ssh-proxy_secret",
	"uaa_clients_tcp_emitter_secret",
	"uaa_clients_tcp_router_secret",
	"uaa_login_client_secret",
	"diego_bbs_encryption_keys_passphrase",
}

var passwordArrays = []*vars.PasswordArray{
	&vars.PasswordArray{
		VarName:      "consul_encrypt_keys",
		NumPasswords: 1,
	},
}
