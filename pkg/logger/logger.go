package logger

import (
	"github.com/NodeFactoryIo/hactar-daemon/pkg/util"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"path"
	"runtime"
	"strings"
)

func SetUpLogger(logLevel log.Level) {
	setUpLogger(logLevel)
}

func SetUpDefaultLogger() {
	result, err := log.ParseLevel(util.String(viper.GetString("log.level")))
	setUpLogger(util.If(err != nil).Level(log.ErrorLevel, result))
}

func setUpLogger(level log.Level) {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			s := strings.Split(f.Function, ".")
			funcname := s[len(s)-1]
			_, filename := path.Split(f.File)
			return funcname, filename
		},
	})
	log.SetLevel(level)
	log.SetOutput(os.Stdout)
}
