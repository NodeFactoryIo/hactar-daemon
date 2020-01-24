package lotus

import (
	"github.com/NodeFactoryIo/hactar-daemon/internal/lotus/requests/miner"
	"github.com/NodeFactoryIo/hactar-daemon/pkg/util"
)

type MinerService interface {
	GetMinerAddress() (string, error)
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
