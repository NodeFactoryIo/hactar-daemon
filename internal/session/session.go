package session

import "github.com/spf13/viper"

type UserSession struct {
	Token            string
	LastCheckedBlock string
	viper            *viper.Viper
}

var CurrentUser = &UserSession{}

func LoadSession(viper *viper.Viper) {
	CurrentUser.viper = viper
	CurrentUser.Token = viper.GetString("hactar.token")
	CurrentUser.LastCheckedBlock = viper.GetString("lotus.block.last-checked")
}

func SaveSession() error {
	err := CurrentUser.viper.SafeWriteConfig()
	return err
}
