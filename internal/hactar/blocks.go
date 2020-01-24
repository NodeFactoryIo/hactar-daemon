package hactar

import "net/http"

type BlocksService interface {
	AddMiningReward(block Block) (*http.Response, error)
}

type blocksService struct {
	client *Client
}

const (
	AddBlockRewardPath = "/user/block/reward"
)

type Block struct {
	Cid string `json:"cid"`
}

func (bs *blocksService) AddMiningReward(block Block) (*http.Response, error)  {
	request, err := bs.client.NewRequest(http.MethodPost, AddBlockRewardPath, block)

	if err != nil {
		return nil, err
	}

	response, err := bs.client.Do(request, nil)

	if err != nil {
		return nil, err
	}

	return response, err
}