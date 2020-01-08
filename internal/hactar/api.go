// Package hactar defines functions for communication with hactar backend
package hactar

import (
	log "github.com/sirupsen/logrus"
)

type Node struct {
	Token        string
	Url          string
	ActorAddress string
}

func SaveNode(node Node) {
	// TODO implement
	log.Info("Saving node information to hactar: ", node)
}

func Auth(username, password string) error {
	// TODO implement
	return nil
}
