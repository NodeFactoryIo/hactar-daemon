package util

import (
	log "github.com/sirupsen/logrus"
	"net"
	"net/http"
)

// Get preferred outbound ip of this machine
func GetOutboundIP() net.IP {
	// get the local up address if it would connect to that target,
	// can be changed to any other IP address.
	conn, err := net.Dial("udp", "8.8.8.8:80") // this
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func HttpResponseStatus2XX(response *http.Response) bool {
	return response.StatusCode >= 200 && response.StatusCode < 300
}
