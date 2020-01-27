package session

import 	"github.com/spf13/viper"

type UserSession struct {
	Token string
	viper *viper.Viper
}

var CurrentUser = &UserSession{}

func LinkViper(viper *viper.Viper) {
	CurrentUser.viper = viper
}

func SaveSession() error {
	err := CurrentUser.viper.SafeWriteConfig()
	return err
}
