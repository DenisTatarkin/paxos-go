package main

import (
	node "../node"
	"os"
)

var proposers[] node.Proposer
var acceptors[] node.Acceptor

func main(){
	proposers, acceptors = node.ParseConfig(os.Args[1])
}
