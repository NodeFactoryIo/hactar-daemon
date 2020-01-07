package token

import (
	"encoding/base64"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io/ioutil"
)

const (
	TokenFileDefaultPath = "~/.lotus/token"
)

// Read lotus token from file
// If token file address is not provided in configuration, default path will be used.
func ReadTokenFromFile() string {
	tokenFile := viper.GetString("jsonrpc.token-address")
	if tokenFile == "" {
		// set to default if not defined in config
		tokenFile = TokenFileDefaultPath
	}

	data, err := ioutil.ReadFile(tokenFile)
	if err != nil {
		log.Error("")
		return ""
	}
	return base64.StdEncoding.EncodeToString(data)
}