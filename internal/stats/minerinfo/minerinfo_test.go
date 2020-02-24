package minerinfo

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
	lotusSectorServiceMock.On("GetNumberOfSectors", "t0101").Return(4, nil)

	lotusClientServiceMock := &mocksLotus.ClientService{}
	clientVersionResponse := &lotus.ClientVersionResponse{
		Version:    "test-version",
		APIVersion: 123,
		BlockDelay: 123,
	}
	lotusClientServiceMock.On("GetClientVersion").Return(clientVersionResponse, nil)

	lotusWalletServiceMock := &mocksLotus.WalletService{}
	lotusWalletServiceMock.On("GetWalletDefaultAddress").Return("test-address", nil)

	lotusMockedClient := &lotus.Client{
		Miner:   lotusMinerServiceMock,
		Blocks:  nil,
		Client:  lotusClientServiceMock,
		Sectors: lotusSectorServiceMock,
		Wallet:lotusWalletServiceMock,
	}

	minerInfoRequest := &hactar.MinerInfo{
		Version:         "test-version",
		WalletAddress: "test-address",
		SectorSize:      "12345678",
		NumberOfSectors: 4,
		MinerPower:      "100",
		TotalPower:      "200",
		Node: hactar.NodeInfo{
			Address: "t0101",
			Url:     url.GetUrl(),
		},
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
	lotusSectorServiceMock.AssertNumberOfCalls(t, "GetNumberOfSectors", 1)
	lotusSectorServiceMock.AssertExpectations(t)
	lotusClientServiceMock.AssertNumberOfCalls(t, "GetClientVersion", 1)
	lotusClientServiceMock.AssertExpectations(t)
	lotusWalletServiceMock.AssertNumberOfCalls(t, "GetWalletDefaultAddress", 1)
	lotusWalletServiceMock.AssertExpectations(t)
	hactarMinerServiceMock.AssertNumberOfCalls(t, "SendMinerInfo", 1)
	hactarMinerServiceMock.AssertExpectations(t)
}
