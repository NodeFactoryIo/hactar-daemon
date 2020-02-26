package pastdealsinfo

import (
	"fmt"
	"github.com/NodeFactoryIo/hactar-daemon/internal/hactar"
	"github.com/NodeFactoryIo/hactar-daemon/internal/lotus"
	"github.com/NodeFactoryIo/hactar-daemon/internal/url"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func SendPastDealsInfo(hactarClient *hactar.Client, lotusClient *lotus.Client) bool {
	minerAddress, err := lotusClient.Miner.GetMinerAddress()
	if err != nil {
		log.Error("Unable to get miner address ", err)
		return false
	}

	pastDeals, err := lotusClient.PastDeals.GetAllPastDeals()
	if err != nil {
		log.Error("Unable to get past deals ", err)
		return false
	}

	pastDealsRequest := hactar.PastDealsInfo{
		Deals: pastDeals,
		Node: hactar.NodeInfo{
			Address: minerAddress,
			Url:     url.GetUrl(),
		},
	}

	response, err := hactarClient.PastDeals.SendPastDealsInfo(pastDealsRequest)

	if err != nil {
		log.Error(
			fmt.Sprintf("Unable to send past deals for node: %s", minerAddress),
			err,
		)
		return false
	}
	if response != nil && response.StatusCode == http.StatusOK {
		log.Info(fmt.Sprintf("Past deals for node %s sent", minerAddress))
	}

	return true
}
