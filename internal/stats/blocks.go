package stats

import (
	"fmt"
	"github.com/NodeFactoryIo/hactar-daemon/internal/hactar"
	"github.com/NodeFactoryIo/hactar-daemon/internal/lotus"
	"github.com/NodeFactoryIo/hactar-daemon/internal/session"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"strconv"
	"time"
)

func SubmitNewBlockReport(hactarClient *hactar.Client, lotusClient *lotus.Client) bool {
	lastHeight, err := lotusClient.Blocks.GetLastHeight()
	if err != nil {
		log.Error("Unable to get last typset", err)
		return false
	}

	// for all unchecked typsets
	for h := session.CurrentUser.LastCheckedHeight; h < lastHeight; h++ {
		// get miner address
		miner, err := lotusClient.Miner.GetMinerAddress()
		if err != nil {
			log.Error("Unable to get miner address", err)
			return false
		}
		typset, err := lotusClient.Blocks.GetTypsetByHeight(h)
		// save block as processed
		beforeLastCheckedHeight := session.CurrentUser.LastCheckedHeight
		session.CurrentUser.LastCheckedHeight = h
		err = session.SaveSession()
		if err != nil {
			log.Error("Unable to save last block processed information.", err)
			return false
		}
		for i, block := range typset.Blocks {
			if miner == block.Miner {
				block := &hactar.Block{
					Cid:   typset.Cids[i],
					Miner: block.Miner,
				}
				response, err := hactarClient.Blocks.AddMiningReward(*block)
				if response != nil && response.StatusCode == http.StatusOK {
					log.Info(fmt.Sprintf("Miner reward for block %s sent", typset.Cids[i]))
					return true
				}
				log.Error(fmt.Sprintf("Unable to send miner reward status for block %s", typset.Cids[i]), err)
				session.CurrentUser.LastCheckedHeight = beforeLastCheckedHeight
				_ = session.SaveSession()
				return false
			}
		}

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
