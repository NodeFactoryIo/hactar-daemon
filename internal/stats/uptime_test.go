package stats

import (
	"errors"
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

func TestSubmitNewNodeUptimeReport_NodeUp(t *testing.T) {
	lotusMinerServiceMock := &mocksLotus.MinerService{}
	lotusMinerServiceMock.On("GetMinerAddress").Return("t0101", nil)
	lotusMockedClient := &lotus.Client{
		Blocks: nil,
		Miner:  lotusMinerServiceMock,
	}

	uptimeReportRequest := &hactar.UptimeReport{
		IsWorking: true,
		Node: hactar.NodeInfo{
			Address: "t0101",
			Url:     url.GetUrl(),
		},
	}

	hactarNodeServiceMock := &mocksHactar.NodesService{}
	hactarNodeServiceMock.On("SendUptimeReport", *uptimeReportRequest).Return(&http.Response{StatusCode: 201}, nil)

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

	success := SubmitNewNodeUptimeReport(hactarMockedClient, lotusMockedClient, sessionMock)
	// assertions
	assert.True(t, success)
	lotusMinerServiceMock.AssertNumberOfCalls(t, "GetMinerAddress", 1)
	lotusMinerServiceMock.AssertExpectations(t)
	hactarNodeServiceMock.AssertNumberOfCalls(t, "SendUptimeReport", 1)
	hactarNodeServiceMock.AssertExpectations(t)
	sessionMock.AssertNumberOfCalls(t, "GetNodeMinerAddress", 1)
	sessionMock.AssertExpectations(t)
}

func TestSubmitNewNodeUptimeReport_NodeDown(t *testing.T) {
	lotusMinerServiceMock := &mocksLotus.MinerService{}
	lotusMinerServiceMock.On("GetMinerAddress").Return("", errors.New("unable to call lotus api"))
	lotusMockedClient := &lotus.Client{
		Blocks: nil,
		Miner:  lotusMinerServiceMock,
	}

	uptimeReportRequest := &hactar.UptimeReport{
		IsWorking: false,
		Node: hactar.NodeInfo{
			Address: "t0101",
			Url:     url.GetUrl(),
		},
	}

	hactarNodeServiceMock := &mocksHactar.NodesService{}
	hactarNodeServiceMock.On("SendUptimeReport", *uptimeReportRequest).Return(&http.Response{StatusCode: 201}, nil)

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

	success := SubmitNewNodeUptimeReport(hactarMockedClient, lotusMockedClient, sessionMock)
	// assertions
	assert.True(t, success)
	lotusMinerServiceMock.AssertNumberOfCalls(t, "GetMinerAddress", 1)
	lotusMinerServiceMock.AssertExpectations(t)
	hactarNodeServiceMock.AssertNumberOfCalls(t, "SendUptimeReport", 1)
	hactarNodeServiceMock.AssertExpectations(t)
	sessionMock.AssertNumberOfCalls(t, "GetNodeMinerAddress", 1)
	sessionMock.AssertExpectations(t)
}
