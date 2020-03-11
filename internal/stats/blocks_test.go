package stats

import (
	"github.com/NodeFactoryIo/hactar-daemon/internal/hactar"
	"github.com/NodeFactoryIo/hactar-daemon/internal/lotus"
	"github.com/NodeFactoryIo/hactar-daemon/internal/url"
	mocksHactar "github.com/NodeFactoryIo/hactar-daemon/mocks/hactar"
	mocksLotus "github.com/NodeFactoryIo/hactar-daemon/mocks/lotus"
	mocks "github.com/NodeFactoryIo/hactar-daemon/mocks/session"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"testing"
)

func TestSubmitNewBlockReport_OneValidTipsetWithReward_Success(t *testing.T) {
	hactarBlockServiceMock := &mocksHactar.BlocksService{}
	hactarBlockServiceMock.On("AddMiningReward", []hactar.Block{
		*(&hactar.Block{
			Cid:    "bafy2bzaceawp7zcx74biecfl3axvoulh4lgdnnwzvgaza2cdhmxx75ymo6jnk",
			Reward: "100",
			Node: hactar.NodeInfo{
				Address: "t0101",
				Url:     url.GetUrl(),
			},
		}),
	}).Return(&http.Response{StatusCode: 200}, nil)
	hactarMockedClient := &hactar.Client{
		BaseURL:  nil,
		Token:    "",
		Nodes:    nil,
		DiskInfo: nil,
		Blocks:   hactarBlockServiceMock,
	}

	lotusBlockServiceMock := &mocksLotus.BlocksService{}
	lotusBlockServiceMock.On("GetLastHeight").Return(int64(2), nil)
	tipsetResponse := &lotus.TipsetResponse{
		Cids: []string{
			"bafy2bzaceawp7zcx74biecfl3axvoulh4lgdnnwzvgaza2cdhmxx75ymo6jnk",
			"bafy2bzaceawp7zcx74biecfl3axvoulh4lgdnnwzvgaza2cdhmxx75ymo7jnk",
		},
		Blocks: []lotus.BlockStruct{
			*(&lotus.BlockStruct{Miner: "t0101"}),
			*(&lotus.BlockStruct{Miner: "t0102"}),
		},
		Height: 2,
	}
	lotusBlockServiceMock.On("GetTipsetByHeight", int64(2)).Return(tipsetResponse, nil)
	lotusMinerServiceMock := &mocksLotus.MinerService{}
	lotusMinerServiceMock.On("GetMinerAddress").Return("t0101", nil)

	lotusRewardServiceMock := &mocksLotus.RewardService{}
	lotusRewardServiceMock.On("GetMiningReward", tipsetResponse.Cids).Return("100", nil)

	lotusMockedClient := &lotus.Client{
		Blocks:  lotusBlockServiceMock,
		Miner:   lotusMinerServiceMock,
		Rewards: lotusRewardServiceMock,
	}

	sessionMock := new(mocks.UserSession)
	sessionMock.On("GetLastCheckedHeight").Return(int64(1))
	sessionMock.On("SetLastCheckedHeight", int64(2)).Return()
	sessionMock.On("SaveSession").Return(nil)

	success := SubmitNewBlockReport(hactarMockedClient, lotusMockedClient, sessionMock)

	// assertions
	assert.True(t, success)
	hactarBlockServiceMock.AssertNumberOfCalls(t, "AddMiningReward", 1)
	hactarBlockServiceMock.AssertExpectations(t)
	lotusBlockServiceMock.AssertNumberOfCalls(t, "GetTipsetByHeight", 1)
	lotusBlockServiceMock.AssertNumberOfCalls(t, "GetTipsetByHeight", 1)
	lotusBlockServiceMock.AssertExpectations(t)
	lotusMinerServiceMock.AssertNumberOfCalls(t, "GetMinerAddress", 1)
	lotusMinerServiceMock.AssertExpectations(t)
	sessionMock.AssertNumberOfCalls(t, "GetLastCheckedHeight", 1)
	sessionMock.AssertNumberOfCalls(t, "SetLastCheckedHeight", 1)
	sessionMock.AssertNumberOfCalls(t, "SaveSession", 1)
	sessionMock.AssertExpectations(t)
}

