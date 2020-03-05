package commands

import (
	"github.com/mkideal/cli"
)

type rootT struct {
	cli.Helper
}

var RootCommand = &cli.Command{
	Argv: func() interface{} { return new(rootT) },
	Desc: "Command line interface for Filecoin nodes monitoring deamon app",
	Fn: func(ctx *cli.Context) error {
		return nil
	},
}

var StartCommand = &cli.Command{
	Name: "start",
	Desc: "Start hactar daemon application",
	Text: "",
	Fn:   RunStartCommand,
	Argv: func() interface{} { return new(StartParams) },
}

var StatusCommand = &cli.Command{
	Name: "status",
	Desc: "",
	Text: "",
	Fn:   RunStatusCommand,
	Argv: func() interface{} { return new(StatusParams) },
}

var TokenCommand = &cli.Command{
	Name: "token",
	Desc: "Show lotus token",
	Text: "",
	Fn:   RunTokenCommand,
	Argv: func() interface{} { return new(TokenParams) },
}
