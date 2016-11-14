package vars

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type Recipe struct {
	Strings        map[string]string       `yaml:"strings"`
	Passwords      []string                `yaml:"passwords"`
	PasswordArrays []*PasswordArray        `yaml:"password_arrays"`
	SSHKeys        []*SSHKeyAndFingerprint `yaml:"ssh_keys"`
	BasicKeyPairs  []*BasicKeyPair         `yaml:"basic_key_pairs"`
	CertSets       []*CertSet              `yaml:"cert_sets"`
}

func LoadRecipe(path string) (*Recipe, error) {
	yamlBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading recipe file: %s", err)
	}

	recipe := &Recipe{}
	err = yaml.Unmarshal(yamlBytes, recipe)
	return recipe, err
}
