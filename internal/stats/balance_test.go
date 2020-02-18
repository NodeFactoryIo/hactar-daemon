package stats

import (
	"github.com/NodeFactoryIo/hactar-daemon/internal/hactar"
	"github.com/NodeFactoryIo/hactar-daemon/internal/lotus"
	"github.com/NodeFactoryIo/hactar-daemon/internal/url"
	mocksHactar "github.com/NodeFactoryIo/hactar-daemon/mocks/hactar"
	mocksLotus "github.com/NodeFactoryIo/hactar-daemon/mocks/lotus"
	mocks "github.com/NodeFactoryIo/hactar-daemon/mocks/session"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestSubmitBalanceReport(t *testing.T) {
	lotusWalletServiceMock := &mocksLotus.WalletService{}
	lotusWalletServiceMock.On("GetWalletDefaultAddress").Return("1234", nil)
	lotusWalletServiceMock.On("GetWalletBalance", "1234").Return("100200", nil)
	lotusMockedClient := &lotus.Client{
		Wallet: lotusWalletServiceMock,
	}

	balanceReportRequest := &hactar.BalanceReport{
		Balance: "100200",
		Node: hactar.NodeInfo{
			Address: "t0101",
			Url:     url.GetUrl(),
		},
	}

	hactarNodeServiceMock := &mocksHactar.NodesService{}
	hactarNodeServiceMock.On("SendBalanceReport", *balanceReportRequest).Return(&http.Response{StatusCode: 201}, nil)

	hactarMockedClient := &hactar.Client{
		BaseURL:  nil,
		Token:    "",
		Nodes:    hactarNodeServiceMock,
		DiskInfo: nil,
		Blocks:   nil,
		Miner:    nil,
	}

	sessionMock := new(mocks.UserSession)
	sessionMock.On("GetNodeMinerAddress").Return("t0101")

	success := SubmitNewBalanceReport(hactarMockedClient, lotusMockedClient, sessionMock)
	// assertions
	assert.True(t, success)
	lotusWalletServiceMock.AssertNumberOfCalls(t, "GetWalletDefaultAddress", 1)
	lotusWalletServiceMock.AssertNumberOfCalls(t, "GetWalletBalance", 1)
	lotusWalletServiceMock.AssertExpectations(t)
	hactarNodeServiceMock.AssertNumberOfCalls(t, "SendBalanceReport", 1)
	hactarNodeServiceMock.AssertExpectations(t)
}
