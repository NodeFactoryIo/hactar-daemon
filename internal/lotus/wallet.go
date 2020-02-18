package lotus

import (
	"github.com/NodeFactoryIo/hactar-daemon/internal/lotus/requests/lotus"
	"github.com/NodeFactoryIo/hactar-daemon/pkg/util"
)

type WalletService interface {
	GetWalletDefaultAddress() (string, error)
	GetWalletBalance(walletAddress string) (string, error)
}

type walletService struct {
	client *Client
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
