package minerinfo

import (
	"github.com/NodeFactoryIo/hactar-daemon/internal/hactar"
	"github.com/NodeFactoryIo/hactar-daemon/internal/lotus"
	"github.com/NodeFactoryIo/hactar-daemon/internal/url"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func SendMinerInfoStats(hactarClient *hactar.Client, lotusClient *lotus.Client) bool {
	minerAddress, err := lotusClient.Miner.GetMinerAddress()
	if err != nil {
		log.Error("Unable to get miner address ", err)
		return false
	}

	minerStats, err := lotusClient.Miner.GetMinerPower(minerAddress)
	if err != nil {
		log.Error("Unable to get miner power ", err)
		return false
	}

	sectorSize, err := lotusClient.Sectors.GetSectorSize(minerAddress)
	if err != nil {
		log.Error("Unable to get sector size ", err)
		return false
	}

	clientVersion, err := lotusClient.Client.GetClientVersion()
	if err != nil {
		log.Error("Unable to get client version ", err)
		return false
	}

	minerInfo := &hactar.MinerInfo{
		Version:    clientVersion.Version,
		SectorSize: sectorSize,
		MinerPower: minerStats.MinerPower,
		TotalPower: minerStats.TotalPower,
		Node: hactar.NodeInfo{
			Address: minerAddress,
			Url:     url.GetUrl(),
		},
	}
	response, err := hactarClient.Miner.SendMinerInfo(*minerInfo)

	if response != nil && response.StatusCode == http.StatusOK {
		log.Info("Miner info successfully sent to hactar")
		return true
	}

	log.Error("Unable to send miner information statistics ", err)
	return false
}
