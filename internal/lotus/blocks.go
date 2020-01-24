package lotus

import (
	"github.com/NodeFactoryIo/hactar-daemon/internal/lotus/requests/lotus"
	log "github.com/sirupsen/logrus"
)

type BlocksService interface {
	GetLastBlock() (*LastBlockResponse, error)
}

type blocksService struct {
	client *Client
}

type LastBlockResponse struct {
	Cids []struct {
		Root string `json:"/"`
	} `json:"Cids"`
	Blocks []struct {
		Miner  string `json:"Miner"`
		Ticket struct {
			VRFProof string `json:"VRFProof"`
		} `json:"Ticket"`
		EPostProof struct {
			Proof      string `json:"Proof"`
			PostRand   string `json:"PostRand"`
			Candidates []struct {
				Partial        string `json:"Partial"`
				SectorID       int    `json:"SectorID"`
				ChallengeIndex int    `json:"ChallengeIndex"`
			} `json:"Candidates"`
		} `json:"EPostProof"`
		Parents []struct {
			Root string `json:"/"`
		} `json:"Parents"`
		ParentWeight    string `json:"ParentWeight"`
		Height          int    `json:"Height"`
		ParentStateRoot struct {
			Root string `json:"/"`
		} `json:"ParentStateRoot"`
		ParentMessageReceipts struct {
			Root string `json:"/"`
		} `json:"ParentMessageReceipts"`
		Messages struct {
			Root string `json:"/"`
		} `json:"Messages"`
		BLSAggregate struct {
			Type string `json:"Type"`
			Data string `json:"Data"`
		} `json:"BLSAggregate"`
		Timestamp int `json:"Timestamp"`
		BlockSig  struct {
			Type string `json:"Type"`
			Data string `json:"Data"`
		} `json:"BlockSig"`
	} `json:"Blocks"`
	Height int `json:"Height"`
}

func (bs *blocksService) GetLastBlock() (*LastBlockResponse, error) {
	response, err := bs.client.lotusNodeClient.Call(lotus.HeadBlock)
	if err = ValidResponse(response, err, lotus.HeadBlock); err != nil {
		return nil, err
	}

	var responseObject *LastBlockResponse
	err = response.GetObject(&responseObject)

	if err != nil {
		log.Error("Unable to parse response for rpc method: ", lotus.HeadBlock, err)
		return nil, err
	}

	return responseObject, nil
}
