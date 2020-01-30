package lotus

import (
	"github.com/NodeFactoryIo/hactar-daemon/internal/lotus/requests/lotus"
	"github.com/NodeFactoryIo/hactar-daemon/internal/lotus/requests/miner"
	"github.com/NodeFactoryIo/hactar-daemon/pkg/util"
	log "github.com/sirupsen/logrus"
)

type MinerService interface {
	GetMinerAddress() (string, error)
	GetMinerPower(miner string) (*MinerPowerResponse, error)
}

type minerService struct {
	client *Client
}

func (ms *minerService) GetMinerAddress() (string, error) {
	response, err := ms.client.lotusMinerClient.Call(miner.ActorAddress)
	if err = ValidResponse(response, err, miner.ActorAddress); err != nil {
		return "", err
	}

	return util.String(response.Result), nil
}

type MinerPowerResponse struct {
	MinerPower int64 `json:"MinerPower"`
	TotalPower int64 `json:"TotalPower"`
}

func (ms *minerService) GetMinerPower(miner string) (*MinerPowerResponse, error) {
	response, err := ms.client.lotusNodeClient.Call(lotus.MinerPower, miner, nil)
	if err = ValidResponse(response, err, lotus.MinerPower); err != nil {
		return nil, err
	}

	var responseObject *MinerPowerResponse
	err = response.GetObject(&responseObject)

	if err != nil {
		log.Error("Unable to parse response for rpc method: ", lotus.MinerPower, err)
		return nil, err
	}

	return responseObject, nil
}
