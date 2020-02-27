package pastdealsinfo

import (
	"github.com/NodeFactoryIo/hactar-daemon/internal/hactar"
	"github.com/NodeFactoryIo/hactar-daemon/internal/lotus"
	"github.com/NodeFactoryIo/hactar-daemon/internal/url"
	mocksHactar "github.com/NodeFactoryIo/hactar-daemon/mocks/hactar"
	mocksLotus "github.com/NodeFactoryIo/hactar-daemon/mocks/lotus"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestSendPastDealsInfo(t *testing.T) {
	lotusMinerServiceMock := &mocksLotus.MinerService{}
	lotusMinerServiceMock.On("GetMinerAddress").Return("t0101", nil)

	pastDealsResponse := []lotus.PastDealResponse{
		{
			Cid:      "t2df1g3",
			State:    5,
			Size:     "1234421",
			Provider: "test-provider",
			Price:    "213213",
			Duration: 5,
		},
		{
			Cid:      "4hj4fss",
			State:    2,
			Size:     "4284111",
			Provider: "test-provider-2",
			Price:    "185041",
			Duration: 4,
		},
	}

	lotusPastDealsServiceMock := &mocksLotus.PastDealsService{}
	lotusPastDealsServiceMock.On("GetAllPastDeals").Return(pastDealsResponse, nil)

	lotusMockedClient := &lotus.Client{
		Miner:     lotusMinerServiceMock,
		PastDeals: lotusPastDealsServiceMock,
	}

	pastDealsInfoRequest := hactar.PastDealsInfo{
		Deals: pastDealsResponse,
		Node: hactar.NodeInfo{
			Address: "t0101",
			Url:     url.GetUrl(),
		},
	}

	hactarPastDealsServiceMock := &mocksHactar.PastDealsService{}
	hactarPastDealsServiceMock.On("SendPastDealsInfo", pastDealsInfoRequest).Return(&http.Response{StatusCode: 200}, nil)

	hactarMockedClient := &hactar.Client{
		Token:     "",
		PastDeals: hactarPastDealsServiceMock,
	}

	// assertions
	success := SendPastDealsInfo(hactarMockedClient, lotusMockedClient)
	assert.True(t, success)
	lotusMinerServiceMock.AssertNumberOfCalls(t, "GetMinerAddress", 1)
	lotusMinerServiceMock.AssertExpectations(t)
	lotusPastDealsServiceMock.AssertNumberOfCalls(t, "GetAllPastDeals", 1)
	lotusPastDealsServiceMock.AssertExpectations(t)
	hactarPastDealsServiceMock.AssertNumberOfCalls(t, "SendPastDealsInfo", 1)
	hactarPastDealsServiceMock.AssertExpectations(t)
}
