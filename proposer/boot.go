package main

import (
	"../node"
	"net/http"
	"os"
)

var proposers[] node.Proposer
var acceptors[] node.Acceptor
var id string

func main(){
	var port = os.Args[1]
	id = os.Args[2]

	proposers, acceptors = node.ParseConfig(os.Args[3])

	configureHandlers(port)

	go manage()
}

func configureHandlers(port string){
	http.HandleFunc("/client/", clientHandler)
}

func clientHandler(w http.ResponseWriter, r *http.Request) {
	//todo
}

func manage(){
	for true{}
	//todo
}