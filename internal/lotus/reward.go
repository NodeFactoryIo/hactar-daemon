package lotus

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"math/big"
)

type RewardService interface {
	GetMiningReward([]string) (string, error)
}

type rewardService struct {
	client *Client
}

func (rs *rewardService) GetMiningReward(cids []string) (string, error) {
	actor, err := rs.client.Miner.GetActor(viper.GetString("lotus.network-address"), cids[0])
	if err != nil {
		log.Error("Unable to get balance for network address", err)
		return "", err
	}
	// calculate mining reward
	if ci, ok := big.NewInt(0).SetString(actor.Balance, 10); ok == true {
		res := ci.Mul(ci, InitialReward)
		res = res.Div(res, MiningRewardTotal)
		res = res.Div(res, BlocksPerEpoch)
		return res.String(), nil
	} else {
		log.Error("Unable to convert balance of network address.")
		return "", errors.New("unable to convert balance of network address")
	}
}
