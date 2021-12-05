package client

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"os"
	"paxos_go/pb"
	"strconv"
	"time"
)

func Execute(serverAddr string, method string, args []string) {
	if len(args) < 2 {
		log.Fatal("Invalid args number\ncmd port_number")
		return
	}

	var _, err = strconv.Atoi(args[1])
	if err != nil {
		log.Fatal("Invalid args\nPort should be a number")
		return
	}

	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("gRPC connection creating error\n%v", err.Error())
		return
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Fatalf("gRPC connection closing error\n%v", err.Error())
		}
	}()

	var client = pb.NewProposalExchangeClient(conn)

	var ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	log.Printf("Client estableshed successfully on the port %s", os.Args[1])

	switch method {
	case "propose":
		proposeHandler(args, ctx, client)
	default:
		log.Fatalf("No such method: %s", method)
	}
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
