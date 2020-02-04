package hactar

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestMinerService_SendMinerInfo(t *testing.T) {
	setup()
	defer teardown()

	minerInfoRequest := &MinerInfo{
		Version:    "test-version",
		SectorSize: 1000,
		MinerPower: 100,
		TotalPower: 200,
		Node: NodeInfo{
			Address: "test-url",
			Url:     "test-address",
		},
	}

	mux.HandleFunc(SendMinerInfoPath, func(w http.ResponseWriter, r *http.Request) {
		v := new(MinerInfo)
		err := json.NewDecoder(r.Body).Decode(v)
		// assert valid request
		assert.Nil(t, err)
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, minerInfoRequest, v)

		resp, _ := json.Marshal(minerInfoRequest)
		_, _ = fmt.Fprintf(w, fmt.Sprintf(`%s`, string(resp)))
	})

	response, err := client.Miner.SendMinerInfo(*minerInfoRequest)
	// assert valid response
	assert.NotNil(t, response)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)
}
