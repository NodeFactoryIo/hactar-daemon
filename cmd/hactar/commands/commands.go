package commands

import (
	"fmt"
	"github.com/NodeFactoryIo/hactar-daemon/internal/hactar"
	"github.com/NodeFactoryIo/hactar-daemon/internal/lotus/services"
	"github.com/NodeFactoryIo/hactar-daemon/internal/stats"
	"github.com/NodeFactoryIo/hactar-daemon/internal/token"
	"github.com/NodeFactoryIo/hactar-daemon/internal/url"
	"github.com/mkideal/cli"
	log "github.com/sirupsen/logrus"
)

// root command
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

// start command
type startT struct {
	cli.Helper
	Email    string `cli:"e,email" usage:"hactar account email" prompt:"Enter your email address"`
	Password string `pw:"p,password" usage:"hactar account password" prompt:"Enter your password"`
}

var StartCommand = &cli.Command{
	Name: "start",
	Desc: "Start hactar daemon application",
	Text: "",
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*startT)
		client := hactar.NewClient(nil)
		// authenticate
		if err := hactar.Auth(argv.Email, argv.Password); err != nil {
			log.Error("Failed to authenticate to Lotus service.")
		}
		log.Info("Successful authentication.")
		// detect miners and allow user to choose actor address
		lotusService := services.NewLotusService(nil, nil)
		actorAddress, err := lotusService.GetMinerAddress()
		if err != nil {
			fmt.Print("Worker down!")
			return nil
		}
		log.Info("Actor address: ", actorAddress)
		// display token and URL
		token.DisplayToken()
		url.DisplayUrl()
		// save node to backend
		n, r, e := client.Nodes.Add(hactar.Node{
			Token:        token.ReadTokenFromFile(),
			Url:          url.GetUrl(),
			ActorAddress: actorAddress,
		})
		fmt.Println(n)
		fmt.Println(r)
		fmt.Println(e)
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
