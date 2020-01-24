package stats

import (
	"fmt"
	"github.com/NodeFactoryIo/hactar-daemon/internal/hactar"
	"github.com/NodeFactoryIo/hactar-daemon/internal/lotus"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"strconv"
	"time"
)

func SubmitNewBlockReport(hactarClient *hactar.Client, lotusClient *lotus.Client) bool {
	lastBlock, err := lotusClient.Blocks.GetLastBlock()
	if err != nil {
		log.Error("Unable to get last block", err)
		return false
	}
	lastBlockMiner := lastBlock.Blocks[0].Miner

	miner, err := lotusClient.Miner.GetMinerAddress()
	if err != nil {
		log.Error("Unable to get miner address", err)
		return false
	}

	// if miner finished last block
	if miner == lastBlockMiner {
		lastBlockCid := lastBlock.Cids[0].Root
		block := &hactar.Block{
			Cid:   lastBlockCid,
			Miner: miner,
		}
		response, err := hactarClient.Blocks.AddMiningReward(*block)

		if response != nil && response.StatusCode == http.StatusOK {
			log.Info(fmt.Sprintf("Miner reward for block %s sent", lastBlockCid))
			return true
		}

		log.Error(fmt.Sprintf("Unable to send miner reward status for block %s", lastBlockCid), err)
		return false
	}
	return true
}

func StartMonitoringBlocks(hactarClient *hactar.Client, lotusClient *lotus.Client) {
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
				SubmitNewBlockReport(hactarClient, lotusClient)
			}
		}
	}()
}