func TestSubmitNewBlockReport_MultipleValidTipsetsWithRewards_Success(t *testing.T) {
	hactarBlockServiceMock := &mocksHactar.BlocksService{}
	// call for height 2
	hactarBlockServiceMock.On("AddMiningReward", []hactar.Block{
		*(&hactar.Block{
			Cid:    "bafy2bzaceawp7zcx74biecfl3axvoulh4lgdnnwzvgaza2cdhmxx75ymo6jnk",
			Reward: "100",
			Node: hactar.NodeInfo{
				Address: "t0101",
				Url:     url.GetUrl(),
			},
		}),
	}).Return(&http.Response{StatusCode: 200}, nil)
	// call for height 3
	hactarBlockServiceMock.On("AddMiningReward", []hactar.Block{
		*(&hactar.Block{
			Cid:    "bafy2bzaceawp7zcx74biecfl3axvoulh4lgdnnwzvgaza2cdhmxx75ymo8jnk",
			Reward: "100",
			Node: hactar.NodeInfo{
				Address: "t0101",
				Url:     url.GetUrl(),
			},
		}),
		*(&hactar.Block{
			Cid:    "bafy2bzaceawp7zcx74biecfl3axvoulh4lgdnnwzvgaza2cdhmxx75ymo9jnk",
			Reward: "100",
			Node: hactar.NodeInfo{
				Address: "t0101",
				Url:     url.GetUrl(),
			},
		}),
	}).Return(&http.Response{StatusCode: 200}, nil)
	// call for height 4
	hactarBlockServiceMock.On("AddMiningReward", []hactar.Block{
		*(&hactar.Block{
			Cid:    "bafy2bzaceawp7zcx74biecfl3axvoulh4lgdnnwzvgaza2cdhmxx75ymo0jnk",
			Reward: "100",
			Node: hactar.NodeInfo{
				Address: "t0101",
				Url:     url.GetUrl(),
			},
		}),
	}).Return(&http.Response{StatusCode: 200}, nil)
	hactarMockedClient := &hactar.Client{
		BaseURL:  nil,
		Token:    "",
		Nodes:    nil,
		DiskInfo: nil,
		Blocks:   hactarBlockServiceMock,
	}

	lotusBlockServiceMock := &mocksLotus.BlocksService{}
	lotusBlockServiceMock.On("GetLastHeight").Return(int64(4), nil)
	// call for height 2
	t2 := &lotus.TipsetResponse{
		Cids: []string{
			"bafy2bzaceawp7zcx74biecfl3axvoulh4lgdnnwzvgaza2cdhmxx75ymo6jnk",
			"bafy2bzaceawp7zcx74biecfl3axvoulh4lgdnnwzvgaza2cdhmxx75ymo7jnk",
		},
		Blocks: []lotus.BlockStruct{
			*(&lotus.BlockStruct{Miner: "t0101"}),
			*(&lotus.BlockStruct{Miner: "t0102"}),
		},
		Height: 2,
	}
	lotusBlockServiceMock.On("GetTipsetByHeight", int64(2)).Return(t2, nil)
	// call for height 3
	t3 := &lotus.TipsetResponse{
		Cids: []string{
			"bafy2bzaceawp7zcx74biecfl3axvoulh4lgdnnwzvgaza2cdhmxx75ymo8jnk",
			"bafy2bzaceawp7zcx74biecfl3axvoulh4lgdnnwzvgaza2cdhmxx75ymo9jnk",
		},
		Blocks: []lotus.BlockStruct{
			*(&lotus.BlockStruct{Miner: "t0101"}),
			*(&lotus.BlockStruct{Miner: "t0101"}),
		},
		Height: 3,
	}
	lotusBlockServiceMock.On("GetTipsetByHeight", int64(3)).Return(t3, nil)
	// call for height 4
	t4 := &lotus.TipsetResponse{
		Cids: []string{
			"bafy2bzaceawp7zcx74biecfl3axvoulh4lgdnnwzvgaza2cdhmxx75ymo0jnk",
		},
		Blocks: []lotus.BlockStruct{
			*(&lotus.BlockStruct{Miner: "t0101"}),
		},
		Height: 4,
	}
	lotusBlockServiceMock.On("GetTipsetByHeight", int64(4)).Return(t4, nil)

	lotusMinerServiceMock := &mocksLotus.MinerService{}
	lotusMinerServiceMock.On("GetMinerAddress").Return("t0101", nil)

	lotusRewardServiceMock := &mocksLotus.RewardService{}
	lotusRewardServiceMock.On("GetMiningReward", mock.AnythingOfType("[]string")).Return("100", nil)

	lotusMockedClient := &lotus.Client{
		Blocks:  lotusBlockServiceMock,
		Miner:   lotusMinerServiceMock,
		Rewards: lotusRewardServiceMock,
	}

	sessionMock := new(mocks.UserSession)
	sessionMock.On("GetLastCheckedHeight").Return(int64(1))
	sessionMock.On("SetLastCheckedHeight", int64(2)).Return()
	sessionMock.On("SetLastCheckedHeight", int64(3)).Return()
	sessionMock.On("SetLastCheckedHeight", int64(4)).Return()
	sessionMock.On("SaveSession").Return(nil)

	success := SubmitNewBlockReport(hactarMockedClient, lotusMockedClient, sessionMock)

	// assertions
	assert.True(t, success)
	hactarBlockServiceMock.AssertNumberOfCalls(t, "AddMiningReward", 3)
	hactarBlockServiceMock.AssertExpectations(t)
	lotusBlockServiceMock.AssertNumberOfCalls(t, "GetTipsetByHeight", 3)
	lotusBlockServiceMock.AssertExpectations(t)
	lotusMinerServiceMock.AssertNumberOfCalls(t, "GetMinerAddress", 1)
	lotusMinerServiceMock.AssertExpectations(t)
	sessionMock.AssertNumberOfCalls(t, "GetLastCheckedHeight", 1)
	sessionMock.AssertNumberOfCalls(t, "SetLastCheckedHeight", 3)
	sessionMock.AssertNumberOfCalls(t, "SaveSession", 3)
	sessionMock.AssertExpectations(t)
}

