package util

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

// Panic if error exists
func Must(err error, message string) {
	if err != nil {
		log.Error(err, message)
		panic(fmt.Errorf("%s: %s\n", message, err.Error()))
	}
}

// Converts interface to string
func String(i interface{}) string {
	return fmt.Sprintf("%v", i)
}
