package lotus

import (
	"github.com/NodeFactoryIo/hactar-daemon/internal/lotus/requests/lotus"
	"github.com/NodeFactoryIo/hactar-daemon/pkg/util"
	log "github.com/sirupsen/logrus"
)

type PastDealsService interface {
	GetAllPastDeals() ([]PastDealResponse, error)
}

type pastDealsService struct {
	client *Client
}

type PastDealResponse struct {
	Cid      string `json:"cid"`
	State    int    `json:"state"`
	Size     string `json:"size"`
	Provider string `json:"provider"`
	Price    string `json:"price"`
	Duration int    `json:"duration"`
}

type rawPastDealResponse struct {
	ProposalCid   interface{} `json:"ProposalCid"`
	State         int         `json:"State"`
	Provider      string      `json:"Provider"`
	PieceRef      string      `json:"PieceRef"`
	Size          string      `json:"Size"`
	PricePerEpoch string      `json:"PricePerEpoch"`
	Duration      int         `json:"Duration"`
}

func (pds *pastDealsService) GetAllPastDeals() ([]PastDealResponse, error) {
	response, err := pds.client.lotusNodeClient.Call(lotus.PastDeals)
	if err = ValidResponse(response, err, lotus.PastDeals); err != nil {
		return nil, err
	}

	var responseObject []rawPastDealResponse
	err = response.GetObject(&responseObject)

	if err != nil {
		log.Error()
		return nil, err
	}

	return convertFromRawPastDeals(responseObject)
}

func convertFromRawPastDeals(rawPastDealResponses []rawPastDealResponse) ([]PastDealResponse, error) {
	deals := make([]PastDealResponse, len(rawPastDealResponses))
	for i := 0; i < len(rawPastDealResponses); i++ {
		deals[i] = PastDealResponse{
			Cid:      util.String(rawPastDealResponses[i].ProposalCid),
			State:    rawPastDealResponses[i].State,
			Size:     rawPastDealResponses[i].Size,
			Provider: rawPastDealResponses[i].Provider,
			Price:    rawPastDealResponses[i].PricePerEpoch,
			Duration: rawPastDealResponses[i].Duration,
		}
	}
	return deals, nil
}
