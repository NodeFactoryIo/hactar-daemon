// Package requests defines jsonrpc method names for communication with lotus
package lotus

// General requests
const (
	Version              = "Filecoin.Version"
	ActorAddress         = "Filecoin.ActorAddress"
	WalletBalance        = "Filecoin.WalletBalance"
	WalletDefaultAddress = "Filecoin.WalletDefaultAddress"
	Actor                = "Filecoin.StateGetActor"
	NumberOfSectors      = "Filecoin.StateMinerSectorCount"
	SectorSize           = "Filecoin.StateMinerSectors"
	MinerInfo            = "Filecoin.StateMinerInfo"
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
