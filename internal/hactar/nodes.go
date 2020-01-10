package hactar

import "net/http"

type NodesService interface {
	Add(node Node) (*http.Response, error)
}

type nodesServices struct {
	client *Client
}

type Node struct {
	Token        string
	Url          string
	ActorAddress string
}

const (
	AddPath = ""
)

func (ns *nodesServices) Add(node Node) (*Node, *http.Response, error) {
	request, err := ns.client.NewRequest(nil, http.MethodPost, AddPath, node)

	if err != nil {
		return nil, nil, err
	}

	root := new(Node)

	response, err := ns.client.Do(nil, request, root)

	if err != nil {
		return nil, response, err
	}

	return root, response, err
}

