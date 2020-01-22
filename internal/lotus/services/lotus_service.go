package services

import (
	"errors"
	"github.com/NodeFactoryIo/hactar-daemon/internal/lotus/requests/lotus"
	"github.com/NodeFactoryIo/hactar-daemon/internal/lotus/requests/miner"
	"github.com/NodeFactoryIo/hactar-daemon/internal/token"
	"github.com/NodeFactoryIo/hactar-daemon/pkg/jsonrpc2client"
	"github.com/NodeFactoryIo/hactar-daemon/pkg/util"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/ybbus/jsonrpc"
)

type LotusService interface {
	GetMinerAddress() (string, error)
}

type lotusService struct {
	lotusClient jsonrpc2client.Client
	minerClient jsonrpc2client.Client
}

func initClient(url, token string) (jsonrpc2client.Client, error) {
	if token == "" || url == ""{
		return nil, errors.New("unable to initialize rpc client")
	}
	return jsonrpc2client.NewClient(url, token), nil
}

func NewLotusService(lClient jsonrpc2client.Client, mClient jsonrpc2client.Client) (*lotusService, error) {
	lotusClient := lClient
	if lotusClient == nil {
		c, err := initClient(viper.GetString("jsonrpc.lotus-node.url"), token.ReadNodeTokenFromFile())
		if err != nil {
			// unable to initialize lotus node client
			return nil, err
		}
		lotusClient = c
	}

	minerClient := mClient
	if minerClient == nil {
		c, err := initClient(viper.GetString("jsonrpc.lotus-miner.url"), token.ReadMinerTokenFromFile())
		if err != nil {
			// unable to initialize lotus node client
			return nil, err
		}
		minerClient = c
	}

	return &lotusService{
		lotusClient: lotusClient,
		minerClient: minerClient,
	}, nil
}

func (ls *lotusService) GetMinerAddress() (string, error) {
	response, err := ls.minerClient.Call(miner.ActorAddress)
	if err != nil {
		log.Error("Unable to get miner address", err)
		return "", err
	}
	return processResult(response)
}

type res struct {
	Cids struct {
		 string
	}
}

func (ls *lotusService) GetLastBlock() (string, error)  {
	response, err := ls.lotusClient.Call(lotus.HeadBlock)
	if err != nil {
		log.Error("Unable to get last block details", err)
		return "", err
	}
	// mapstructure.Decode(response.Result, )
	return processResult(response)
}

func processResult(response *jsonrpc.RPCResponse) (string, error) {
	if response != nil {
		if response.Error == nil {
			return util.String(response.Result), nil
		}
		log.Error(response.Error)
		return "", errors.New(response.Error.Message)
	}
	return "", errors.New("unable to process rpc response")
}
