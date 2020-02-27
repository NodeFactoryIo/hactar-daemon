package session

import (
	"github.com/spf13/viper"
)

type UserSession interface {
	SaveSession() error
	GetHactarToken() string
	SetHactarToken(token string)
	GetLastCheckedHeight() int64
	SetLastCheckedHeight(height int64)
	GetNodeMinerAddress() string
	SetNodeMinerAddress(address string)
}

type userSession struct {
	// persisted values
	hactarToken       string
	lastCheckedHeight int64
	// memory values
	nodeMinerAddress string
	filepath         string
	viper            *viper.Viper
}

// instance holding information about current session
var CurrentSession *userSession

// initialization function
func InitSession(viper *viper.Viper, filepath string) {
	CurrentSession = &userSession{
		hactarToken:       viper.GetString("hactar.token"),
		lastCheckedHeight: viper.GetInt64("lotus.block.last-checked"),
		filepath:          filepath,
		viper:             viper,
	}
}

// implementation of UserSession interface

func (session *userSession) SaveSession() error {
	session.viper.Set("lotus.block.last-checked", session.lastCheckedHeight)
	session.viper.Set("hactar.token", session.hactarToken)
	err := session.viper.WriteConfigAs(session.filepath)
	return err
}

func (session *userSession) GetNodeMinerAddress() string {
	return session.nodeMinerAddress
}

func (session *userSession) SetNodeMinerAddress(address string) {
	session.nodeMinerAddress = address
}

func (session *userSession) GetHactarToken() string {
	return session.hactarToken
}

func (session *userSession) SetHactarToken(token string) {
	session.hactarToken = token
}

func (session *userSession) GetLastCheckedHeight() int64 {
	return session.lastCheckedHeight
}

func (session *userSession) SetLastCheckedHeight(height int64) {
	session.lastCheckedHeight = height
}
