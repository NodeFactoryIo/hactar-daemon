package stats

import (
	"fmt"
	"github.com/NodeFactoryIo/hactar-daemon/internal/hactar"
	"github.com/NodeFactoryIo/hactar-daemon/internal/lotus/services"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"strconv"
	"time"
)

func SubmitNewStatsReport() bool  {
	client := hactar.NewClient(nil)
	lotusService := services.NewLotusService(nil, nil)

	actorAddress, err := lotusService.GetMinerAddress()
	if err != nil {
		fmt.Print("Worker down!")
		// return nil
	}


	sendDiskInfoStats(client, actorAddress, "")
	// TODO collect all statistic and send it to backend
	fmt.Println("Collecting stats and sending report....")
	return true
}

func sendDiskInfoStats(client *hactar.Client, actorAddress string, nodeUrl string) {
	usage := DiskUsage("/")

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
		log.Info(fmt.Sprintf("Disk: %d (free) %d (used)", usage.Free, usage.Used))
	}
}

func StartMonitoringStats() {
	interval, _ := strconv.Atoi(viper.GetString("stats.interval"))
	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				log.Info("Stats monitor ticked")
				SubmitNewStatsReport()
			}
		}
	}()
}