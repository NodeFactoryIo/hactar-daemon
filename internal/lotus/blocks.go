package lotus

import (
	"github.com/NodeFactoryIo/hactar-daemon/internal/lotus/requests/lotus"
	"github.com/NodeFactoryIo/hactar-daemon/pkg/util"
	"github.com/getsentry/sentry-go"
)

type BlocksService interface {
	GetLastTipset() (*TipsetResponse, error)
	GetLastHeight() (int64, error)
	GetTipsetByHeight(height int64) (*TipsetResponse, error)
}

type blocksService struct {
	client *Client
}

type TipsetResponse struct {
	Cids   []string
	Blocks []BlockStruct
	Height int64
}

type BlockStruct struct {
	Miner string
}

type rawTipsetResponse struct {
	Cids   []interface{} `json:"Cids"`
	Blocks []interface{} `json:"Blocks"`
	Height int64         `json:"Height"`
}

func (bs *blocksService) GetLastTipset() (*TipsetResponse, error) {
	tipsetRaw, err := bs.getLastTipsetRaw()
	if err != nil {
		sentry.CaptureException(err)
		return nil, err
	}
	return convertFromRawTipset(tipsetRaw)
}

func (bs *blocksService) GetLastHeight() (int64, error) {
	tipsetRaw, err := bs.getLastTipsetRaw()
	if err != nil {
		sentry.CaptureException(err)
		return -1, err
	}
	return tipsetRaw.Height, nil
}

func (bs *blocksService) GetTipsetByHeight(height int64) (*TipsetResponse, error) {
	response, err := bs.client.lotusNodeClient.Call(lotus.TipSetByHeight, height, nil)
	if err = ValidResponse(response, err, lotus.TipSetByHeight); err != nil {
		return nil, err
	}

	var responseObject *rawTipsetResponse
	err = response.GetObject(&responseObject)

	return convertFromRawTipset(responseObject)
}

func (bs *blocksService) getLastTipsetRaw() (*rawTipsetResponse, error) {
	response, err := bs.client.lotusNodeClient.Call(lotus.HeadBlock)
	if err = ValidResponse(response, err, lotus.HeadBlock); err != nil {
		return nil, err
	}

	var responseObject *rawTipsetResponse
	err = response.GetObject(&responseObject)

	return responseObject, err
}

func convertFromRawTipset(tipsetResponse *rawTipsetResponse) (*TipsetResponse, error) {
	// get all cids
	var cids []string
	for i := 0; i < len(tipsetResponse.Cids); i++ {
		c := util.String(tipsetResponse.Cids[i].(map[string]interface{})["/"])
		cids = append(cids, c)
	}
	// get all miners
	var miners []BlockStruct
	for i := 0; i < len(tipsetResponse.Blocks); i++ {
		m := util.String(tipsetResponse.Blocks[i].(map[string]interface{})["Miner"])
		miners = append(miners, *(&BlockStruct{Miner: m}))
	}
	// create TipsetResponse
	response := &TipsetResponse{
		Cids:   cids,
		Blocks: miners,
		Height: tipsetResponse.Height,
	}
	return response, nil
}
