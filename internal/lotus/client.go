package lotus

import (
	"github.com/NodeFactoryIo/hactar-daemon/internal/lotus/requests/lotus"
	log "github.com/sirupsen/logrus"
)

type ClientService interface {
	GetClientVersion() (*ClientVersionResponse, error)
}

type clientService struct {
	client *Client
}

type ClientVersionResponse struct {
	Version    string `json:"Version"`
	APIVersion int    `json:"APIVersion"`
	BlockDelay int    `json:"BlockDelay"`
}

func (cs *clientService) GetClientVersion() (*ClientVersionResponse, error) {
	response, err := cs.client.lotusNodeClient.Call(lotus.Version)
	if err = ValidResponse(response, err, lotus.Version); err != nil {
		return nil, err
	}

	var responseObject *ClientVersionResponse
	err = response.GetObject(&responseObject)

	if err != nil {
		log.Error("Unable to parse response for rpc method: ", lotus.Version, err)
		return nil, err
	}

	return responseObject, nil
}
