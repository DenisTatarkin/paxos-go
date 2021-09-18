package node

import (
	"bufio"
	"fmt"
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
	var f,err = os.Open(filename); if err!=nil{fmt.Println(err.Error()); return nil,nil}; if f == nil{fmt.Println("No config file"); return nil,nil}
	defer f.Close()
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
