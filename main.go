package main

import (
	"fmt"
	"github.com/NodeFactoryIo/hactar-daemon/cmd/hactar/commands"
	"github.com/NodeFactoryIo/hactar-daemon/internal/session"
	"github.com/NodeFactoryIo/hactar-daemon/pkg/logger"
	"github.com/NodeFactoryIo/hactar-daemon/pkg/util"
	"github.com/mkideal/cli"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

// Load configuration and initialize commands
func main() {
	// start program
	log.Info("Starting hactar")
	// load config file
	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath(".")      // look for config in the working directory
	util.Must(viper.ReadInConfig(), "Fatal error reading config file")
	// load status file
	status := viper.New()
	status.SetConfigName("status") // name of config file (without extension)
	status.AddConfigPath(".")      // look for config in the working directory
	util.Must(status.ReadInConfig(), "Fatal error reading status file")
	session.LoadSession(status)

	command := os.Args[1:]
	// special handling needed if start command called
	if command[0] == "start" {
		// if new username and password are not provided
		if len(os.Args) != 6 {
			// check for token in status file
			if session.CurrentUser.Token != "" {
				command = append(command, "-e", "token", "-p", "token")
				fmt.Println("Successfully found saved hactar token.")
			}
		}
		// show before credentials prompt
		if len(os.Args) <= 2 && session.CurrentUser.Token == "" {
			fmt.Println("Please enter your hactar credentials.")
		}
	}


	// initialize logger
	logger.SetUpLogger()

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
