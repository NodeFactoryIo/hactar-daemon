package services

import (
	"github.com/NodeFactoryIo/hactar-daemon/internal/lotus/requests/miner"
	"github.com/NodeFactoryIo/hactar-daemon/pkg/jsonrpc2client"
	"github.com/NodeFactoryIo/hactar-daemon/pkg/util"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/ybbus/jsonrpc"
)

type LotusService interface {
	GetMinerAddress() string
}

type lotusService struct {
	lotusClient jsonrpc2client.Client
	minerClient jsonrpc2client.Client
}

func NewLotusService(lClient jsonrpc2client.Client, mClient jsonrpc2client.Client) *lotusService {
	lotusClient := lClient
	if lotusClient == nil {
		url := viper.GetString("jsonrpc.lotus.url")
		lotusClient = jsonrpc2client.NewClient(url)
	}

	minerClient := mClient
	if minerClient == nil {
		url := viper.GetString("jsonrpc.miner.url")
		minerClient = jsonrpc2client.NewClient(url)
	}

	return &lotusService{
		lotusClient: lotusClient,
		minerClient: minerClient,
	}
}

func (ls *lotusService) GetMinerAddress() string {
	response, err := ls.minerClient.Call(miner.ActorAddress)
	if err != nil {
		log.Error("Unable to get miner address", err)
		return ""
	}
	return processResult(response)
}

func processResult(response *jsonrpc.RPCResponse) string {
	if response != nil {
		if response.Error == nil {
			return util.String(response.Result)
		}
		log.Error(response.Error)
	}
	return ""
}
