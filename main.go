package main

import (
	"fmt"
	"github.com/NodeFactoryIo/hactar-daemon/cmd/hactar/commands"
	"github.com/NodeFactoryIo/hactar-daemon/internal/config"
	"github.com/NodeFactoryIo/hactar-daemon/internal/session"
	"github.com/NodeFactoryIo/hactar-daemon/pkg/logger"
	"github.com/mkideal/cli"
	log "github.com/sirupsen/logrus"
	"os"
)

// Load configuration and initialize commands
func main() {
	// config and status setup
	config.InitMainConfig()
	status := config.InitStatusConfig()
	session.InitSession(status)

	// initialize logger
	logger.SetUpLogger()

	// start program
	log.Info("Starting hactar")
	fmt.Println("Starting hactar daemon app...")
	command := os.Args[1:]
	// special handling needed if start command called
	if command[0] == "start" {
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
	).Run(command); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
