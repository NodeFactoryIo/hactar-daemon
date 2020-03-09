package config

import (
	"github.com/spf13/viper"
	"os"
	"github.com/getsentry/sentry-go"
	"time"
)

func InitMainConfig() {
	setDefaultValuesForMainConfig()
	// load config file
	viper.SetConfigName(getMainConfigName()) // name of config file (without extension)
	viper.AddConfigPath(".")                 // look for config in the working directory
	_ = viper.ReadInConfig()

	if os.Getenv("ENV") != "test" {
		setupSentry()
	}
}

func setDefaultValuesForMainConfig() {
	viper.SetDefault("stats.interval", 60)
	viper.SetDefault("stats.blocks.interval", 50)
	viper.SetDefault("stats.uptime.interval", 10)
	viper.SetDefault("stats.balance.interval", 150)
	viper.SetDefault("jsonrpc.lotus-miner.url", "")
	viper.SetDefault("jsonrpc.lotus-miner.token", "")
	viper.SetDefault("jsonrpc.lotus-node.url", "")
	viper.SetDefault("jsonrpc.lotus-node.token", "")
	viper.SetDefault("hactar.api-url", "")
	viper.SetDefault("log.level", "error")
	viper.SetDefault("lotus.network-address", "t01")
	viper.SetDefault("sentry.key", "")
	viper.SetDefault("sentry.project", "")
}

// depending on ENV variable creates name for config file
func getMainConfigName() string {
	configFileName := "config"
	if env := os.Getenv("ENV"); env != "" {
		configFileName = configFileName + "-" + env
	}
	return configFileName
}

func setupSentry() {
	key := viper.GetString("sentry.key")
	project := viper.GetString("sentry.project")
	if key != "" && project != "" {
		sentry.Init(sentry.ClientOptions{
			Dsn: "https://" + key + "@sentry.io/" + project,
			Debug: true,
		})

		// Flush buffered events before the program terminates.
		defer sentry.Flush(2 * time.Second)
	}
}
