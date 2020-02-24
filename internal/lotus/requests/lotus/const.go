// Package requests defines jsonrpc method names for communication with lotus
package lotus

// General requests
const (
	Version              = "Filecoin.Version"
	ActorAddress         = "Filecoin.ActorAddress"
	WalletBalance        = "Filecoin.WalletBalance"
	WalletDefaultAddress = "Filecoin.WalletDefaultAddress"
	MinerBalance         = "Filecoin.StateGetActor"
	Sectors              = "Filecoin.StateMinerSectors"
	SectorSize           = "Filecoin.StateMinerSectorSize"
	MinerPower           = "Filecoin.StateMinerPower"
	PastDeals            = "Filecoin.ClientListDeals"
	DealDetails          = "Filecoin.ClientGetDealInfo"
	HeadBlock            = "Filecoin.ChainHead"
	Block                = "Filecoin.ChainGetBlock"
	TipSetByHeight       = "Filecoin.ChainGetTipSetByHeight"
)

// Payment channels requests
const (
	PayChannels        = "Filecoin.PaychList"
	PayChannel         = "Filecoin.PaychStatus"
	PayChannelVouchers = "Filecoin.PaychVoucherList"
)
