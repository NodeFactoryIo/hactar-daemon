package hactar

import (
	"github.com/getsentry/sentry-go"
	"net/http"
)

type DiskInfoService interface {
	SendDiskInfo(diskInfo DiskInfo) (*http.Response, error)
}

type diskInfoService struct {
	client *Client
}

type DiskInfo struct {
	FreeSpace  string   `json:"freeSpace"`
	TakenSpace string   `json:"takenSpace"`
	Node       NodeInfo `json:"nodeInfo"`
}

const (
	DiskInfoPath = "/user/node/diskinformation"
)

func (dis *diskInfoService) SendDiskInfo(diskInfo DiskInfo) (*http.Response, error) {
	request, err := dis.client.NewRequest(http.MethodPost, DiskInfoPath, diskInfo)

	if err != nil {
		return nil, err
	}

	response, err := dis.client.Do(request, nil)

	if err != nil {
		sentry.CaptureException(err)
		return nil, err
	}

	return response, err
}
