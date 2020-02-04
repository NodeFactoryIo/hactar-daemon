package hactar

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestDiskInfoService_SendDiskInfo(t *testing.T) {
	setup()
	defer teardown()

	createRequest := &DiskInfo{
		FreeSpace:    "1000",
		TakenSpace:   "2000",
		Node: NodeInfo{
			Address: "test-url",
			Url:     "test-address",
		},
	}

	mux.HandleFunc(DiskInfoPath, func(w http.ResponseWriter, r *http.Request) {
		v := new(DiskInfo)
		err := json.NewDecoder(r.Body).Decode(v)
		// assert valid request
		assert.Nil(t, err)
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, createRequest, v)

		resp, _ := json.Marshal(createRequest)
		_, _ = fmt.Fprintf(w, fmt.Sprintf(`%s`, string(resp)))
	})
	response, err := client.DiskInfo.SendDiskInfo(*createRequest)
	// assert valid response
	assert.NotNil(t, response)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)
}
