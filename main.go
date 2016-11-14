package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/rosenhouse/cf-filler/vars"

	yaml "gopkg.in/yaml.v2"
)

func mainWithError() error {
	var dnsName string
	var mysqlHost string
	var recipePath string

	flag.StringVar(&dnsName, "dnsname", "myenv.example.com", "DNS name for the deployment")
	flag.StringVar(&mysqlHost, "mysqlHost", "10.0.31.193", "MySQL server host")
	flag.StringVar(&recipePath, "recipe", "", "Recipe file specifying vars to generate")

	flag.Parse()

	if recipePath == "" {
		return fmt.Errorf("missing required flag 'recipe'")
	}

	recipe, err := vars.LoadRecipe(recipePath)
	if err != nil {
		return err
	}

	vars, err := CreateVars(dnsName, mysqlHost, recipe)
	if err != nil {
		return fmt.Errorf("applying config: %s", err)
	}

	outBytes, err := yaml.Marshal(vars)
	if err != nil {
		return fmt.Errorf("marshaling output as yaml: %s", err)
	}
	_, err = os.Stdout.Write(outBytes)
	return err
}

func main() {
	if err := mainWithError(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}
