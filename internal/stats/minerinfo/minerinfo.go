package minerinfo

import (
	"github.com/NodeFactoryIo/hactar-daemon/internal/hactar"
	"github.com/NodeFactoryIo/hactar-daemon/internal/lotus"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
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
	minerPower, err := strconv.ParseInt(minerStats.MinerPower, 10, 64)
	if err != nil {
		log.Error("Unable to parse miner power ", err)
		return false
	}
	totalPower, err := strconv.ParseInt(minerStats.TotalPower, 10, 64)
	if err != nil {
		log.Error("Unable to parse total power ", err)
		return false
	}

	sectorSizeString, err := lotusClient.Sectors.GetSectorSize(minerAddress)
	if err != nil {
		log.Error("Unable to get sector size ", err)
		return false
	}
	sectorSize, err := strconv.ParseInt(sectorSizeString, 10, 64)
	if err != nil {
		log.Error("Unable to parse sector size ", err)
		return false
	}

	clientVersion, err := lotusClient.Client.GetClientVersion()
	if err != nil {
		log.Error("Unable to get client version ", err)
		return false
	}

	minerInfo := &hactar.MinerInfo{
		Miner:      minerAddress,
		Version:    clientVersion.Version,
		SectorSize: sectorSize,
		MinerPower: minerPower,
		TotalPower: totalPower,
	}
	response, err := hactarClient.Miner.SendMinerInfo(*minerInfo)

	if response != nil && response.StatusCode == http.StatusOK {
		log.Info("Miner info successfully sent to hactar")
		return true
	}

	log.Error("Unable to send miner information statistics ", err)
	return false
}
