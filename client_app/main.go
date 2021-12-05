package main

import cli "paxos_go/client_app/cli"

func main() {
	var app cli.CLIapp
	app.Start()
	if err := app.Runner.Run(); err != nil {
		panic(err)
	}
}
