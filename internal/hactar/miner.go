package hactar

import (
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
	Version    string   `json:"version"`
	SectorSize int64    `json:"sectorSize"`
	MinerPower int64    `json:"minerPower"`
	TotalPower int64    `json:"totalPower"`
	Node       NodeInfo `json:"nodeInfo"`
}

func (ms *minerService) SendMinerInfo(info MinerInfo) (*http.Response, error) {
	request, err := ms.client.NewRequest(http.MethodPost, SendMinerInfoPath, info)

	if err != nil {
		return nil, err
	}

	response, err := ms.client.Do(request, nil)

	if err != nil {
		return nil, err
	}

	return response, err
}
