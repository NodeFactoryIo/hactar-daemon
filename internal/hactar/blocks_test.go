package hactar

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestBlocksService_AddMiningReward(t *testing.T) {
	setup()
	defer teardown()

	addRequest := []Block{
		*(&Block{
			Cid: "test-cid-1",
			Node: NodeInfo{
				Address: "test-url",
				Url:     "test-address",
			},
		}),
		*(&Block{
			Cid: "test-cid-2",
			Node: NodeInfo{
				Address: "test-url",
				Url:     "test-address",
			},
		}),
	}

	mux.HandleFunc(AddBlockRewardPath, func(w http.ResponseWriter, r *http.Request) {
		v := new([]Block)
		err := json.NewDecoder(r.Body).Decode(v)
		// assert valid request
		assert.Nil(t, err)
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, addRequest, *v)

		resp, _ := json.Marshal(addRequest)
		_, _ = fmt.Fprintf(w, fmt.Sprintf(`%s`, string(resp)))
	})

	response, err := client.Blocks.AddMiningReward(addRequest)
	// assert valid response
	assert.NotNil(t, response)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)
}
