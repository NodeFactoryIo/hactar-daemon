package config

import (
	"github.com/magiconair/properties/assert"
	"github.com/spf13/viper"
	"os"
	"testing"
)

func TestInitMainConfig_getMainConfigName(t *testing.T) {
	_ = os.Setenv("ENV", "test")
	configName := getMainConfigName()
	assert.Equal(t, configName, "config-test")

	_ = os.Setenv("ENV", "")
	configName = getMainConfigName()
	assert.Equal(t, configName, "config")
}

func TestInitMainConfig_testDefaultValues(t *testing.T) {
	InitMainConfig()
	assert.Equal(t, viper.GetInt("stats.interval"), 60)
	assert.Equal(t, viper.GetInt("stats.blocks.interval"), 50)
	assert.Equal(t, viper.GetInt("stats.uptime.interval"), 10)
	assert.Equal(t, viper.GetInt("stats.balance.interval"), 150)
	assert.Equal(t, viper.GetString("log.level"), "error")
}
