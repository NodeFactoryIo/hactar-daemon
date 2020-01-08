package services

import (
	"fmt"
	"github.com/NodeFactoryIo/hactar-daemon/internal/lotus/requests"
	"github.com/NodeFactoryIo/hactar-daemon/pkg/jsonrpc2client"
	"github.com/NodeFactoryIo/hactar-daemon/pkg/util"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type LotusService interface {
	GetMinerAddress() string
}

type lotusService struct {
	rpcClient jsonrpc2client.Client
}

func NewLotusService(client jsonrpc2client.Client) *lotusService {
	jsonrpcClient := client
	if jsonrpcClient == nil {
		// if client not provided
		url := viper.Get("jsonrpc.url")
		jsonrpcClient = jsonrpc2client.NewClient(util.String(url))
	}

	return &lotusService{
		rpcClient: jsonrpcClient,
	}
}

func (ls *lotusService) SetJsonRpcClient(c jsonrpc2client.Client)  {
	ls.rpcClient = c
}

func (ls *lotusService) GetJsonRpcClient() jsonrpc2client.Client   {
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