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
	actorAddress, err := lotusClient.Miner.GetMinerAddress()

	// TODO this uptime check is not valid one! or actor address needs to be saved!
	response, err := hactarClient.Nodes.SendUptimeReport(hactar.UptimeReport{
		IsWorking: err != nil,
		Node: hactar.NodeInfo{
			Address: actorAddress,
			Url:     url.GetUrl(),
		},
	})

	if err != nil {
		log.Error("Unable to send uptime report")
		return false
	}

	if response != nil && response.StatusCode == http.StatusCreated {
		log.Info("Successfully sent uptime report")
		return true
	}

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
