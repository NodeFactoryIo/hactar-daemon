package stats

import (
	"github.com/NodeFactoryIo/hactar-daemon/internal/hactar"
	"github.com/NodeFactoryIo/hactar-daemon/internal/lotus"
	"github.com/NodeFactoryIo/hactar-daemon/internal/stats/diskinfo"
	"github.com/NodeFactoryIo/hactar-daemon/internal/stats/minerinfo"
	"github.com/NodeFactoryIo/hactar-daemon/internal/stats/pastdealsinfo"
	"github.com/NodeFactoryIo/hactar-daemon/internal/url"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"strconv"
	"time"
)

func SubmitNewStatsReport(hactarClient *hactar.Client, lotusClient *lotus.Client) bool {
	nodeUrl := url.GetUrl()
	actorAddress, err := lotusClient.Miner.GetMinerAddress()
	if err != nil {
		log.Error("Unable to send stats report because worker is down.")
		log.Info("Shutting down hactar reporting...")
		os.Exit(2)
		return false
	}

	// send all statistics
	log.Info("Collecting stats and sending report.")
	diskinfo.SendDiskInfoStats(hactarClient, actorAddress, nodeUrl)
	minerinfo.SendMinerInfoStats(hactarClient, lotusClient)
	pastdealsinfo.SendPastDealsInfo(hactarClient, lotusClient)
	return true
}

func StartMonitoringStats(hactarClient *hactar.Client, lotusClient *lotus.Client) {
	interval, _ := strconv.Atoi(viper.GetString("stats.interval"))
	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	done := make(chan bool)

	// Tick once immediately
	SubmitNewStatsReport(hactarClient, lotusClient)

	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				log.Info("Stats monitor ticked.")
				SubmitNewStatsReport(hactarClient, lotusClient)
			}
		}
	}()
}
