package core

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

// check go-playground/validator/
type flag struct {
	Use   string `yaml:"use"`
	Short string `yaml:"short"`
	Alias string `yaml:"alias"`
	Type  string `yaml:"type"`
}

// this could be a tree of commands with `Commands []*command`
type command struct {
	Use   string `yaml:"use"`
	Alias string `yaml:"alias"`
	Short string `yaml:"short"`
	Long  string `yaml:"long"`
	Flags []flag `yaml:"flags"`
}

type Config struct {
	Use      string    `yaml:"use"`
	Short    string    `yaml:"short"`
	Long     string    `yaml:"long"`
	Flags    []flag    `yaml:"flags"`
	Commands []command `yaml:"commands"`
}

func ReadConfig() Config {
	var config Config

	// Read the input from stdin
	inputBytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("Error reading stdin:", err)
	}
	// Unmarshal the input YAML into the config struct
	if err := yaml.Unmarshal(inputBytes, &config); err != nil {
		fmt.Println("Error unmarshalling YAML:", err)
	}

	return config
}
