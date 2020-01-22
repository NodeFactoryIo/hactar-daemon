package stats

import (
	"github.com/NodeFactoryIo/hactar-daemon/internal/lotus/services"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strconv"
	"time"
)

func SubmitNewBlockReport() bool {
	lotusService, err := services.NewLotusService(nil, nil)
	if err != nil {
		// TODO
	}

	if err != nil {
		log.Error("Unable to initialize lotus service", err)
		return false
	}

	log.Info(lotusService.GetLastBlock())
	return true

	//nodeUrl := url.GetUrl()
	//actorAddress, err := lotusService.GetMinerAddress()
	//if err != nil {
	//	log.Error("Unable to send block report because worker is down.")
	//	return false
	//}
}

func StartMonitoringBlocks() {
	interval, _ := strconv.Atoi(viper.GetString("stats.blocks.interval"))
	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				log.Info("Block monitor ticked.")
				SubmitNewBlockReport()
			}
		}
	}()
}
