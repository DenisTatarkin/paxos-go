package cli

import (
	"context"
	"github.com/cristalhq/acmd" //Copyright (c) 2021 cristaltech
	client "paxos_go/client"
)

type CLIapp struct {
	cmds   []acmd.Command
	Runner *acmd.Runner
}

func (app *CLIapp) Start() {
	var cmds = []acmd.Command{
		{
			Name:        "propose",
			Description: "Prepare and send Proposal to proposer",
			Do:          proposeHandler,
		},
	}
	app.Runner = acmd.RunnerOf(cmds, acmd.Config{
		AppName: "paxos-cli",
	})
}

func proposeHandler(_ context.Context, args []string) error {
	var err = client.Execute(args[0], args[1], args[2:])
	return err
}
