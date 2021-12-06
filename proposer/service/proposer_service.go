package service

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"paxos_go/pb"
	"strconv"
)

type server struct {
	pb.ProposalExchangeServer
}

func Run(port string) error {
	var _, err = strconv.Atoi(port)
	if err != nil {
		return fmt.Errorf("incorrect port: %v", err)
	}

	listener, err := net.Listen("tcp", port)
	if err != nil {
		return fmt.Errorf("Failed to listen port %s: %v", port, err)
	}
	proposalExchServer := grpc.NewServer()
	pb.RegisterProposalExchangeServer(proposalExchServer, &server{})

	log.Printf("Starting gRPC listener on port %s", port)
	if err := proposalExchServer.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}

	return nil
}
