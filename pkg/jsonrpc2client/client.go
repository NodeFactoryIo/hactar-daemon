package jsonrpc2client

import (
	"github.com/NodeFactoryIo/hactar-daemon/internal/token"
	"github.com/ybbus/jsonrpc"
)

type Client interface {
	Call(method string, params ...interface{}) (*jsonrpc.RPCResponse, error)
}

type client struct {
	baseURL   string
	rpcClient jsonrpc.RPCClient
}

func NewClient(rpcurl string) *client {
	c := &client{}
	c.baseURL = rpcurl
	c.rpcClient = jsonrpc.NewClientWithOpts(c.baseURL, &jsonrpc.RPCClientOpts{
		CustomHeaders: map[string]string{
			"Authorization": "Bearer " + token.ReadTokenFromFile(),
		},
	})
	return c
}

func (c *client) Call(method string, params ...interface{}) (*jsonrpc.RPCResponse, error) {
	return c.rpcClient.Call(method, params)
}
