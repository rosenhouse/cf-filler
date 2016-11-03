package main

import (
	"flag"
	"fmt"
	"os"

	yaml "gopkg.in/yaml.v2"
)

func mainWithError() error {
	var dnsName string
	flag.StringVar(&dnsName, "dnsname", "myenv.example.com", "DNS name for environment")

	flag.Parse()

	o, err := CreateVars(dnsName)
	if err != nil {
		return fmt.Errorf("applying config: %s", err)
	}

	outBytes, err := yaml.Marshal(o)
	if err != nil {
		return fmt.Errorf("marshaling output as yaml: %s", err)
	}
	os.Stdout.Write(outBytes)

	return nil
}

func main() {
	if err := mainWithError(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}

/* TODO: add these

blobstore_tls_ca_cert
	blobstore_tls_cert
	blobstore_tls_private_key

consul_agent_ca_cert
	consul_agent_agent_key
	consul_agent_cert
	consul_agent_server_cert
	consul_agent_server_key

diego_bbs_ca_cert
	diego_bbs_client_cert
	diego_bbs_client_key
	diego_bbs_server_cert
	diego_bbs_server_key

diego_bbs_sql_db_connection_string

diego_ssh_proxy_host_key
diego_ssh_proxy_host_key_fingerprint

loggregator_tls_ca_cert
	doppler_tls_server_cert
	doppler_tls_server_key
	metron_metron_agent_tls_client_cert
	metron_metron_agent_tls_client_key


uaa_jwt_signing_key
uaa_jwt_verification_key

uaa_sslCertificate
uaa_sslPrivateKey

*/
