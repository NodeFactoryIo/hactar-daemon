package lotus

import (
	"errors"
	"github.com/NodeFactoryIo/hactar-daemon/pkg/util/bigint"
	log "github.com/sirupsen/logrus"
	"math/big"
)

type RewardService interface {
	GetMiningReward() (string, error)
}

type rewardService struct {
	client *Client
}

func (rs *rewardService) GetMiningReward() (string, error) {
	balance, err := rs.client.Wallet.GetWalletBalance("t01") // todo add to config
	if err != nil {
		log.Error("Unable to get balance for network address", err)
		return "", err
	}
	// calculate mining reward
	if ci, ok := big.NewInt(0).SetString(balance, 10); ok == true {
		res := ci.Mul(ci, InitialReward)
		res = res.Div(res, FromFil(MiningRewardTotal).Int)
		res = res.Div(res, bigint.NewInt(BlocksPerEpoch).Int)
		return res.String(), nil
	} else {
		log.Error("Unable to convert balance of network address.")
		return "", errors.New("unable to convert balance of network address")
	}
}
