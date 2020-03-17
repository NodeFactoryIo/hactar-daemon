package config

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/getsentry/sentry-go"
	"github.com/subosito/gotenv"
	"os"
	"time"
)

func InitMainConfig() {
	setDefaultValuesForMainConfig()
	// load config file
	viper.SetConfigName(getMainConfigName()) // name of config file (without extension)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../../")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// alternatively look for config in the working directory
			viper.AddConfigPath(".")
			_ = viper.ReadInConfig()

		} else {
			fmt.Printf("Error while loading config: %s \n", err)
		}
	}

	// Load env variables from .env
	gotenv.Load()

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
	dsn := os.Getenv("SENTRY_DSN")
	sentry.Init(sentry.ClientOptions{
		Dsn: dsn,
		Debug: false,
	})

	// Flush buffered events before the program terminates.
	defer sentry.Flush(2 * time.Second)
}
