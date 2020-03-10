package hactar

import (
	"github.com/getsentry/sentry-go"
	"net/http"
)

type NodesService interface {
	Add(node Node) (*Node, *http.Response, error)
	GetAllNodes() ([]NodeInfo, *http.Response, error)
	SendUptimeReport(report UptimeReport) (*http.Response, error)
	SendBalanceReport(report BalanceReport) (*http.Response, error)
}

type nodesServices struct {
	client *Client
}

type Node struct {
	Token string   `json:"token"`
	Node  NodeInfo `json:"nodeInfo"`
}

const (
	NodePath        = "/user/node"
	NodeUptimePath  = NodePath + "/uptime"
	NodeBalancePath = NodePath + "/balance"
)

func (ns *nodesServices) GetAllNodes() ([]NodeInfo, *http.Response, error) {
	request, err := ns.client.NewRequest(http.MethodGet, NodePath, nil)

	if err != nil {
		return nil, nil, err
	}

	root := new([]NodeInfo)

	response, err := ns.client.Do(request, root)

	if err != nil {
		sentry.CaptureException(err)
		return nil, response, err
	}

	return *root, response, err
}

type UptimeReport struct {
	IsWorking bool     `json:"isWorking"`
	Node      NodeInfo `json:"nodeInfo"`
}

func (ns *nodesServices) SendUptimeReport(report UptimeReport) (*http.Response, error) {
	request, err := ns.client.NewRequest(http.MethodPost, NodeUptimePath, report)

	if err != nil {
		return nil, err
	}

	response, err := ns.client.Do(request, nil)

	if err != nil {
		return response, err
	}

	return response, err
}

type BalanceReport struct {
	Balance string   `json:"balance"`
	Node    NodeInfo `json:"nodeInfo"`
}

func (ns *nodesServices) SendBalanceReport(report BalanceReport) (*http.Response, error) {
	request, err := ns.client.NewRequest(http.MethodPost, NodeBalancePath, report)

	if err != nil {
		return nil, err
	}

	response, err := ns.client.Do(request, nil)

	if err != nil {
		sentry.CaptureException(err)
		return response, err
	}

	return response, err
}

func (ns *nodesServices) Add(node Node) (*Node, *http.Response, error) {
	request, err := ns.client.NewRequest(http.MethodPost, NodePath, node)

	if err != nil {
		sentry.CaptureException(err)
		return nil, nil, err
	}

	return sendSingleNodeRequest(request, ns)
}

func sendSingleNodeRequest(request *http.Request, ns *nodesServices) (*Node, *http.Response, error) {
	root := new(Node)

	response, err := ns.client.Do(request, root)

	if err != nil {
		sentry.CaptureException(err)
		return nil, response, err
	}

	return root, response, err
}
