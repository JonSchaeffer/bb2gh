package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Bitbucket Bitbucket `yaml:"bitbucket"`
	Github		Github    `yaml:"github"`
}

type Bitbucket struct {
	Org string `yaml:"org"`
	Repos []string `yaml:"repos"`
}

type Github struct {
	Org string `yaml:"org"`
}

func ReadConfig(filename string) (Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return Config{}, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil

}
