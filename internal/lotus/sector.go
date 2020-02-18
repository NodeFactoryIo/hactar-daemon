package lotus

import (
	"github.com/NodeFactoryIo/hactar-daemon/internal/lotus/requests/lotus"
	"github.com/NodeFactoryIo/hactar-daemon/pkg/util"
)

type SectorService interface {
	GetSectorSize(miner string) (string, error)
	GetNumberOfSectors(miner string) (int, error)
}

type sectorService struct {
	client *Client
}

func (ss *sectorService) GetSectorSize(miner string) (string, error) {
	response, err := ss.client.lotusNodeClient.Call(lotus.SectorSize, miner, nil)
	if err = ValidResponse(response, err, lotus.SectorSize); err != nil {
		return "", err
	}

	return util.String(response.Result), nil
}

func (ss *sectorService) GetNumberOfSectors(miner string) (int, error) {
	response, err := ss.client.lotusNodeClient.Call(lotus.Sectors, miner, nil)
	if err = ValidResponse(response, err, lotus.Sectors); err != nil {
		return -1, err
	}

	sectors, _ := response.Result.([]interface{})
	return len(sectors), nil
}
