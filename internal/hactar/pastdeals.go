package hactar

import (
	"github.com/NodeFactoryIo/hactar-daemon/internal/lotus"
	"net/http"
)

type PastDealsService interface {
	SendPastDealsInfo(info PastDealsInfo) (*http.Response, error)
}

type pastDealsService struct {
	client *Client
}

const (
	SendPastDealsInfoPath = "/user/node/pastdeals"
)

type PastDealsInfo struct {
	Deals []lotus.PastDealResponse `json:"pastDeals"`
	Node  NodeInfo               `json:"nodeInfo"`
}

func (pds *pastDealsService) SendPastDealsInfo(info PastDealsInfo) (*http.Response, error) {
	request, err := pds.client.NewRequest(http.MethodPost, SendPastDealsInfoPath, info)

	if err != nil {
		return nil, err
	}

	response, err := pds.client.Do(request, nil)

	if err != nil {
		return nil, err
	}

	return response, err
}
