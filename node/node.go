package node

import (
	"bufio"
	"os"
	"strings"
)

type Proposer struct {
	Id string
	Address string
	IsLeader bool
}

type Acceptor struct{
	Id string
	Address string
}

func ParseConfig(filename string) ([]Proposer, []Acceptor) {
	var f,_= os.Open(filename)
	var scanner = bufio.NewScanner(f)
	var proposers[] Proposer
	var acceptors[] Acceptor

	for scanner.Scan(){
		tokens := strings.Split(scanner.Text(), " ")
		switch tokens[2] {
			case "proposer" : proposers = append(proposers,
				Proposer{Id: tokens[0], Address: tokens[1], IsLeader : false})
			case "acceptor" : acceptors = append(acceptors,
				Acceptor{Id: tokens[0], Address: tokens[1]})
			default: panic("Incorrect role")
		}
	}

	return proposers, acceptors
}
