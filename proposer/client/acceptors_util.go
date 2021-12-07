package client

import (
	"fmt"
	linq "github.com/ahmetb/go-linq/v3"
	"google.golang.org/grpc"
	"log"
	config "paxos_go/proposer/config"
)

type AcceptorServer struct {
	Id      string
	Address string
}

func SelectAcceptors() ([]AcceptorServer, error) {
	var allAcceptors, err = config.GetAcceptors()
	if err != nil {
		return nil, fmt.Errorf("getting acceptors list error:%v", err)
	}

	var amount = len(allAcceptors)/2 + 1
	var selectedAcceptors = make([]AcceptorServer, 10, 10)
	for _, acceptor := range allAcceptors {
		conn, err := grpc.Dial(acceptor.Address, grpc.WithInsecure())
		if err != nil {
			log.Printf("gRPC connection with acceptor %s creating error\n%v", acceptor.Address, err.Error())
			continue
		}

		selectedAcceptors = append(selectedAcceptors, acceptor)

		err = conn.Close()
		if err != nil {
			log.Printf("gRPC connection with acceptor %s closing error\n%v", acceptor.Address, err.Error())
		}

		if len(selectedAcceptors) == amount {
			break
		}
	}

	if len(selectedAcceptors) < amount {
		return nil,
			fmt.Errorf("unable to connect with sufficient acceptors\n%d/%dconnected", len(selectedAcceptors), amount)
	}

	return selectedAcceptors, nil
}

func ChangeAcceptor(acceptor *AcceptorServer, acceptors *[]AcceptorServer) (*AcceptorServer, error) {
	var all, err = config.GetAcceptors()
	if err != nil {
		return nil, fmt.Errorf("getting acceptors list error:%v", err)
	}

	var selectFrom = make([]AcceptorServer, 10, 10)
	linq.From(all).Except(linq.From(acceptors)).ToSlice(&selectFrom)

	var newAcceptor *AcceptorServer

	for _, _newAcceptor := range selectFrom {
		conn, err := grpc.Dial(acceptor.Address, grpc.WithInsecure())
		if err != nil {
			log.Printf("gRPC connection with acceptor %s creating error\n%v", acceptor.Address, err.Error())
			continue
		}

		err = conn.Close()
		if err != nil {
			log.Printf("gRPC connection with acceptor %s closing error\n%v", acceptor.Address, err.Error())
		}

		newAcceptor = &_newAcceptor
	}
	if newAcceptor == nil {
		return nil, fmt.Errorf("unable to connect to some new acceptor instead of %s", acceptor.Address)
	}

	linq.From(acceptors).Except(linq.From(acceptor)).Union(linq.From(newAcceptor)).ToSlice(acceptors)
	return newAcceptor, nil
}
