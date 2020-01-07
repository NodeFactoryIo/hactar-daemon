package token

import (
	"encoding/base64"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io/ioutil"
)

const (
	DefaultPathTokenFile = "~/.lotus/token"
)

// Read lotus token from file
// If token file address is not provided in configuration, default path will be used.
func ReadTokenFromFile() string {
	tokenFile := viper.GetString("jsonrpc.token-address")
	if tokenFile == "" {
		// set to default if not defined in config
		tokenFile = DefaultPathTokenFile
	}

	data, err := ioutil.ReadFile(tokenFile)
	if err != nil {
		log.Error(fmt.Sprintf("Error reading token from file: %s", tokenFile))
		return ""
	}
	return base64.StdEncoding.EncodeToString(data)
}