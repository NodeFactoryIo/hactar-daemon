package stats

import (
	"fmt"
	"github.com/NodeFactoryIo/hactar-daemon/internal/hactar"
	log "github.com/sirupsen/logrus"
	"net/http"
	"syscall"
)

type DiskStatus struct {
	All  uint64
	Used uint64
	Free uint64
}

func SendDiskInfoStats(client *hactar.Client, actorAddress string, nodeUrl string) {
	usage := diskUsage("/")

	response, err := client.DiskInfo.SendDiskInfo(hactar.DiskInfo{
		FreeSpace:    string(usage.Free),
		TakenSpace:   string(usage.Used),
		NodeUrl:      nodeUrl,
		ActorAddress: actorAddress,
	})

	if err != nil {
		log.Error("Unable to send disk information statistics. ", err)
	}

	if response != nil && response.StatusCode == http.StatusOK {
		log.Info(fmt.Sprintf("Disk stats: %d (free) %d (used)", usage.Free, usage.Used))
	}
}

// calculates disk usage of path/disk
func diskUsage(path string) (disk DiskStatus) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)
	if err != nil {
		return
	}
	disk.All = fs.Blocks * uint64(fs.Bsize)
	disk.Free = fs.Bfree * uint64(fs.Bsize)
	disk.Used = disk.All - disk.Free
	return
}