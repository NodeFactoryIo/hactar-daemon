package commands

import (
	"fmt"
	"github.com/NodeFactoryIo/hactar-daemon/internal/lotus"
	"github.com/NodeFactoryIo/hactar-daemon/internal/session"
	"github.com/NodeFactoryIo/hactar-daemon/internal/url"
	"github.com/cheynewallace/tabby"
	"github.com/mkideal/cli"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type StatusParams struct {
	cli.Helper
	Debug    bool `cli:"d,debug" usage:"turn debug mode, showing all application logs"`
}

func RunStatusCommand(ctx *cli.Context) error {
	lotusClient, err := lotus.NewClient(nil, nil)
	if err != nil {
		fmt.Println("Unable to detect lotus client")    // show to user
		log.Error("Failed to initialize lotus client", err) // log
		return err
	}
	actorAddress, err := lotusClient.Miner.GetMinerAddress()
	if err != nil {
		fmt.Println("Lotus miner worker down")    // show to user
		log.Error("Lotus miner worker down", err) // log
		return err
	}
	defaultAddress, err := lotusClient.Wallet.GetWalletDefaultAddress()
	if err != nil {
		fmt.Println("Unable to get owner's address")       // show to user
		log.Error("Failed to get owner's default address") // log
		return err
	}

	t := tabby.New()
	t.AddHeader("Hactar daemon app stats")
	t.AddLine("Actor address:", actorAddress)
	t.AddLine("Node url:", url.GetUrl())
	t.AddLine("Owner address:", defaultAddress)
	t.AddLine("")
	t.Print()

	t = tabby.New()
	t.AddHeader("Monitoring block rewards")
	t.AddLine("Last reported tipset height:", session.CurrentSession.GetLastCheckedHeight())
	t.AddLine("Last tipset reported at:", session.CurrentSession.GetLastCheckedHeightTimestamp())
	t.AddLine("")
	t.Print()

	t = tabby.New()
	t.AddHeader("Reporting stats on the following intervals (seconds):")
	t.AddLine("General info report", viper.GetString("stats.interval"))
	t.AddLine("Uptime report", viper.GetString("stats.uptime.interval"))
	t.AddLine("Block reward report", viper.GetString("stats.blocks.interval"))
	t.AddLine("Balance report", viper.GetString("stats.balance.interval"))
	t.Print()

	return nil
}
