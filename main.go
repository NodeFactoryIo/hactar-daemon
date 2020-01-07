package main

import (
	"fmt"
	"github.com/NodeFactoryIo/hactar-daemon/cmd/hactar/commands"
	"github.com/NodeFactoryIo/hactar-daemon/internal/util"
	"github.com/NodeFactoryIo/hactar-daemon/pkg/logger"
	"github.com/mkideal/cli"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"strconv"
)

// Load configuration and initialize commands
func main() {
	// load config file
	viper.SetConfigName("config") 			// name of config file (without extension)
	viper.AddConfigPath("/etc/appname/")   	// path to look for the config file in
	viper.AddConfigPath("$HOME/.appname")  	// call multiple times to add many search paths
	viper.AddConfigPath(".")               	// optionally look for config in the working directory
	util.Must(viper.ReadInConfig(), "Fatal error reading config file")

	// initialize logger
	logger.SetUpLogger()

	// start program
	log.Info("Starting hactar")
	interval, _ := strconv.Atoi(viper.GetString("stats.interval"))
	log.Info(fmt.Sprintf("Stats refresh interval is %d minutes.", interval / 60))

	// initialize commands
	if err := cli.Root(commands.RootCommand,
		cli.Tree(cli.HelpCommand("Display help information")),
		cli.Tree(commands.StartCommand),
		cli.Tree(commands.TokenCommand),
	).Run(os.Args[1:]); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}


