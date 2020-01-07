package util

import "fmt"

// Panic if error exists
func Must(err error, message string) {
	if err != nil {
		panic(fmt.Errorf("%s: %s\n", message, err.Error()))
	}
}