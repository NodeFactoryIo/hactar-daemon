package hactar

import (
	"encoding/json"
	"fmt"
	"github.com/NodeFactoryIo/hactar-daemon/internal/lotus"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestPastDealsService_SendPastDealsInfo(t *testing.T) {
	setup()
	defer teardown()

	addRequest := PastDealsInfo{
		Deals: []lotus.PastDealResponse{
			{
				Cid:      "test-cid-1",
				State:    1,
				Size:     "1234",
				Provider: "test-provider",
				Price:    "1234",
				Duration: 3,
			},
			{
				Cid:      "test-cid-2",
				State:    2,
				Size:     "2345",
				Provider: "test-provider-2",
				Price:    "1234",
				Duration: 3,
			},
		},
		Node: NodeInfo{
			Address: "test-url",
			Url:     "test-address",
		},
	}

	mux.HandleFunc(SendPastDealsInfoPath, func(w http.ResponseWriter, r *http.Request) {
		v := new(PastDealsInfo)
		err := json.NewDecoder(r.Body).Decode(v)
		// assert valid request
		assert.Nil(t, err)
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, addRequest, *v)

		resp, _ := json.Marshal(addRequest)
		_, _ = fmt.Fprintf(w, fmt.Sprintf(`%s`, string(resp)))
	})

	response, err := client.PastDeals.SendPastDealsInfo(addRequest)
	// assert valid response
	assert.NotNil(t, response)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)
}
