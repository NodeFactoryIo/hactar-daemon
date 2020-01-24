package stats

import (
	"github.com/NodeFactoryIo/hactar-daemon/internal/lotus"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strconv"
	"time"
)

func SubmitNewBlockReport() bool {
	lotusClient, err := lotus.NewClient(nil, nil)
	if err != nil {
		log.Error("Unable to initialize lotus client", err)
		return false
	}

	lastBlock, err := lotusClient.Blocks.GetLastBlock()
	if err != nil {
		log.Error("Unable to get last block")
		return false
	}
	lastBlockMiner := lastBlock.Blocks[0].Miner

	miner, err := lotusClient.Miner.GetMinerAddress()
	if err != nil {
		// TODO
	}

	if miner == lastBlockMiner {
		log.Info("I collected prize")
		// hactarClient := hactar.NewClient(session.CurrentUser.Token)
		// nodeUrl := url.GetUrl()
	}
	return true

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
