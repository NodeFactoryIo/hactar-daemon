package diskinfo

import (
	"github.com/NodeFactoryIo/hactar-daemon/internal/hactar"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type DiskStatus struct {
	All  uint64
	Used uint64
	Free uint64
}

func SendDiskInfoStats(client *hactar.Client, actorAddress string, nodeUrl string) bool {
	usage := DiskUsage("/")

	response, err := client.DiskInfo.SendDiskInfo(hactar.DiskInfo{
		FreeSpace:  strconv.FormatUint(usage.Free, 10),
		TakenSpace: strconv.FormatUint(usage.Used, 10),
		Node: hactar.NodeInfo{
			Address: actorAddress,
			Url:     nodeUrl,
		},
	})

	if response != nil && response.StatusCode == http.StatusCreated {
		log.Info("Disk info successfully sent to hactar")
		return true
	}

	log.Error("Unable to send disk information statistics ", err)
	return false
}
