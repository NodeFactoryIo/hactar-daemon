package hactar

import (
	"net/http"
)

type NodesService interface {
	Add(node Node) (*Node, *http.Response, error)
}

type nodesServices struct {
	client *Client
}

type Node struct {
	Token        string `json:"token"`
	Node NodeInfo `json:"nodeInfo"`
}

const (
	NodePath = "/user/node"
)

func (ns *nodesServices) Add(node Node) (*Node, *http.Response, error) {
	request, err := ns.client.NewRequest(http.MethodPost, NodePath, node)

	if err != nil {
		return nil, nil, err
	}

	return sendSingleNodeRequest(request, ns)
}

func sendSingleNodeRequest(request *http.Request, ns *nodesServices) (*Node, *http.Response, error) {
	root := new(Node)

	response, err := ns.client.Do(request, root)

	if err != nil {
		return nil, response, err
	}

	return root, response, err
}
