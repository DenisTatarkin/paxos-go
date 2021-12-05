package client

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"os"
	"paxos_go/pb"
	"time"
)

func Execute(serverAddr string, method string, args []string) error {
	if len(serverAddr) == 0 || len(method) == 0 || len(args) < 2 {
		return errors.New("Invalid args number\ncmd proposer_address proposer_method proposal_key proposal_value")
	}

	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("gRPC connection creating error\n%v", err.Error())
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Printf("gRPC connection closing error\n%v", err.Error())
		}
	}()

	var client = pb.NewProposalExchangeClient(conn)

	var ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	log.Printf("Client estableshed successfully on the port %s", os.Args[1])

	switch method {
	case "propose":
		err := proposeHandler(args, ctx, client)
		if err != nil {
			return err
		}
	default:
		log.Fatalf("No such method: %s", method)
	}

	return nil
}

func proposeHandler(args []string, ctx context.Context, client pb.ProposalExchangeClient) error {
	if len(args) < 2 {
		return errors.New("propose command error: there should be 2 arguments: key and value")
	}

	var r, err = client.SendProposal(ctx, &pb.Proposal{
		Key:   args[0],
		Value: args[1],
	})
	if err != nil {
		return fmt.Errorf("sending proposal error: %v", err.Error())
	}

	log.Printf("sending proposal is successful: %s", r.Value)
	return nil
}
