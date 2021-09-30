package main

import (
	"bufio"
	"fmt"
	"hash/fnv"
	"log"
	"net/http"
	"os"
	"paxos-go/node"
	"strings"
)

var proposers []node.Proposer
var acceptors []node.Acceptor
var id string
var leader node.Proposer
var iLeader = false
var address string

func main() {
	var port = os.Args[1]
	id = os.Args[2]

	proposers, acceptors = node.ParseConfig(os.Args[3])

	configureHandlers(port)

	//todo: this function should be executed after all proposers connected.
	//So should be implemented console command "start" or ssomething like that
	electLeader()

	go manage()

	address = "localhost:" + port

	log.Fatal(http.ListenAndServe(address, nil))
}

func configureHandlers(port string) {
	http.HandleFunc("/", requestHandler)
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	var tokens = strings.Split(r.RequestURI, "/")
	var role = tokens[1]
	var reqBody = strings.Split(tokens[2], "=")
	var key = reqBody[0]
	var value = reqBody[1]

	if role == "client" {
		if iLeader {
			clientHandler(key, value)
		} else {
			ch := make(chan bool)
			go node.Message("leader", key+"="+value, ch)
			//todo: send to leader
			access := <-ch
			if !access {
				fmt.Println("Err")
			}
		}
	}
}

func clientHandler(key string, value string) {
	for _, acceptor := range acceptors {
		var ch = make(chan bool)
		go node.Message(acceptor.Address, key+"="+value, ch)
		access := <-ch
		if !access {
			fmt.Println("Err")
		}
	}
}

func manage() {
	var scanner = bufio.NewScanner(os.Stdin)

	for scanner.Scan() {

	}
}

func electLeader() {
	h := fnv.New32a()
	var hashesProposers []node.Proposer = make([]node.Proposer, len(proposers))

	for _, proposer := range proposers {
		h.Write([]byte(proposer.Address))
		hash := (int)(h.Sum32()) % len(proposers)
		hashesProposers[hash] = proposer
	}

	for _, proposer := range hashesProposers {
		if &proposer != nil {
			leader = proposer
			break
		}
	}

	iLeader = leader.Address == address
	if leader.Address == address {
		iLeader = true
		fmt.Println("This node is leader\n")
	}
}
