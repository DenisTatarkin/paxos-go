package main

import (
	"../node"
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

var proposers[] node.Proposer
var acceptors[] node.Acceptor
var id string
var leader = false
var address string

func main(){
	var port = os.Args[1]
	id = os.Args[2]

	proposers, acceptors = node.ParseConfig(os.Args[3])

	configureHandlers(port)

	checkIfLeader()

	go manage()

	address = "localhost:" + port

	log.Fatal(http.ListenAndServe(address, nil))
}

func configureHandlers(port string){
	http.HandleFunc("/", requestHandler)
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	var tokens = strings.Split(r.RequestURI, "/")
	var role = tokens[1]
	var reqBody = strings.Split(tokens[2], "=")
	var key = reqBody[0]
	var value = reqBody[1]

	if role == "client" {
		if leader{
			clientHandler(key, value)
		}
		else{
			ch := make(chan <- bool)
			go node.Message("leader", key + "=" + value, ch)
			//todo send to leader
		}
	}
}

func clientHandler(key string, value string) {
	for _,acceptor := range acceptors{
		var ch = make(chan bool)
		go node.Message(acceptor.Address, key + "=" + value, ch)
		access:= <-ch; if !access {fmt.Println("Err")}
	}
}

func manage(){
	var scanner = bufio.NewScanner(os.Stdin)

	for scanner.Scan(){

	}
}

func checkIfLeader() {
	//todo
}