package url

import (
	"fmt"
	"github.com/NodeFactoryIo/hactar-daemon/pkg/util"
)

func DisplayUrl() {
	fmt.Printf("Node URL:\n%s\n", GetUrl())
}

func GetUrl() string {
	return util.GetOutboundIP().String()
}
