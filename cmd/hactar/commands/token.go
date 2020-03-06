package commands

import (
	"github.com/NodeFactoryIo/hactar-daemon/internal/token"
	"github.com/cheynewallace/tabby"
	"github.com/mkideal/cli"
)

type TokenParams struct {
	cli.Helper
	Debug    bool `cli:"d,debug" usage:"turn debug mode, showing all application logs"`
}

func RunTokenCommand(ctx *cli.Context) error {
	t := tabby.New()
	t.AddLine("Node token:", token.ReadNodeTokenFromFile())
	t.AddLine("Miner token:", token.ReadMinerTokenFromFile())
	t.Print()
	return nil
}
