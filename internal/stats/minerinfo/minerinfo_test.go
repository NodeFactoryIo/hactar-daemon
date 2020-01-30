package minerinfo

import (
	"github.com/NodeFactoryIo/hactar-daemon/internal/hactar"
	"github.com/NodeFactoryIo/hactar-daemon/internal/lotus"
	mocksHactar "github.com/NodeFactoryIo/hactar-daemon/mocks/hactar"
	mocksLotus "github.com/NodeFactoryIo/hactar-daemon/mocks/lotus"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestSendMinerInfoStats(t *testing.T) {
	lotusMinerServiceMock := &mocksLotus.MinerService{}
	lotusMinerServiceMock.On("GetMinerAddress").Return("t0101", nil)
	minerPowerResponse := &lotus.MinerPowerResponse{
		MinerPower: "100",
		TotalPower: "200",
	}
	lotusMinerServiceMock.On("GetMinerPower", "t0101").Return(minerPowerResponse, nil)

	lotusSectorServiceMock := &mocksLotus.SectorService{}
	lotusSectorServiceMock.On("GetSectorSize", "t0101").Return("12345678", nil)

	lotusClientServiceMock := &mocksLotus.ClientService{}
	clientVersionResponse := &lotus.ClientVersionResponse{
		Version:    "test-version",
		APIVersion: 123,
		BlockDelay: 123,
	}
	lotusClientServiceMock.On("GetClientVersion").Return(clientVersionResponse, nil)

	lotusMockedClient := &lotus.Client{
		Miner:   lotusMinerServiceMock,
		Blocks:  nil,
		Client:  lotusClientServiceMock,
		Sectors: lotusSectorServiceMock,
	}

	minerInfoRequest := &hactar.MinerInfo{
		Miner:      "t0101",
		Version:    "test-version",
		SectorSize: "12345678",
		MinerPower: "100",
		TotalPower: "200",
	}

	hactarMinerServiceMock := &mocksHactar.MinerService{}
	hactarMinerServiceMock.On("SendMinerInfo", *minerInfoRequest).Return(&http.Response{StatusCode: 200}, nil)

	hactarMockedClient := &hactar.Client{
		BaseURL:  nil,
		Token:    "",
		Nodes:    nil,
		DiskInfo: nil,
		Blocks:   nil,
		Miner:    hactarMinerServiceMock,
	}

	success := SendMinerInfoStats(hactarMockedClient, lotusMockedClient)

	// assertions
	assert.True(t, success)
	lotusMinerServiceMock.AssertNumberOfCalls(t, "GetMinerAddress", 1)
	lotusMinerServiceMock.AssertNumberOfCalls(t, "GetMinerPower", 1)
	lotusMinerServiceMock.AssertExpectations(t)
	lotusSectorServiceMock.AssertNumberOfCalls(t, "GetSectorSize", 1)
	lotusSectorServiceMock.AssertExpectations(t)
	lotusClientServiceMock.AssertNumberOfCalls(t, "GetClientVersion", 1)
	lotusClientServiceMock.AssertExpectations(t)
	hactarMinerServiceMock.AssertNumberOfCalls(t, "SendMinerInfo", 1)
	hactarMinerServiceMock.AssertExpectations(t)
}
