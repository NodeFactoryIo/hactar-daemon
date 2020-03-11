package hactar

import (
	"github.com/getsentry/sentry-go"
	"net/http"
)

type MinerService interface {
	SendMinerInfo(info MinerInfo) (*http.Response, error)
}

type minerService struct {
	client *Client
}

const (
	SendMinerInfoPath = "/user/node/generalminerinfo"
)

type MinerInfo struct {
	Version         string   `json:"version"`
	WalletAddress   string   `json:"walletAddress"`
	SectorSize      string   `json:"sectorSize"`
	NumberOfSectors int      `json:"numberOfSectors"`
	MinerPower      string   `json:"minerPower"`
	TotalPower      string   `json:"totalPower"`
	Node            NodeInfo `json:"nodeInfo"`
}

func (ms *minerService) SendMinerInfo(info MinerInfo) (*http.Response, error) {
	request, err := ms.client.NewRequest(http.MethodPut, SendMinerInfoPath, info)

	if err != nil {
		return nil, err
	}

	response, err := ms.client.Do(request, nil)

	if err != nil {
		sentry.CaptureException(err)
		return nil, err
	}

	return response, err
}
