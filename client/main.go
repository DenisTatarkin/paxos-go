package main

import (
	"context"
	"github.com/cristalhq/acmd"
	"log"
	"os"
	"strconv"
	"time"

	"google.golang.org/grpc"
	pb "paxos_go/pb"
)

const address = "localhost:"

var ctx context.Context

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

	var client = pb.NewProposalExchangeClient(conn)

	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//todo: console manage passing
	log.Printf("Client estableshed successfully on the port %s", os.Args[1])
	client.SendProposal(ctx, &pb.Proposal{
		Key:   "",
		Value: "",
	})
	//todo: console manage passing
}
