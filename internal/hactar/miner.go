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
	SendMinerInfoPath = "/user/miner"
)

type MinerInfo struct {
	Miner      string `json:"miner"`
	Version    string `json:"version"`
	SectorSize string `json:"sector_size"`
	MinerPower string `json:"miner_power"`
	TotalPower string `json:"total_power"`
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
