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
	DefaultPathNodeTokenFile = "/.lotus/token"
	DefaultPathMinerTokenFile = "/.lotusstorage/token"
)

// Read lotus token from file
// If token file address is not provided in configuration, default path will be used.
func ReadNodeTokenFromFile() string {
	tokenFile := viper.GetString("jsonrpc.lotus-node.token")
	if tokenFile == "" {
		// set to default if not defined in config
		usr, _ := user.Current()
		dir := usr.HomeDir
		tokenFile = dir + DefaultPathNodeTokenFile
	}

	return readTokenFromFile(tokenFile)
}

func ReadMinerTokenFromFile() string {
	tokenFile := viper.GetString("jsonrpc.lotus-miner.token")
	if tokenFile == "" {
		// set to default if not defined in config
		usr, _ := user.Current()
		dir := usr.HomeDir
		tokenFile = dir + DefaultPathMinerTokenFile
	}

	return readTokenFromFile(tokenFile)
}

func readTokenFromFile(tokenFile string) string {
	data, err := ioutil.ReadFile(tokenFile)
	if err != nil {
		log.Error(fmt.Sprintf("Error reading token from file: %s", tokenFile))
		return ""
	}
	return strings.TrimSpace(string(data))
}

func DisplayTokens() {
	fmt.Printf("Node token:\n%s\n", ReadNodeTokenFromFile())
	fmt.Printf("Miner token:\n%s\n", ReadMinerTokenFromFile())
}
