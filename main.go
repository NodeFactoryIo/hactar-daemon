package main

import (
	"fmt"
	"github.com/NodeFactoryIo/hactar-daemon/cmd"
	"os"
)

func main() {
	must(cmd.RootCmd.Execute())
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
