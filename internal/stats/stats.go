package stats

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strconv"
	"time"
)

func SubmitNewStatsReport() bool  {
	// TODO collect all statistic and send it to backend
	fmt.Println("Collecting stats and sending report....")
	return true
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