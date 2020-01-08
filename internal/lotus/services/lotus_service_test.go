package services_test

import (
	"github.com/NodeFactoryIo/hactar-daemon/internal/lotus/services"
	"github.com/NodeFactoryIo/hactar-daemon/pkg/jsonrpc2client/mock"
	"testing"
)

// TODO write tests
func TestServiceInitialization(t *testing.T) {
	lotusMockClient := mock.NewClient()
	minerMockClient := mock.NewClient()
	services.NewLotusService(lotusMockClient, minerMockClient)
}