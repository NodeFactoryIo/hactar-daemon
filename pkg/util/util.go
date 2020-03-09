package util

import (
	"bytes"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
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

func ReaderToString(reader io.ReadCloser) string {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(reader)
	if err != nil {
		log.Error(err)
	}
	return buf.String()
}
