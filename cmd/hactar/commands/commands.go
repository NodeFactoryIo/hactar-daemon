package commands

import (
	"fmt"
	"github.com/NodeFactoryIo/hactar-daemon/internal/hactar"
	"github.com/NodeFactoryIo/hactar-daemon/internal/lotus/services"
	"github.com/NodeFactoryIo/hactar-daemon/internal/session"
	"github.com/NodeFactoryIo/hactar-daemon/internal/stats"
	"github.com/NodeFactoryIo/hactar-daemon/internal/token"
	"github.com/NodeFactoryIo/hactar-daemon/internal/url"
	"github.com/mkideal/cli"
	log "github.com/sirupsen/logrus"
	"net/http"
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
		client, err := hactar.NewAuthClient(argv.Email, argv.Password)
		// authenticate
		if err != nil {
			log.Error("Failed to authenticate to Hactar service.")
			return err
		}
		log.Info("Successful authentication.")
		// save jwt token for current session
		session.CurrentUser.Token = client.Token
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
		if !client.IsActive() {
			log.Error("Hactar not responding")
			return nil
		}
		node, resp, err := client.Nodes.Add(hactar.Node{
			Token:        token.ReadTokenFromFile(),
			Url:          url.GetUrl(),
			ActorAddress: actorAddress,
		})
		if err != nil {
			log.Error("Adding new node failed.", err)
			return nil
		} else if resp != nil && resp.StatusCode == http.StatusOK {
			log.Info(fmt.Sprintf("New node added, url: %s address: %s", node.Url, node.ActorAddress))
		}
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
