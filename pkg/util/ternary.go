package util

import "github.com/sirupsen/logrus"

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

func (c If) Float32(a, b float32) float32 {
	if c {
		return a
	}
	return b
}

func (c If) Float64(a, b float64) float64 {
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
