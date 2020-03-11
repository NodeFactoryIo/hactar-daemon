package lotus

import (
	"github.com/NodeFactoryIo/hactar-daemon/internal/lotus/requests/lotus"
	"github.com/NodeFactoryIo/hactar-daemon/pkg/util"
)

type WalletService interface {
	GetWalletDefaultAddress() (string, error)
	GetWalletBalance(walletAddress string) (string, error)
	GetMinerBalance(minerAddress string, cids []string) (string, error)
}

type walletService struct {
	client *Client
}

type MinerBalanceResponse struct {
	Balance string `json:"Balance"`
}

func (ws *walletService) GetWalletDefaultAddress() (string, error) {
	response, err := ws.client.lotusNodeClient.Call(lotus.WalletDefaultAddress)
	if err = ValidResponse(response, err, lotus.WalletDefaultAddress); err != nil {
		return "", err
	}

	return util.String(response.Result), nil
}

func (ws *walletService) GetWalletBalance(walletAddress string) (string, error) {
	response, err := ws.client.lotusNodeClient.Call(lotus.WalletBalance, walletAddress)
	if err = ValidResponse(response, err, lotus.WalletBalance); err != nil {
		return "", err
	}

	return util.String(response.Result), nil
}

func (ws *walletService) GetMinerBalance(miner string, cids []string) (string, error) {
	var serializedCids []interface{}
	for i := 0; i < len(cids); i++ {
		c := [string]interface{"/": cids[i]}
		serializedCids = append(serializedCids, c)
	}

	response, err := ws.client.lotusNodeClient.Call(lotus.MinerBalance, miner, cids)
	if err = ValidResponse(response, err, lotus.MinerBalance); err != nil {
		return "", err
	}

	var responseObject *MinerBalanceResponse
	err = response.GetObject(&responseObject)

	return responseObject.Balance, nil
}
