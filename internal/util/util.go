package util

import (
	"fmt"
	"github.com/sirupsen/logrus"
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

// Kind of ternary operator, used for conditional initialization
type If bool

func (c If) String(a, b string) string {
	if c {
		return a
	}
	return b
}

func (c If) Int(a, b int) int {
	if c {
		return a
	}
	return b
}

func (c If) Level(a, b logrus.Level) logrus.Level {
	if c {
		return a
	}
	return b
}