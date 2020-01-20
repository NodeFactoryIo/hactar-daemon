package mock

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMockResponse(t *testing.T) {
	c := NewClient()

	c.MockResponse("RPC.Test", "test-response", "param")

	resp, err := c.Call("RPC.Test", "param")

	assert.Equal(t, "test-response", resp.Result)
	assert.Nil(t, resp.Error)
	assert.Nil(t, err)
	c.AssertExpectations(t)
}

func TestMockResponseWithMultipleParams(t *testing.T) {
	c := NewClient()

	c.MockResponse("RPC.Test", "test-response", "param1", "param2")

	resp, err := c.Call("RPC.Test", "param1", "param2")

	assert.Equal(t, "test-response", resp.Result)
	assert.Nil(t, resp.Error)
	assert.Nil(t, err)
	c.AssertExpectations(t)
}

func TestMockError(t *testing.T) {
	c := NewClient()

	c.MockError("RPC.Test", errors.New("test error"), "param1")

	resp, err := c.Call("RPC.Test", "param1")

	assert.Nil(t, resp.Result)
	if assert.Error(t, err) {
		assert.Equal(t, errors.New("test error"), err)
	}
	c.AssertExpectations(t)
}
