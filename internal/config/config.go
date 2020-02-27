package config

import (
	"github.com/spf13/viper"
	"os"
)

func InitMainConfig() {
	setDefaultValuesForMainConfig()
	// load config file
	viper.SetConfigName(getMainConfigName()) // name of config file (without extension)
	viper.AddConfigPath(".")                 // look for config in the working directory
	_ = viper.ReadInConfig()
}

func InitStatusConfig() (*viper.Viper, string) {
	// load status file
	status := viper.New()
	status.SetConfigName("status") // name of config file (without extension)
	status.AddConfigPath(".")      // look for config in the working directory
	status.SetConfigType("yaml")
	// set default values
	status.SetDefault("hactar.token", "")
	status.SetDefault("lotus.block.last-checked", "")
	_ = status.ReadInConfig()
	return status, "./status.yaml"
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
}

// depending on ENV variable creates name for config file
func getMainConfigName() string {
	configFileName := "config"
	if env := os.Getenv("ENV"); env != "" {
		configFileName = configFileName + "-" + env
	}
	return configFileName
}
