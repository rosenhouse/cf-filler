package main

import (
	"flag"
	"fmt"
	"os"

	yaml "gopkg.in/yaml.v2"
)

func mainWithError() error {
	var dnsName string
	var mysqlHost string
	flag.StringVar(&dnsName, "dnsname", "myenv.example.com", "DNS name for the deployment")
	flag.StringVar(&mysqlHost, "mysqlHost", "10.0.31.193", "MySQL server host")

	flag.Parse()

	o, err := CreateVars(dnsName, mysqlHost)
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
