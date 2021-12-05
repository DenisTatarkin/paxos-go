package client_config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type ProposerServer struct {
	Id      string
	Address string
}

func GetProposers() ([]ProposerServer, error) {
	file, err := ioutil.ReadFile("configs.yaml")
	if err != nil {
		return nil, err
	}

	proposers := make([]ProposerServer, 10, 20)
	err = yaml.Unmarshal(file, &proposers)
	if err != nil {
		return nil, err
	}

	return proposers, nil
}
