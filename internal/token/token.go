package token

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io/ioutil"
	"os/user"
	"strings"
)

const (
	DefaultPathTokenFile = "/.lotus/token"
)

// Read lotus token from file
// If token file address is not provided in configuration, default path will be used.
func ReadTokenFromFile() string {
	tokenFile := viper.GetString("jsonrpc.token-address")
	if tokenFile == "" {
		// set to default if not defined in config
		usr, _ := user.Current()
		dir := usr.HomeDir
		tokenFile = dir + DefaultPathTokenFile
	}

	data, err := ioutil.ReadFile(tokenFile)
	if err != nil {
		log.Error(fmt.Sprintf("Error reading token from file: %s", tokenFile))
		return ""
	}
	return strings.TrimSpace(string(data))
}

func DisplayToken() {
	token := ReadTokenFromFile()
	fmt.Printf("Token:\n%s\n", token)
}