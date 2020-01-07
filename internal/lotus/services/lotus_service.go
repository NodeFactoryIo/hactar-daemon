package services

import (
	"fmt"
	"github.com/NodeFactoryIo/hactar-daemon/internal/lotus/requests"
	"github.com/NodeFactoryIo/hactar-daemon/pkg/jsonrpc2client"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type LotusService interface {
	GetMinerAddress() string
}

func NewLotusService() *lotusService {
	return &lotusService{}
}

type lotusService struct {
	rpcClient jsonrpc2client.Client
}

func (ls *lotusService) SetJsonRpcClient(c jsonrpc2client.Client)  {
	ls.rpcClient = c
}

func (ls *lotusService) GetJsonRpcClient() jsonrpc2client.Client   {
	if ls.rpcClient != nil {
		return ls.rpcClient
	}

	url := viper.GetString("jsonrpc.url")
	ls.rpcClient = jsonrpc2client.NewClient(url)
	return ls.rpcClient
}

func (ls *lotusService) GetMinerAddress() string {
	response, err := ls.GetJsonRpcClient().Call(requests.ActorAddress)
	if err != nil {
		log.Error("Unable to get miner address", err)
	} else if response != nil && response.Error == nil {
		return fmt.Sprintf("%v", response.Result)
	}
	return "t0101" // TODO tmp
}