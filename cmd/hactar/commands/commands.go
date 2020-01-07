package commands

import (
	"github.com/NodeFactoryIo/hactar-daemon/internal/hactar"
	"github.com/NodeFactoryIo/hactar-daemon/internal/lotus/services"
	"github.com/NodeFactoryIo/hactar-daemon/internal/stats"
	"github.com/NodeFactoryIo/hactar-daemon/internal/token"
	"github.com/mkideal/cli"
	log "github.com/sirupsen/logrus"
)

// root command
type rootT struct {
	cli.Helper
}

var RootCommand = &cli.Command{
	Desc: "this is root command",
	Argv: func() interface{} { return new(rootT) },
	Fn: func(ctx *cli.Context) error {
		return nil
	},
}

// start command
type startT struct {
	cli.Helper
	Username string `cli:"u,username" usage:"github account" prompt:"Enter Lotus account"`
	Password string `pw:"p,password" usage:"password of github account" prompt:"Enter Lotus password"`
}

var StartCommand = &cli.Command{
	Name: "start",
	Desc: "Start hactar daemon application",
	Text: "",
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*startT)
		// authenticate
		if err := hactar.Auth(argv.Username, argv.Password); err != nil {
			log.Error("Failed to authenticate to Lotus service.")
		}
		log.Info("Successful authentication.")
		// detect miners and allow user to choose actor address
		lotusService := services.NewLotusService()
		actorAddress := lotusService.GetMinerAddress()
		log.Info("Actor address: ", actorAddress)
		// display token and URL
		token.DisplayToken()
		// save node to backend
		hactar.SaveNode(hactar.Node{
			Token:        token.ReadTokenFromFile(),
			Url:          "",
			ActorAddress: actorAddress,
		})
		// start stats monitoring
		stats.StartMonitoringStats()
		select {}
	},
	Argv: func() interface{} { return new(startT) },
}

// token command
type tokenT struct {
	cli.Helper
}

var TokenCommand = &cli.Command{
	Name: "token",
	Desc: "Show lotus token",
	Text: "",
	Fn: func(ctx *cli.Context) error {
		ctx.String("Token: %s\n", token.ReadTokenFromFile())
		return nil
	},
	Argv: func() interface{} { return new(tokenT) },
}
