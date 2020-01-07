// Package requests defines jsonrpc method names for communication with lotus
package requests

// General requests
const (
	Version       = "Filecoin.Version"
	ActorAddress  = "Filecoin.ActorAddress"
	WalletBalance = "Filecoin.WalletBalance"
	MinerBalance  = "Filecoin.StateGetActor"
	Sectors       = "Filecoin.StateMinerSectors"
	SectorSize    = "Filecoin.StateMinerSectorSize"
	MinerPower    = "Filecoin.StateMinerPower"
	PastDeals     = "Filecoin.ClientListDeals"
	DealDetails   = "Filecoin.ClientGetDealInfo"
	HeadBlock     = "Filecoin.ChainHead"
	Block         = "Filecoin.ChainGetBlock"
)

// Payment channels requests
const (
	PayChannels = "Filecoin.PaychList"
	PayChannel  = "Filecoin.PaychStatus"
	PayChannelVouchers = "Filecoin.PaychVoucherList"
)