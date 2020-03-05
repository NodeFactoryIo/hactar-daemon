package commands

import (
	"fmt"
	"github.com/NodeFactoryIo/hactar-daemon/internal/lotus"
	"github.com/NodeFactoryIo/hactar-daemon/internal/session"
	"github.com/NodeFactoryIo/hactar-daemon/internal/url"
	"github.com/mkideal/cli"
	log "github.com/sirupsen/logrus"
)

type StatusParams struct {
	cli.Helper
}

func RunStatusCommand(ctx *cli.Context) error {
	lotusClient, err := lotus.NewClient(nil, nil)
	if err != nil {
		fmt.Println("Unable to initialize lotus client")    // show to user
		log.Error("Failed to initialize lotus client", err) // log
		return err
	}
	actorAddress, err := lotusClient.Miner.GetMinerAddress()
	if err != nil {
		fmt.Println("Lotus worker down")    // show to user
		log.Error("Lotus worker down", err) // log
		return err
	}
	defaultAddress, err := lotusClient.Wallet.GetWalletDefaultAddress()
	if err != nil {
		fmt.Println("Unable to get owner's address")
		log.Error("Failed to get owner's default address")
		return err
	}

	fmt.Printf("Node url: %s\n", url.GetUrl())
	fmt.Printf("Actor address: %s\n", actorAddress)
	fmt.Printf("Owner address: %s\n", defaultAddress)
	fmt.Printf(
		"Last reported tipset of height %d at %s",
		session.CurrentSession.GetLastCheckedHeight(),
		session.CurrentSession.GetLastCheckedHeightTimestamp())

	return nil
}
