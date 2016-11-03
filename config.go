package main

import "fmt"

func CreateVars(dnsName string) (OutputData, error) {
	o := OutputData{}
	system_domain := dnsName
	o["system_domain"] = system_domain
	o["app_domain"] = dnsName
	o.AddSystemComponent("uaa", CfgWithSubdomainURI|CfgWithHTTPSURL)
	o["uaa_token_url"] = fmt.Sprintf("https://%s/oauth/token", o["uaa_uri"])

	o.AddSystemComponent("login", CfgWithSubdomainURI)
	o.AddSystemComponent("api", CfgWithHTTPSURL)
	o.AddSystemComponent("loggregator", CfgNone)
	o.AddSystemComponent("doppler", CfgWithSubdomainURI)
	o.AddSystemComponent("blobstore", CfgNone)
	o["blobstore_public_url"] = fmt.Sprintf("http://%s", o["blobstore_uri"])
	o["blobstore_private_url"] = "https://blobstore.service.cf.internal:4443"
	o["metron_agent_deployment_name"] = dnsName

	o.GeneratePasswords(
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
		"consul_encrypt_keys",
		"diego_bbs_encryption_keys_passphrase",
	)

	o["uaa_scim_users_admin_name"] = "admin"
	o["blobstore_admin_users_username"] = "blobstore-user"
	o["cc_staging_upload_user"] = "staging_user"
	o["cf_mysql_mysql_galera_healthcheck_endpoint_username"] = "galera_healthcheck"
	o["cf_mysql_mysql_seeded_databases_cc_username"] = "cloud_controller"
	o["cf_mysql_mysql_seeded_databases_diego_username"] = "diego"
	o["cf_mysql_mysql_seeded_databases_uaa_username"] = "uaa"
	o["nats_user"] = "nats"
	o["router_status_user"] = "router-status"

	for setName, certSet := range certSets {
		if err := certSet.Generate(o); err != nil {
			return o, fmt.Errorf("generate cert set '%s': %s", setName, err)
		}
	}

	return o, nil
}

var certSets = map[string]*CertSet{
	"etcd_servers": &CertSet{
		CA: &CA{
			VarName_CA: "etcd_ca_cert",
			CommonName: "etcdCA",
		},
		CertKeyPairs: []*CertKeyPair{
			&CertKeyPair{
				VarName_Cert: "etcd_server_cert",
				VarName_Key:  "etcd_server_key",
				CommonName:   "etcd.service.cf.internal",
				Domains: []string{
					"*.etcd.service.cf.internal",
					"etcd.service.cf.internal",
				},
			},
			&CertKeyPair{
				VarName_Cert: "etcd_client_cert",
				VarName_Key:  "etcd_client_key",
				CommonName:   "clientName",
			},
		},
	},

	"etcd_peers": &CertSet{
		CA: &CA{
			VarName_CA: "etcd_peer_ca_cert",
			CommonName: "peerCA",
		},
		CertKeyPairs: []*CertKeyPair{
			&CertKeyPair{
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

	"blobstore": &CertSet{
		CA: &CA{
			VarName_CA: "blobstore_tls_ca_cert",
			CommonName: "blobstore_ca",
		},
		CertKeyPairs: []*CertKeyPair{
			&CertKeyPair{
				VarName_Cert: "blobstore_tls_cert",
				VarName_Key:  "blobstore_tls_private_key",
				CommonName:   "blobstore.service.cf.internal",
			},
		},
	},

	"consul_agent": &CertSet{
		CA: &CA{
			VarName_CA: "consul_agent_ca_cert",
			CommonName: "consulCA",
		},
		CertKeyPairs: []*CertKeyPair{
			&CertKeyPair{
				VarName_Cert: "consul_agent_cert",
				VarName_Key:  "consul_agent_agent_key",
				CommonName:   "consul_agent",
			},
			&CertKeyPair{
				VarName_Cert: "consul_agent_server_cert",
				VarName_Key:  "consul_agent_server_key",
				CommonName:   "server.dc1.cf.internal",
			},
		},
	},

	"diego_bbs": &CertSet{
		CA: &CA{
			VarName_CA: "diego_bbs_ca_cert",
			CommonName: "diegoCA",
		},
		CertKeyPairs: []*CertKeyPair{
			&CertKeyPair{
				VarName_Cert: "diego_bbs_client_cert",
				VarName_Key:  "diego_bbs_client_key",
				CommonName:   "bbs_client",
			},
			&CertKeyPair{
				VarName_Cert: "diego_bbs_server_cert",
				VarName_Key:  "diego_bbs_server_key",
				CommonName:   "bbs.service.cf.internal",
			},
		},
	},

	"loggregator": &CertSet{
		CA: &CA{
			VarName_CA: "loggregator_tls_ca_cert",
			CommonName: "loggregatorCA",
		},
		CertKeyPairs: []*CertKeyPair{
			&CertKeyPair{
				VarName_Cert: "doppler_tls_server_cert",
				VarName_Key:  "doppler_tls_server_key",
				CommonName:   "doppler",
			},
			&CertKeyPair{
				VarName_Cert: "metron_metron_agent_tls_client_cert",
				VarName_Key:  "metron_metron_agent_tls_client_key",
				CommonName:   "metron_agent",
			},
		},
	},
}

/* TODO: add these

diego_bbs_sql_db_connection_string

diego_ssh_proxy_host_key
diego_ssh_proxy_host_key_fingerprint

uaa_jwt_signing_key
uaa_jwt_verification_key

uaa_sslCertificate
uaa_sslPrivateKey

*/