func TestSubmitNewBlockReport_ValidTipsetWithoutReward_Success(t *testing.T) {
	hactarMockedClient := &hactar.Client{
		BaseURL:  nil,
		Token:    "",
		Nodes:    nil,
		DiskInfo: nil,
		Blocks:   nil,
	}

	lotusBlockServiceMock := &mocksLotus.BlocksService{}
	lotusBlockServiceMock.On("GetLastHeight").Return(int64(2), nil)
	t2 := &lotus.TipsetResponse{
		Cids: []string{
			"bafy2bzaceawp7zcx74biecfl3axvoulh4lgdnnwzvgaza2cdhmxx75ymo6jnk",
			"bafy2bzaceawp7zcx74biecfl3axvoulh4lgdnnwzvgaza2cdhmxx75ymo7jnk",
		},
		Blocks: []lotus.BlockStruct{
			*(&lotus.BlockStruct{Miner: "t0102"}),
			*(&lotus.BlockStruct{Miner: "t0102"}),
		},
		Height: 2,
	}

	lotusBlockServiceMock.On("GetTipsetByHeight", int64(2)).Return(t2, nil)
	lotusMinerServiceMock := &mocksLotus.MinerService{}
	lotusMinerServiceMock.On("GetMinerAddress").Return("t0101", nil)
	lotusMockedClient := &lotus.Client{
		Blocks: lotusBlockServiceMock,
		Miner:  lotusMinerServiceMock,
	}

	sessionMock := new(mocks.UserSession)
	sessionMock.On("GetLastCheckedHeight").Return(int64(1))
	sessionMock.On("SetLastCheckedHeight", int64(2)).Return()
	sessionMock.On("SaveSession").Return(nil)

	success := SubmitNewBlockReport(hactarMockedClient, lotusMockedClient, sessionMock)

	// assertions
	assert.True(t, success)
	lotusBlockServiceMock.AssertNumberOfCalls(t, "GetTipsetByHeight", 1)
	lotusBlockServiceMock.AssertNumberOfCalls(t, "GetTipsetByHeight", 1)
	lotusBlockServiceMock.AssertExpectations(t)
	lotusMinerServiceMock.AssertNumberOfCalls(t, "GetMinerAddress", 1)
	lotusMinerServiceMock.AssertExpectations(t)
	sessionMock.AssertNumberOfCalls(t, "GetLastCheckedHeight", 1)
	sessionMock.AssertNumberOfCalls(t, "SetLastCheckedHeight", 1)
	sessionMock.AssertNumberOfCalls(t, "SaveSession", 1)
	sessionMock.AssertExpectations(t)
}
