package commands

import (
	"fmt"
	"github.com/NodeFactoryIo/hactar-daemon/pkg/logger"
	"github.com/mkideal/cli"
	log "github.com/sirupsen/logrus"
	"strings"
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
	Text: "Debug flag [-d][--debug] can be used with any command to show app logs",
	OnRootBefore: func(context *cli.Context) error {
		isDebug := context.IsSet("-d", "-debug")
		if isDebug {
			logger.SetUpLogger(log.InfoLevel)
		} else {
			logger.SetUpDefaultLogger()
		}
		log.Info(fmt.Sprintf("Set logger level: %s", strings.ToUpper(log.GetLevel().String())))
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
	Desc: "Show hactar daemon status",
	Text: "",
	Fn:   RunStatusCommand,
	Argv: func() interface{} { return new(StatusParams) },
}

var TokenCommand = &cli.Command{
	Name: "token",
	Desc: "Show lotus-node and lotus-miner tokens",
	Text: "",
	Fn:   RunTokenCommand,
	Argv: func() interface{} { return new(TokenParams) },
}
