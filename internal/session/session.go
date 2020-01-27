package session

import (
	"github.com/spf13/viper"
)

type UserSession struct {
	HactarToken      string
	LastCheckedBlock string
	viper            *viper.Viper
}

var CurrentUser = &UserSession{}

func LoadSession(viper *viper.Viper) {
	CurrentUser.viper = viper
	CurrentUser.HactarToken = viper.GetString("hactar.token")
	CurrentUser.LastCheckedBlock = viper.GetString("lotus.block.last-checked")
}

func SaveSession() error {
	CurrentUser.viper.Set("lotus.block.last-checked", CurrentUser.LastCheckedBlock)
	CurrentUser.viper.Set("hactar.token", CurrentUser.HactarToken)
	err := CurrentUser.viper.WriteConfig()
	return err
}
