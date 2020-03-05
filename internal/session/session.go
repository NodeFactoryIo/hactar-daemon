package session

import (
	"github.com/NodeFactoryIo/hactar-daemon/pkg/util"
	"github.com/spf13/viper"
	"os"
	"time"
)

type UserSession interface {
	SaveSession() error
	GetHactarToken() string
	SetHactarToken(token string)
	GetLastCheckedHeight() int64
	SetLastCheckedHeight(height int64)
	GetLastCheckedHeightTimestamp() string
	GetNodeMinerAddress() string
	SetNodeMinerAddress(address string)
}

type userSession struct {
	// persisted values
	hactarToken                string
	lastCheckedHeight          int64
	lastCheckedHeightTimestamp string
	// memory values
	nodeMinerAddress string
	filepath         string
	viper            *viper.Viper
}

// instance holding information about current session
var CurrentSession *userSession

func InitSession() {
	// define file path
	rootDir := getHactarDir()
	filepath := rootDir + "/status.yaml"
	// define status file
	status := viper.New()
	status.SetConfigName("status") // name of config file (without extension)
	status.AddConfigPath(rootDir)  // look for config in the working directory
	status.SetConfigType("yaml")
	// set default values
	status.SetDefault("hactar.token", "")
	status.SetDefault("lotus.block.last-checked", "")
	status.SetDefault("lotus.block.last-checked-timestamp", "")
	// try to read existing file
	_ = status.ReadInConfig()
	// save current session
	CurrentSession = &userSession{
		hactarToken:       viper.GetString("hactar.token"),
		lastCheckedHeight: viper.GetInt64("lotus.block.last-checked"),
		filepath:          filepath,
		viper:             status,
	}
	// creates status file if not created
	_ = CurrentSession.SaveSession()
}

func getHactarDir() string {
	// find user home dir
	userHomeDir, err := os.UserHomeDir()
	util.Must(err, "Unable to find home directory.")
	// hactar directory
	hactarDir := userHomeDir + "/.lotus/hactar"
	err = os.MkdirAll(hactarDir, os.ModePerm)
	util.Must(err, "Unable to create hactar folder.")
	return hactarDir
}

// implementation of UserSession interface

func (session *userSession) SaveSession() error {
	session.viper.Set("lotus.block.last-checked", session.lastCheckedHeight)
	session.viper.Set("hactar.token", session.hactarToken)
	session.viper.Set("lotus.block.last-checked-timestamp", session.lastCheckedHeightTimestamp)
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
	// save timestamp
	timestamp := time.Now().Format("2002-02-02 15:04:33")
	session.lastCheckedHeightTimestamp = timestamp
}

func (session *userSession) GetLastCheckedHeightTimestamp() string {
	return session.lastCheckedHeightTimestamp
}
