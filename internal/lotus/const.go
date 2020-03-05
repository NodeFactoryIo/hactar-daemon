package lotus

import (
	"github.com/NodeFactoryIo/hactar-daemon/pkg/util/bigint"
	"math/big"
)

// InitialReward
const initialRewardStr = "153856861913558700202"

var InitialReward *big.Int

// MiningRewardTotal
const miningRewardTotal = 1_400_000_000
const filecoinPrecision = 1_000_000_000_000_000_000

var MiningRewardTotal *big.Int

// BlocksPerEpoch
const blocksPerEpoch = 5

var BlocksPerEpoch *big.Int

func init() {
	InitialReward = new(big.Int)
	var ok bool
	InitialReward, ok = InitialReward.
		SetString(initialRewardStr, 10)
	if !ok {
		panic("could not parse InitialRewardStr")
	}

	MiningRewardTotal = fromFil(miningRewardTotal).Int

	BlocksPerEpoch = bigint.NewInt(blocksPerEpoch).Int
}

func fromFil(i uint64) bigint.BigInt {
	return bigint.BigMul(bigint.NewInt(i), bigint.NewInt(filecoinPrecision))
}
