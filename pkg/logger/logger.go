package logger

import (
	log "github.com/sirupsen/logrus"
	"os"
	"path"
	"runtime"
	"strings"
)

func init()  {
	log.SetFormatter(&log.TextFormatter{})
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{
		DisableTimestamp: false,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			s := strings.Split(f.Function, ".")
			funcname := s[len(s)-1]
			_, filename := path.Split(f.File)
			return funcname, filename
		},
	})
	// TODO make logger write to file and change format of logging
	log.SetOutput(os.Stdout)
}
