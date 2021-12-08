package client

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"paxos_go/pb"
	"time"
)

type sendingError struct {
	error
	acceptor *AcceptorServer
	success  bool
}

type promiseErrorWrap struct {
	Promise *pb.Promise
	err     sendingError
}

func Prepare(proposalKey string, version string) error {
	if len(proposalKey) == 0 || len(version) == 0 {
		return errors.New("invalid key or version in Prepare message")
	}

	var acceptors, err = SelectAcceptors()
	if err != nil {
		return fmt.Errorf("selecting acceptors error: %v", err)
	}

	var channels = make([]chan promiseErrorWrap, 10, 10)

	for _, acceptor := range acceptors {
		ch := make(chan promiseErrorWrap)
		go sendPrepareToAcceptor(
			&acceptor,
			&pb.Prepare{
				ProposalKey: proposalKey,
				Version:     version,
			},
			ch)
		channels = append(channels, ch)
	}

	for _, ch := range channels {
		if err := <-ch; !err.err.success {
			_ch := make(chan sendingError)
			tries := 1
			newAcc, _err := ChangeAcceptor(err.err.acceptor, &acceptors)
			if _err == nil {
				go sendPrepareToAcceptor(
					newAcc,
					&pb.Prepare{
						ProposalKey: proposalKey,
						Version:     version,
					},
					ch)
			}
			_err_ := <-_ch
			for !_err_.success && tries < 100 {
				tries++
				newAcc, _err := ChangeAcceptor(err.err.acceptor, &acceptors)
				if _err == nil {
					go sendPrepareToAcceptor(
						newAcc,
						&pb.Prepare{
							ProposalKey: proposalKey,
							Version:     version,
						},
						ch)
				}
				_err_ = <-_ch
			}

			if !_err_.success {
				return errors.New("unable to connect sufficient number of acceptors")
			}
		}
	}

	return nil
}

func sendPrepareToAcceptor(acceptor *AcceptorServer, prepare *pb.Prepare, ch chan<- promiseErrorWrap) {
	conn, err := grpc.Dial(acceptor.Address, grpc.WithInsecure())
	if err != nil {
		ch <- promiseErrorWrap{
			nil,
			sendingError{
				fmt.Errorf("gRPC connection creating error\n%v", err),
				acceptor,
				false,
			},
		}
		return
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Printf("gRPC connection closing error\n%v", err.Error())
		}
	}()

	var client = pb.NewPhaseAClient(conn)

	var ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	promise, err := client.SendPrepare(ctx, prepare)
	if err != nil {
		ch <- promiseErrorWrap{
			nil,
			sendingError{
				fmt.Errorf("sending prepare error: %v", err),
				acceptor,
				false,
			},
		}
		return
	}

	ch <- promiseErrorWrap{
		promise,
		sendingError{
			error:    nil,
			acceptor: nil,
			success:  true,
		},
	}
	return
}
