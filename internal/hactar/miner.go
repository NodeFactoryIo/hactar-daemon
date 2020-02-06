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
	SectorSize string   `json:"sectorSize"`
	MinerPower string   `json:"minerPower"`
	TotalPower string   `json:"totalPower"`
	Node       NodeInfo `json:"nodeInfo"`
}

func (ms *minerService) SendMinerInfo(info MinerInfo) (*http.Response, error) {
	request, err := ms.client.NewRequest(http.MethodPut, SendMinerInfoPath, info)

	if err != nil {
		return nil, err
	}

	response, err := ms.client.Do(request, nil)

	if err != nil {
		return nil, err
	}

	return response, err
}
