package config

import (
	"fmt"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestNewConfigs(t *testing.T) {
	r := Root{}
	d, err := yaml.Marshal(r)
	if err != nil {
		t.Fatalf("err: %s\n", err)
	}
	fmt.Printf("%s\n", d)
}

func TestReadConfigs(t *testing.T) {
	r := ReadConfigs("configs.yaml")
	fmt.Printf("%#v\n", r)
}
