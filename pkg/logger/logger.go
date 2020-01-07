package logger

import (
	"github.com/NodeFactoryIo/hactar-daemon/internal/util"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"path"
	"runtime"
	"strings"
)

func init()  {
	// init formatter
	log.SetFormatter(&log.TextFormatter{
		DisableTimestamp: false,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			s := strings.Split(f.Function, ".")
			funcname := s[len(s)-1]
			_, filename := path.Split(f.File)
			return funcname, filename
		},
	})
	// set log level
	result, err := log.ParseLevel(util.String(viper.Get("log.level")))
	log.SetLevel(util.If(err != nil).Level(log.WarnLevel, result))
	// TODO make logger write to file and change format of logging
	log.SetOutput(os.Stdout)
}
