package lotus

import (
	"github.com/NodeFactoryIo/hactar-daemon/internal/lotus/requests/lotus"
	log "github.com/sirupsen/logrus"
	"strconv"
)

type SectorService interface {
	GetSectorSize(miner string) (string, error)
	GetNumberOfSectors(miner string) (int, error)
}

type sectorService struct {
	client *Client
}

type MinerInfoResponse struct {
	Owner                      string      `json:"Owner"`
	Worker                     string      `json:"Worker"`
	NewWorker                  string      `json:"NewWorker"`
	ControlAddresses           interface{} `json:"ControlAddresses"`
	WorkerChangeEpoch          int         `json:"WorkerChangeEpoch"`
	PeerId                     string      `json:"PeerId"`
	Multiaddrs                 interface{} `json:"Multiaddrs"`
	SealProofType              int         `json:"SealProofType"`
	SectorSize                 int         `json:"SectorSize"`
	WindowPoStPartitionSectors int         `json:"WindowPoStPartitionSectors"`
	ConsensusFaultElapsed      int         `json:"ConsensusFaultElapsed"`
}

func (ss *sectorService) GetSectorSize(miner string) (string, error) {
	response, err := ss.client.lotusNodeClient.Call(lotus.MinerInfo, miner, nil)
	if err = ValidResponse(response, err, lotus.MinerInfo); err != nil {
		return "", err
	}

	var responseObject *MinerInfoResponse
	err = response.GetObject(&responseObject)

	if err != nil {
		log.Error("Unable to parse response for rpc method: ", lotus.MinerInfo, err)
		return "", err
	}

	return strconv.Itoa(responseObject.SectorSize), nil
}

func (ss *sectorService) GetNumberOfSectors(miner string) (int, error) {
	response, err := ss.client.lotusNodeClient.Call(lotus.Sectors, miner, nil, nil)
	if err = ValidResponse(response, err, lotus.Sectors); err != nil {
		return -1, err
	}

	sectors, _ := response.Result.([]interface{})
	return len(sectors), nil
}
