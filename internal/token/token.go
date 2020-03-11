package token

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io/ioutil"
	"os/user"
	"strings"
)

const (
	DefaultPathNodeTokenFile  = "/.lotus/token"
	DefaultPathMinerTokenFile = "/.lotusstorage/token"
)

// Read lotus node token from file
// If token file address is not provided in configuration, default path will be used.
func ReadNodeTokenFromFile() string {
	return readTokenFromFile(viper.GetString("jsonrpc.lotus-node.token"), DefaultPathNodeTokenFile)
}

// Read lotus miner token from file
// If token file address is not provided in configuration, default path will be used.
func ReadMinerTokenFromFile() string {
	return readTokenFromFile(viper.GetString("jsonrpc.lotus-miner.token"), DefaultPathMinerTokenFile)
}

func readTokenFromFile(tokenFile string, defaultPath string) string {
	if tokenFile == "" {
		// set to default if not defined in config
		usr, _ := user.Current()
		dir := usr.HomeDir
		tokenFile = dir + defaultPath
	}

	data, err := ioutil.ReadFile(tokenFile)
	if err != nil {
		log.Error(fmt.Sprintf("Error reading token from file: %s", tokenFile))
		sentry.CaptureException(err)
		return ""
	}
	return strings.TrimSpace(string(data))
}
