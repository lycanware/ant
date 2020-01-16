package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Config structure representation of the yml file
type Config struct {
	Actions []actionItem `yaml:"actions"`
}

type actionItem struct {
	Action        string `yaml:"action"`
	Src           string `yaml:"src"`
	Dst           string `yaml:"dst"`
	DstClearFirst bool   `yaml:"dst-clear-first"`
}

// NewConfig takes the file path of the yml file and calls a load function to read and unmarshal the data.
// An error is returned if the file can't be read or unmarshalling fails
func NewConfig(filePath string) (*Config, error) {
	conf := &Config{}
	err := conf.load(filePath)
	return conf, err
}

func (c *Config) load(filePath string) error {
	var err error
	var fcontent []byte

	if fcontent, err = ioutil.ReadFile(filePath); err != nil {
		return err
	}

	if err = yaml.Unmarshal(fcontent, c); err != nil {
		return err
	}

	return nil
}
