package commands

import (
	"fmt"
	"github.com/NodeFactoryIo/hactar-daemon/internal/token"
	"github.com/mkideal/cli"
)

type TokenParams struct {
	cli.Helper
}

func RunTokenCommand(ctx *cli.Context) error {
	fmt.Printf(
		"Node token:\n%s\nMiner token:\n%s\n",
		token.ReadNodeTokenFromFile(),
		token.ReadMinerTokenFromFile(),
	)
	return nil
}
