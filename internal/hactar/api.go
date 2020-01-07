// Package hactar defines functions for communication with hactar backend
package hactar

import (
	log "github.com/sirupsen/logrus"
)

func Auth(username, password string) error {
	return nil
}

type Node struct {
	Token        string
	Url          string
	ActorAddress string
}

func SaveNode(node Node) {
	log.Info("Saving node information to hactar: ", node)
}
