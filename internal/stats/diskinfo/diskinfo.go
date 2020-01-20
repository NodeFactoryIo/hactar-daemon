package diskinfo

import (
	"fmt"
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
		NodeUrl:      nodeUrl,
		ActorAddress: actorAddress,
	})

	if response != nil && response.StatusCode == http.StatusOK {
		log.Info(fmt.Sprintf("Disk stats: %d (free) %d (used)", usage.Free, usage.Used))
		return true
	}

	log.Error("Unable to send disk information statistics. ", err)
	return false
}