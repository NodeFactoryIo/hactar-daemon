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

func SetUpLogger() {
	// init formatter
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			s := strings.Split(f.Function, ".")
			funcname := s[len(s)-1]
			_, filename := path.Split(f.File)
			return funcname, filename
		},
	})
	// set log level
	result, err := log.ParseLevel(util.String(viper.Get("log.level")))
	log.SetLevel(util.If(err != nil).Level(log.InfoLevel, result))
	// TODO make logger write to file and possibly change format of logging
	log.SetOutput(os.Stdout)
}
