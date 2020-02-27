package lotus

import (
	"errors"
	"github.com/NodeFactoryIo/hactar-daemon/internal/token"
	"github.com/NodeFactoryIo/hactar-daemon/pkg/jsonrpc2client"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/ybbus/jsonrpc"
)

type Client struct {
	lotusNodeClient  jsonrpc2client.Client
	lotusMinerClient jsonrpc2client.Client
	Blocks           BlocksService
	Miner            MinerService
	Sectors          SectorService
	Client           ClientService
	Wallet           WalletService
	PastDeals        PastDealsService
}

func NewClient(lClient jsonrpc2client.Client, mClient jsonrpc2client.Client) (*Client, error) {
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

	client := &Client{
		lotusNodeClient:  lotusClient,
		lotusMinerClient: minerClient,
	}

	client.Blocks = &blocksService{client: client}
	client.Miner = &minerService{client: client}
	client.Sectors = &sectorService{client: client}
	client.Client = &clientService{client: client}
	client.Wallet = &walletService{client: client}
	client.PastDeals = &pastDealsService{client: client}

	return client, nil
}

func initClient(url, token string) (jsonrpc2client.Client, error) {
	if token == "" || url == "" {
		return nil, errors.New("unable to initialize rpc client")
	}
	return jsonrpc2client.NewClient(url, token), nil
}

func ValidResponse(response *jsonrpc.RPCResponse, err error, method string) error {
	if err != nil {
		log.Error("Failed to call rpc method: ", method, err)
		return err
	}
	if response != nil {
		if response.Error != nil {
			log.Error(response.Error)
			return errors.New(response.Error.Message)
		} else {
			return nil
		}
	}
	log.Error("Failed rpc call")
	return errors.New("failed rpc call")
}
