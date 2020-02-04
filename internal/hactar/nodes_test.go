package hactar

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNodesServices_Add(t *testing.T) {
	setup()
	defer teardown()

	createRequest := &Node{
		Token: "test-token",
		Node: NodeInfo{
			Address: "test-url",
			Url:     "test-address",
		},
	}

	mux.HandleFunc(NodePath, func(w http.ResponseWriter, r *http.Request) {
		v := new(Node)
		err := json.NewDecoder(r.Body).Decode(v)
		// assert valid request
		assert.Nil(t, err)
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, createRequest, v)

		resp, _ := json.Marshal(createRequest)
		_, _ = fmt.Fprintf(w, fmt.Sprintf(`%s`, string(resp)))
	})
	node, response, err := client.Nodes.Add(*createRequest)
	// assert valid response
	assert.NotNil(t, response)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, createRequest, node)
}
