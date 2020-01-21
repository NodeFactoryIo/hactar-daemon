package services_test

import (
	"errors"
	"github.com/NodeFactoryIo/hactar-daemon/internal/lotus/requests/miner"
	"github.com/NodeFactoryIo/hactar-daemon/internal/lotus/services"
	"github.com/NodeFactoryIo/hactar-daemon/pkg/jsonrpc2client/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLotusService_GetMinerAddress_ValidResponse(t *testing.T) {
	lotusMockClient := mock.NewClient()
	minerMockClient := mock.NewClient()

	minerMockClient.MockResponse(miner.ActorAddress, "t0101")

	lotusService, _ := services.NewLotusService(lotusMockClient, minerMockClient)

	address, err := lotusService.GetMinerAddress()

	assert.Equal(t, "t0101", address)
	assert.Nil(t, err)
}
func TestLotusService_GetMinerAddress_ErrorResponse(t *testing.T) {
	lotusMockClient := mock.NewClient()
	minerMockClient := mock.NewClient()

	minerMockClient.MockError(miner.ActorAddress, errors.New("test-error"))

	lotusService, _ := services.NewLotusService(lotusMockClient, minerMockClient)

	address, err := lotusService.GetMinerAddress()

	assert.Equal(t, errors.New("test-error"), err)
	assert.Equal(t, "", address)
}
