package minerinfo

import (
	"github.com/NodeFactoryIo/hactar-daemon/internal/hactar"
	"github.com/NodeFactoryIo/hactar-daemon/internal/lotus"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func SendMinerInfoStats(hactarClient *hactar.Client, lotusClient *lotus.Client) bool {
	minerAddress, err := lotusClient.Miner.GetMinerAddress()
	if err != nil {
		// TODO
		return false
	}

	minerStats, err := lotusClient.Miner.GetMinerPower(minerAddress)
	if err != nil {
		// TODO
		return false
	}

	sectorSize, err := lotusClient.Sectors.GetSectorSize(minerAddress)
	if err != nil {
		// TODO
		return false
	}

	clientVersion, err := lotusClient.Client.GetClientVersion()
	if err != nil {
		// TODO
		return false
	}

	minerInfo := &hactar.MinerInfo{
		Miner:      minerAddress,
		Version:    clientVersion.Version,
		SectorSize: sectorSize,
		MinerPower: minerStats.MinerPower,
		TotalPower: minerStats.TotalPower,
	}
	response, err := hactarClient.Miner.SendMinerInfo(*minerInfo)

	if response != nil && response.StatusCode == http.StatusOK {
		log.Info()
		return true
	}

	log.Error("Unable to send miner information statistics. ", err)
	return false
}
