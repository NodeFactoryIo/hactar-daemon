package lotus

import (
	"github.com/NodeFactoryIo/hactar-daemon/pkg/util/bigint"
	"math/big"
)

var InitialReward *big.Int

const InitialRewardStr = "153856861913558700202"

const MiningRewardTotal = 1_400_000_000
const FilecoinPrecision = 1_000_000_000_000_000_000
const BlocksPerEpoch = 5

func FromFil(i uint64) bigint.BigInt {
	return bigint.BigMul(bigint.NewInt(i), bigint.NewInt(FilecoinPrecision))
}

func init() {
	InitialReward = new(big.Int)

	var ok bool
	InitialReward, ok = InitialReward.
		SetString(InitialRewardStr, 10)
	if !ok {
		panic("could not parse InitialRewardStr")
	}
}
