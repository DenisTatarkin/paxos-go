package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/cristalhq/acmd"
	"log"
	"os"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"paxos_go/pb"
)

const address = "localhost:"

var ctx context.Context
var client pb.ProposalExchangeClient

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Invalid args number\ncmd port_number")
		return
	}

	var _, err = strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal("Invalid args\nPort should be a number")
		return
	}

	conn, err := grpc.Dial(address+os.Args[1], grpc.WithInsecure())
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

	client = pb.NewProposalExchangeClient(conn)

	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	log.Printf("Client estableshed successfully on the port %s", os.Args[1])

	startCLI()
}

func startCLI() {
	var cmds = []acmd.Command{
		{
			Name:        "propose",
			Description: "Prepare and send Proposal to proposer",
			Do:          proposeHandler,
		},
	}

	var runner = acmd.RunnerOf(cmds, acmd.Config{
		AppName: "paxos-cli",
	})

	if err := runner.Run(); err != nil {
		panic(err)
	}
}

func proposeHandler(_ context.Context, args []string) error {
	if len(args) < 2 {
		return errors.New("propose command error: there should be 2 arguments: key and value")
	}

	var r, err = client.SendProposal(ctx, &pb.Proposal{
		Key:   os.Args[0],
		Value: os.Args[1],
	})
	if err != nil {
		return fmt.Errorf("sending proposal error: %v", err.Error())
	}

	log.Printf("sending proposal is successful: %s", r.Value)
	return nil
}
