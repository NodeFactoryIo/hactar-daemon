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
	assert.Equal(t, viper.GetInt("stats.interval"), 3600)
	assert.Equal(t, viper.GetInt("stats.blocks.interval"), 900)
	assert.Equal(t, viper.GetInt("stats.uptime.interval"), 3600)
	assert.Equal(t, viper.GetInt("stats.balance.interval"), 900)
	assert.Equal(t, viper.GetString("log.level"), "error")
}
