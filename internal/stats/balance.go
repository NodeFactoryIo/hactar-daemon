package stats

import (
	"github.com/NodeFactoryIo/hactar-daemon/internal/hactar"
	"github.com/NodeFactoryIo/hactar-daemon/internal/lotus"
	"github.com/NodeFactoryIo/hactar-daemon/internal/session"
	"github.com/NodeFactoryIo/hactar-daemon/internal/url"
	"github.com/getsentry/sentry-go"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"strconv"
	"time"
)

func SubmitNewBalanceReport(hactarClient *hactar.Client, lotusClient *lotus.Client, currentSession session.UserSession) bool {
	walletAddress, err := lotusClient.Wallet.GetWalletDefaultAddress()
	if err != nil {
		log.Error("Unable to get wallet address")
		sentry.CaptureException(err)
		return false
	}

	balance, err := lotusClient.Wallet.GetWalletBalance(walletAddress)
	if err != nil {
		log.Error("Unable to get balance for wallet")
		sentry.CaptureException(err)
	}

	response, err := hactarClient.Nodes.SendBalanceReport(hactar.BalanceReport{
		Balance: balance,
		Node: hactar.NodeInfo{
			Address: currentSession.GetNodeMinerAddress(),
			Url:     url.GetUrl(),
		},
	})

	if response != nil && response.StatusCode == http.StatusCreated {
		log.Info("Successfully sent balance report")
		return true
	}

	log.Error("Unable to send balance report", balance)
	sentry.CaptureException(err)
	return false
}

func StartMonitoringBalance(hactarClient *hactar.Client, lotusClient *lotus.Client, currentSession session.UserSession) {
	interval, _ := strconv.Atoi(viper.GetString("stats.balance.interval"))
	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	done := make(chan bool)

	// Tick once immediately
	SubmitNewBalanceReport(hactarClient, lotusClient, currentSession)

	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				log.Info("Node balance monitor ticked.")
				SubmitNewBalanceReport(hactarClient, lotusClient, currentSession)
			}
		}
	}()
}
