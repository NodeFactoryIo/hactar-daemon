package lotus

import (
	"github.com/NodeFactoryIo/hactar-daemon/internal/lotus/requests/lotus"
	"github.com/NodeFactoryIo/hactar-daemon/internal/lotus/requests/miner"
	"github.com/NodeFactoryIo/hactar-daemon/pkg/util"
	log "github.com/sirupsen/logrus"
	"github.com/ybbus/jsonrpc"
)

type MinerService interface {
	GetMinerAddress() (string, error)
	GetMinerPower(miner string) (*MinerPowerResponse, error)
	GetActor(miner string, tipSetKey string) (*ActorResponse, error)
	GetLatestActor(miner string) (*ActorResponse, error)
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
	MinerPower string `json:"MinerPower"`
	TotalPower string `json:"TotalPower"`
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

type ActorResponse struct {
	Code    interface{} `json:"Code"`
	Head    interface{} `json:"Head"`
	Nonce   int64       `json:"Nonce"`
	Balance string      `json:"Balance"`
}

type TipSetKey struct {
	Cid string `json:"/"`
}

func (ms *minerService) GetActor(miner string, tipSetKey string) (*ActorResponse, error) {
	response, err := ms.client.lotusNodeClient.Call(lotus.Actor, miner, []TipSetKey{{Cid: tipSetKey}})
	return getActor(response, err)
}

func (ms *minerService) GetLatestActor(miner string) (*ActorResponse, error) {
	response, err := ms.client.lotusNodeClient.Call(lotus.Actor, miner, nil)
	return getActor(response, err)
}

func getActor(response *jsonrpc.RPCResponse, err error) (*ActorResponse, error) {
	if err = ValidResponse(response, err, lotus.Actor); err != nil {
		return nil, err
	}

	var responseObject *ActorResponse
	err = response.GetObject(&responseObject)

	if err != nil {
		log.Error()
		return nil, err
	}

	return responseObject, nil
}