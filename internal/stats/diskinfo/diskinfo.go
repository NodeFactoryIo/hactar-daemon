package diskinfo

import (
	"github.com/NodeFactoryIo/hactar-daemon/internal/hactar"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type DiskStatus struct {
	All  uint64
	Used uint64
	Free uint64
}

func SendDiskInfoStats(client *hactar.Client, actorAddress string, nodeUrl string) bool {
	usage := DiskUsage("/")

	response, err := client.DiskInfo.SendDiskInfo(hactar.DiskInfo{
		FreeSpace:    string(usage.Free),
		TakenSpace:   string(usage.Used),
		Node: hactar.NodeInfo{
			Address: actorAddress,
			Url:     nodeUrl,
		},
	})

	if response != nil && response.StatusCode == http.StatusOK {
		log.Info("Disk info successfully sent to hactar")
		return true
	}

	log.Error("Unable to send disk information statistics ", err)
	return false
}
