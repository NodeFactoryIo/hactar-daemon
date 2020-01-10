//+build mage

package main

import (
	"github.com/magefile/mage/sh"
	"os"
)

const (
	packageName  = "github.com/NodeFactoryIo/hactar-deamon"
)

var goexe = "go"

func init() {
	if exe := os.Getenv("GOEXE"); exe != "" {
		goexe = exe
	}

	// We want to use Go 1.11 modules even if the source lives inside GOPATH.
	// The default is "auto".
	os.Setenv("GO111MODULE", "on")
}

func Build() error {
	return sh.Run("go", "install", "./...")
}