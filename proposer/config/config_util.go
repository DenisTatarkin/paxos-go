package client_config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	client "paxos_go/proposer/client"
)

func GetAcceptors() ([]client.AcceptorServer, error) {
	file, err := ioutil.ReadFile("configs.yaml")
	if err != nil {
		return nil, err
	}

	acceptors := make([]client.AcceptorServer, 10, 20)
	err = yaml.Unmarshal(file, &acceptors)
	if err != nil {
		return nil, err
	}

	return acceptors, nil
}
