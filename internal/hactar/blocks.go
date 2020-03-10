package hactar

import (
	"github.com/getsentry/sentry-go"
	"net/http"
)

type BlocksService interface {
	AddMiningReward(blocks []Block) (*http.Response, error)
}

type blocksService struct {
	client *Client
}

const (
	AddBlockRewardPath = "/user/node/miningrewards"
)

type Block struct {
	Cid    string   `json:"cid"`
	Reward string   `json:"reward"`
	Node   NodeInfo `json:"nodeInfo"`
}

func (bs *blocksService) AddMiningReward(blocks []Block) (*http.Response, error) {
	request, err := bs.client.NewRequest(http.MethodPost, AddBlockRewardPath, blocks)

	if err != nil {
		return nil, err
	}

	response, err := bs.client.Do(request, nil)

	if err != nil {
		sentry.CaptureException(err)
		return nil, err
	}

	return response, err
}
