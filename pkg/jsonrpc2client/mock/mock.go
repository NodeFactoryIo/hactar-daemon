package mock

import (
	"github.com/stretchr/testify/mock"
	"github.com/ybbus/jsonrpc"
)

type JsonRpc2Client struct {
	mock.Mock
}

func NewClient() *JsonRpc2Client {
	return &JsonRpc2Client{}
}

func (c *JsonRpc2Client) Call(method string, params ...interface{}) (*jsonrpc.RPCResponse, error) {
	args := c.Called(method, params)
	resp := &jsonrpc.RPCResponse{
		JSONRPC: "2.0",
		Result:  args.Get(0),
		Error:   nil,
		ID:      0,
	}
	return resp, args.Error(1)
}

func (c *JsonRpc2Client) MockResponse(method string, response interface{}, params ...interface{}) {
	call := c.On("Call", method, params)
	call.Once()
	call.Return(response, nil)
}

func (c *JsonRpc2Client) MockError(method string, err error, params ...interface{}) {
	call := c.On("Call", method, params)
	call.Once()
	call.Return(nil, err)
}