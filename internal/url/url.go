package url

import (
	"github.com/NodeFactoryIo/hactar-daemon/pkg/util"
)

func GetUrl() string {
	return util.GetOutboundIP().String()
}
