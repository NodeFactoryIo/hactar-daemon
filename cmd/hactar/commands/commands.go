package commands

import (
	"fmt"
	"github.com/NodeFactoryIo/hactar-daemon/internal/hactar"
	"github.com/NodeFactoryIo/hactar-daemon/internal/lotus"
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

		currentSession := session.CurrentSession

		hactarClient := new(hactar.Client)
		if currentSession.GetHactarToken() != "" {
			// create client with saved token
			hactarClient = hactar.NewClient(currentSession.GetHactarToken())
		} else {
			// create client with provided email and password
			c, err := hactar.NewAuthClient(argv.Email, argv.Password)
			if err != nil {
				log.Error("Failed to authenticate to Hactar service.")
				return err
			}
			hactarClient = c
			// save jwt token for current session
			currentSession.SetHactarToken(hactarClient.Token)
			err = currentSession.SaveSession()
			if err != nil {
				log.Error("Unable to save hactar token.", err)
			}
		}
		log.Info("Successful authentication.")

		// detect miners and allow user to choose actor address
		lotusClient, err := lotus.NewClient(nil, nil)
		if err != nil {
			log.Error("Failed to initialize lotus service")
			return nil
		}
		actorAddress, err := lotusClient.Miner.GetMinerAddress()
		if err != nil {
			log.Error("Worker down!")
			return nil
		}
		log.Info("Actor address: ", actorAddress)
		// display token and URL
		token.DisplayTokens()
		url.DisplayUrl()
		// save node to backend
		node, resp, err := hactarClient.Nodes.Add(hactar.Node{
			Token:        token.ReadNodeTokenFromFile(),
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
		stats.StartMonitoringStats(hactarClient, lotusClient)
		stats.StartMonitoringBlocks(hactarClient, lotusClient, currentSession)
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
		ctx.String(
			"Node token:\n%s\nMiner token:\n%s\n",
			token.ReadNodeTokenFromFile(),
			token.ReadMinerTokenFromFile(),
		)
		return nil
	},
	Argv: func() interface{} { return new(tokenT) },
}
