package mock

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMockResponse(t *testing.T) {
	c := NewClient()

	c.MockResponse("RPC.Test", "world", "hello")

	resp, err := c.Call("RPC.Test", "hello")

	assert.Equal(t, "world", resp.Result)
	assert.Equal(t, nil, err)
	c.AssertExpectations(t)
}

func TestMockResponseWithMultipleParams(t *testing.T) {
	c := NewClient()

	c.MockResponse("RPC.Test", "world", "hello", "world")

	resp, err := c.Call("RPC.Test", "hello", "world")

	assert.Equal(t, "world", resp.Result)
	assert.Equal(t, nil, err)
	c.AssertExpectations(t)
}

func TestMockError(t *testing.T) {
	c := NewClient()

	c.MockError("RPC.Test", errors.New("test error"), "hello")

	resp, err := c.Call("RPC.Test", "hello")

	assert.Equal(t, nil, resp.Result)
	if assert.Error(t, err) {
		assert.Equal(t, errors.New("test error"), err)
	}
	c.AssertExpectations(t)
}