package stats

import (
	"github.com/NodeFactoryIo/hactar-daemon/internal/hactar"
	"github.com/NodeFactoryIo/hactar-daemon/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"testing"
)

func TestSendDiskInfoStats_OkResponse(t *testing.T) {
	// mock disk info service
	diskInfoServiceMock := &mocks.DiskInfoService{}
	diskInfoServiceMock.On("SendDiskInfo", mock.Anything).Return(&http.Response{StatusCode: 200}, nil)
	// create hactar client with mocked service
	mockedClient := &hactar.Client{
		BaseURL:  nil,
		Nodes:    nil,
		DiskInfo: diskInfoServiceMock,
	}
	assert.True(t, SendDiskInfoStats(mockedClient, "", ""))
	assert.True(t, diskInfoServiceMock.AssertNumberOfCalls(t, "SendDiskInfo", 1))
}

func TestSendDiskInfoStats_ErrorResponse(t *testing.T) {
	// mock disk info service
	diskInfoServiceMock := &mocks.DiskInfoService{}
	diskInfoServiceMock.On("SendDiskInfo", mock.Anything).Return(&http.Response{StatusCode: 400}, nil)

	// create hactar client with mocked service
	mockedClient := &hactar.Client{
		BaseURL:  nil,
		Nodes:    nil,
		DiskInfo: diskInfoServiceMock,
	}

	assert.False(t, SendDiskInfoStats(mockedClient, "", ""))
	assert.True(t, diskInfoServiceMock.AssertNumberOfCalls(t, "SendDiskInfo", 1))
}

func TestDiskUsage(t *testing.T) {
	usage := DiskUsage("/")
	assert.NotNil(t, usage.Used)
	assert.NotNil(t, usage.Free)
	assert.NotNil(t, usage.All)
	assert.True(t, usage.All > 0)
}
