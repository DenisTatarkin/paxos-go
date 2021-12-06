package service

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"paxos_go/pb"
	"strconv"
)

type server struct {
	pb.ProposalExchangeServer
}

func Run() {
	if len(os.Args) < 2 {
		log.Fatal("incorrect number of arguments\ncmd port")
	}

	var _, err = strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("incorrect port: %v", err)
	}

	listener, err := net.Listen("tcp", os.Args[1])
	if err != nil {
		log.Fatalf("Failed to listen port %s: %v", os.Args[1], err)
	}
	proposalExchServer := grpc.NewServer()
	pb.RegisterProposalExchangeServer(proposalExchServer, &server{})

	log.Printf("Starting gRPC listener on port %s", os.Args[1])
	if err := proposalExchServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
