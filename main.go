package main

import (
	"fmt"
	"github.com/NodeFactoryIo/hactar-daemon/cmd/hactar/commands"
	"github.com/NodeFactoryIo/hactar-daemon/internal/config"
	"github.com/NodeFactoryIo/hactar-daemon/internal/session"
	"github.com/mkideal/cli"
	"os"
)

// Load configuration and initialize commands
func main() {
	config.InitMainConfig()
	session.InitSession()

	command := os.Args[1:]
	// special handling needed if start command called
	if len(os.Args) <= 1 {
		command = append(command, "help")
	} else if command[0] == "start" {
		// if new username and password are not provided
		if len(os.Args) != 6 {
			// check for token in status file
			if session.CurrentSession.GetHactarToken() != "" {
				command = append(command, "-e", "token", "-p", "token")
				fmt.Println("Successfully found saved hactar token.")
			}
		}
		// show before credentials prompt
		if len(os.Args) <= 2 && session.CurrentSession.GetHactarToken() == "" {
			fmt.Println("Please enter your hactar credentials:")
		}
	}

	// initialize commands
	if err := cli.Root(commands.RootCommand,
		cli.Tree(cli.HelpCommand("Display help information")),
		cli.Tree(commands.StartCommand),
		cli.Tree(commands.TokenCommand),
		cli.Tree(commands.StatusCommand),
	).Run(command); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
