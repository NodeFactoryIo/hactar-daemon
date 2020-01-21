package stats

import (
	"github.com/NodeFactoryIo/hactar-daemon/internal/hactar"
	"github.com/NodeFactoryIo/hactar-daemon/internal/lotus/services"
	"github.com/NodeFactoryIo/hactar-daemon/internal/session"
	"github.com/NodeFactoryIo/hactar-daemon/internal/stats/diskinfo"
	"github.com/NodeFactoryIo/hactar-daemon/internal/url"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strconv"
	"time"
)

func SubmitNewStatsReport() bool {
	client := hactar.NewClient(session.CurrentUser.Token)
	lotusService := services.NewLotusService(nil, nil)

	nodeUrl := url.GetUrl()
	actorAddress, err := lotusService.GetMinerAddress()
	if err != nil {
		log.Error("Unable to send stats report because worker is down.")
		return false
	}

	// send all statistics
	log.Info("Collecting stats and sending report.")
	diskinfo.SendDiskInfoStats(client, actorAddress, nodeUrl)

	return true
}

func StartMonitoringStats() {
	interval, _ := strconv.Atoi(viper.GetString("stats.interval"))
	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				log.Info("Stats monitor ticked.")
				SubmitNewStatsReport()
			}
		}
	}()
}
