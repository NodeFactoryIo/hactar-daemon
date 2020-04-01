package commands

import (
	"fmt"
	"github.com/NodeFactoryIo/hactar-daemon/internal/hactar"
	"github.com/NodeFactoryIo/hactar-daemon/internal/lotus"
	"github.com/NodeFactoryIo/hactar-daemon/internal/session"
	"github.com/NodeFactoryIo/hactar-daemon/internal/stats"
	"github.com/NodeFactoryIo/hactar-daemon/internal/token"
	"github.com/NodeFactoryIo/hactar-daemon/internal/url"
	"github.com/cheynewallace/tabby"
	"github.com/mkideal/cli"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type StartParams struct {
	cli.Helper
	Email    string `cli:"e,email" usage:"hactar account email" prompt:"Enter your email address"`
	Password string `pw:"p,password" usage:"hactar account password" prompt:"Enter your password"`
	Debug    bool   `cli:"d,debug" usage:"turn debug mode, showing all application logs"`
}

func RunStartCommand(ctx *cli.Context) error {
	argv := ctx.Argv().(*StartParams)
	// load current session
	currentSession := session.CurrentSession
	// initialize hactar client && auth
	hactarClient := new(hactar.Client)
	if currentSession.GetHactarToken() != "" {
		// create client with saved token
		hactarClient = hactar.NewClient(currentSession.GetHactarToken())
	} else {
		// create client with provided email and password
		c, err := hactar.NewAuthClient(argv.Email, argv.Password)
		if err != nil {
			log.Error("Failed to authenticate to Hactar service.", err)
			return nil
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
		log.Error("Lotus miner worker is down!")
		return nil
	}
	log.Info("Actor address: ", actorAddress)
	// display token and URL
	nodeUrl := url.GetUrl()
	t := tabby.New()
	t.AddLine("Actor address: ", actorAddress)
	t.AddLine("Node url: ", nodeUrl)
	//t.AddLine("Node token:", token.ReadNodeTokenFromFile())
	//t.AddLine("Miner token:", token.ReadMinerTokenFromFile())
	t.Print()
	// this check for existing nodes is just placeholder
	nodes, _, err := hactarClient.Nodes.GetAllNodes()
	if err == nil {
		// search if node already added
		nodeAdded := false
		for i := range nodes {
			if nodes[i].Address == actorAddress && nodes[i].Url == nodeUrl {
				nodeAdded = true
				break
			}
		}
		// save node to backend if not added
		if !nodeAdded {
			node, resp, err := hactarClient.Nodes.Add(hactar.Node{
				Token: token.ReadNodeTokenFromFile(),
				Node: hactar.NodeInfo{
					Address: actorAddress,
					Url:     nodeUrl,
				},
			})
			if err != nil {
				log.Error("Adding new node failed.", err)
				return nil
			} else if resp != nil && resp.StatusCode == http.StatusCreated {
				log.Info(fmt.Sprintf("New node added, url: %s address: %s", node.Node.Url, node.Node.Address))
			}
		} else {
			log.Info("Node already added.")
		}
	}

	currentSession.SetNodeMinerAddress(actorAddress)

	fmt.Println("Everything is initialized. Starting with monitoring...")

	// start stats monitoring
	stats.StartMonitoringStats(hactarClient, lotusClient)
	stats.StartMonitoringNodeUptime(hactarClient, lotusClient, currentSession)
	stats.StartMonitoringBalance(hactarClient, lotusClient, currentSession)
	stats.StartMonitoringBlocks(hactarClient, lotusClient, currentSession)
	select {}
}
