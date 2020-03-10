package stats

import (
	"github.com/NodeFactoryIo/hactar-daemon/internal/hactar"
	"github.com/NodeFactoryIo/hactar-daemon/internal/lotus"
	"github.com/NodeFactoryIo/hactar-daemon/internal/session"
	"github.com/NodeFactoryIo/hactar-daemon/internal/url"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"strconv"
	"time"
)

func SubmitNewNodeUptimeReport(hactarClient *hactar.Client, lotusClient *lotus.Client, currentSession session.UserSession) bool {
	isWorking := true
	// check for miner address
	miner, err := lotusClient.Miner.GetMinerAddress()
	if err != nil {
		isWorking = false
	} else {
		// check for last tipset
		tipset, err := lotusClient.Blocks.GetLastTipset()
		if err != nil {
			isWorking = false
		}
		// get actor
		if err == nil {
			actor, _ := lotusClient.Miner.GetActor(miner, tipset.Cids[0])
			isWorking = actor != nil
		}
	}
	// send uptime report
	response, err := hactarClient.Nodes.SendUptimeReport(hactar.UptimeReport{
		IsWorking: isWorking,
		Node: hactar.NodeInfo{
			Address: currentSession.GetNodeMinerAddress(),
			Url:     url.GetUrl(),
		},
	})

	if response != nil && response.StatusCode == http.StatusCreated {
		log.Info("Successfully sent uptime report")
		return true
	}

	log.Error("Unable to send uptime report")
	return false
}

func StartMonitoringNodeUptime(hactarClient *hactar.Client, lotusClient *lotus.Client, currentSession session.UserSession) {
	interval, _ := strconv.Atoi(viper.GetString("stats.uptime.interval"))
	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				log.Info("Node uptime monitor ticked.")
				SubmitNewNodeUptimeReport(hactarClient, lotusClient, currentSession)
			}
		}
	}()
}
