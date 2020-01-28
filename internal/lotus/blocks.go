package lotus

import (
	"github.com/NodeFactoryIo/hactar-daemon/internal/lotus/requests/lotus"
	"github.com/NodeFactoryIo/hactar-daemon/pkg/util"
	log "github.com/sirupsen/logrus"
)

type BlocksService interface {
	GetLastTypset() (*TypsetResponse, error)
	GetLastHeight() (int64, error)
	GetTypsetByHeight(height int64) (*TypsetResponse, error)
}

type blocksService struct {
	client *Client
}

type RawTypsetResponse struct {
	Cids   []interface{} `json:"Cids"`
	Blocks []interface{} `json:"Blocks"`
	Height int64         `json:"Height"`
}

type TypsetResponse struct {
	Cids   []string
	Blocks []BlockStruct
	Height int64
}

type BlockStruct struct {
	Miner string
}

func (bs *blocksService) GetLastTypset() (*TypsetResponse, error) {
	typsetRaw, err := bs.getLastTypsetRaw()
	if err != nil {
		log.Error()
		return nil, err
	}
	return convertFromRawTypset(typsetRaw)
}

func (bs *blocksService) GetLastHeight() (int64, error) {
	typsetRaw, err := bs.getLastTypsetRaw()
	if err != nil {
		log.Error()
		return -1, err
	}
	return typsetRaw.Height, nil
}

func (bs *blocksService) GetTypsetByHeight(height int64) (*TypsetResponse, error) {
	response, err := bs.client.lotusNodeClient.Call(lotus.TipSetByHeight, height)
	if err = ValidResponse(response, err, lotus.TipSetByHeight); err != nil {
		return nil, err
	}

	var responseObject *RawTypsetResponse
	err = response.GetObject(&responseObject)

	return convertFromRawTypset(responseObject)
}

func (bs *blocksService) getLastTypsetRaw() (*RawTypsetResponse, error) {
	response, err := bs.client.lotusNodeClient.Call(lotus.HeadBlock)
	if err = ValidResponse(response, err, lotus.HeadBlock); err != nil {
		return nil, err
	}

	var responseObject *RawTypsetResponse
	err = response.GetObject(&responseObject)

	return responseObject, err
}

func convertFromRawTypset(typsetRaw *RawTypsetResponse) (*TypsetResponse, error) {
	// get all cids
	var cids []string
	for i := 0; i < len(typsetRaw.Cids); i++ {
		c := util.String(typsetRaw.Cids[i].(map[string]interface{})["/"])
		cids = append(cids, c)
	}

	// get all miners
	var miners []BlockStruct
	for i := 0; i < len(typsetRaw.Blocks); i++ {
		m := util.String(typsetRaw.Blocks[i].(map[string]interface{})["Miner"])
		miners = append(miners, *(&BlockStruct{Miner: m}))
	}

	typsetResponse := &TypsetResponse{
		Cids:   cids,
		Blocks: miners,
		Height: typsetRaw.Height,
	}

	return typsetResponse, nil
}
