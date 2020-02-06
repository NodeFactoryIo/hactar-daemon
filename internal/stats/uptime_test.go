package stats

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

func TestSubmitNewNodeUptimeReport(t *testing.T) {
	lotusMinerServiceMock := &mocksLotus.MinerService{}
	lotusMinerServiceMock.On("GetMinerAddress").Return("t0101", nil)
	lotusMockedClient := &lotus.Client{
		Blocks: nil,
		Miner:  lotusMinerServiceMock,
	}
	
	uptimeReportRequest := &hactar.UptimeReport{
		IsWorking: true,
		Node:      hactar.NodeInfo{
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

	success := SubmitNewNodeUptimeReport(hactarMockedClient, lotusMockedClient, nil)
	// assertions
	assert.True(t, success)
	lotusMinerServiceMock.AssertNumberOfCalls(t, "GetMinerAddress", 1)
	lotusMinerServiceMock.AssertExpectations(t)
	hactarNodeServiceMock.AssertNumberOfCalls(t, "SendUptimeReport", 1)
	hactarNodeServiceMock.AssertExpectations(t)
}