package stats

import (
	"fmt"
	"github.com/NodeFactoryIo/hactar-daemon/internal/hactar"
	"github.com/NodeFactoryIo/hactar-daemon/internal/lotus"
	"github.com/NodeFactoryIo/hactar-daemon/internal/session"
	"github.com/NodeFactoryIo/hactar-daemon/pkg/util"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"strconv"
	"time"
)

func SubmitNewBlockReport(hactarClient *hactar.Client, lotusClient *lotus.Client, currentSession session.UserSession) bool {
	// get last typset height
	lastHeight, err := lotusClient.Blocks.GetLastHeight()
	if err != nil {
		log.Error("Unable to get last typset height", err)
		return false
	}
	// get miner address
	miner, err := lotusClient.Miner.GetMinerAddress()
	if err != nil {
		log.Error("Unable to get miner address", err)
		return false
	}
	// iterate over all unchecked typsets
	for h := currentSession.GetLastCheckedHeight() + 1; h <= lastHeight; h++ {
		// get tipset for height h
		tipset, err := lotusClient.Blocks.GetTipsetByHeight(h)
		if err != nil {
			log.Error(fmt.Sprintf("Unable to get tipset of height %d", h))
			return false
		}
		// check for all blocks
		var blocks []hactar.Block
		for i, block := range tipset.Blocks {
			if miner == block.Miner {
				block := &hactar.Block{
					Cid:   tipset.Cids[i],
					Miner: block.Miner,
				}
				blocks = append(blocks, *block)
			}
		}
		// if mining reward present
		if len(blocks) > 0 {
			// send mining reward for this tipset
			response, err := hactarClient.Blocks.AddMiningReward(blocks)
			if err != nil {
				log.Error(
					fmt.Sprintf("Unable to send miner reward status for tipset of height %d", tipset.Height),
					err,
				)
				return false
			}
			if response != nil && response.StatusCode == http.StatusOK {
				log.Info(fmt.Sprintf("Miner reward for tipset of height %d sent", tipset.Height))
			}
		}
		currentSession.SetLastCheckedHeight(h)
		err = currentSession.SaveSession()
		util.Must(err, "Unable to save last block processed information.")
	}
	return true
}

func StartMonitoringBlocks(hactarClient *hactar.Client, lotusClient *lotus.Client, currentSession session.UserSession) {
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
				SubmitNewBlockReport(hactarClient, lotusClient, currentSession)
			}
		}
	}()
}
