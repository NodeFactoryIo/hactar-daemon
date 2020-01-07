package util

import (
	"fmt"
)

// Panic if error exists
func Must(err error, message string) {
	if err != nil {
		panic(fmt.Errorf("%s: %s\n", message, err.Error()))
	}
}

// Converts interface to string
func String(i interface{}) string {
	return fmt.Sprintf("%v", i)
}