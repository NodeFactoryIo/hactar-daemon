package hactar

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

var (
	mux    *http.ServeMux
	client *Client
	server *httptest.Server
)

// setup function for all hactar client tests using client
func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	client = NewClient(nil)
	baseUrl, _ := url.Parse(server.URL)
	client.BaseURL = baseUrl
}

// teardown function for all hactar client tests using client
func teardown() {
	server.Close()
}

// helper function for testing default client setup
// checks if all client services initialized and base url set up properly
func testClientDefaults(t *testing.T, c *Client) {
	services := []string{
		"Nodes",
		"DiskInfo",
	}

	cp := reflect.ValueOf(c)
	cv := reflect.Indirect(cp)

	for _, s := range services {
		assert.NotNil(t, cv.FieldByName(s), fmt.Sprintf("client.%s shouldn't be nil", s))
	}

	assert.NotNil(t, c.BaseURL, "base url shouldn't be nil")
}

func TestNewClient(t *testing.T) {
	c := NewClient(nil)
	testClientDefaults(t, c)
}

func TestNewRequest(t *testing.T) {
	c := NewClient(nil)

	inURL, outURL := "/foo", "/foo"
	inBody, outBody := &DiskInfo{
		FreeSpace:    "1000",
		TakenSpace:   "1000",
		Node: NodeInfo{
			Address: "test-url",
			Url:     "test-address",
		},
	},
		`{"freeSpace":"1000","takenSpace":"1000",`+
			`"node":{"address":"test-url","url":"test-address"}}`+"\n"

	req, _ := c.NewRequest(http.MethodGet, inURL, inBody)

	assert.Equal(t, outURL, req.URL.String())

	body, _ := ioutil.ReadAll(req.Body)
	assert.Equal(t, outBody, string(body))
}

func TestNewAuthRequest(t *testing.T) {
	c := NewClient("test-token")

	inURL, outURL := "/foo", "/foo"
	inBody, outBody := &DiskInfo{
		FreeSpace:    "1000",
		TakenSpace:   "1000",
		Node: NodeInfo{
			Address: "test-url",
			Url:     "test-address",
		},
	},
		`{"freeSpace":"1000","takenSpace":"1000",`+
			`"node":{"address":"test-url","url":"test-address"}}`+"\n"

	req, _ := c.NewRequest(http.MethodGet, inURL, inBody)

	assert.Equal(t, outURL, req.URL.String())

	body, _ := ioutil.ReadAll(req.Body)
	assert.Equal(t, outBody, string(body))

	assert.Equal(t, "Bearer test-token", req.Header.Get("Authorization"))
}

func TestNewRequest_BadURL(t *testing.T) {
	c := NewClient(nil)
	_, err := c.NewRequest(http.MethodGet, ":", nil)
	assert.NotNil(t, err, "expected error on bad url")
	_, ok := err.(*url.Error)
	assert.True(t, ok)
	assert.Equal(t, "parse", err.(*url.Error).Op, "expected url parse error")
}

func TestDo(t *testing.T) {
	setup()
	defer teardown()

	type foo struct {
		A string
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		fmt.Fprint(w, `{"A":"a"}`)
	})

	req, _ := client.NewRequest(http.MethodGet, "/", nil)
	body := new(foo)
	_, err := client.Do(req, body)
	assert.Nil(t, err)

	expected := &foo{"a"}
	assert.Equal(t, body, expected)
}

func TestDo_HttpError(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Bad Request", 400)
	})

	req, _ := client.NewRequest(http.MethodGet, "/", nil)
	_, err := client.Do(req, nil)

	assert.NotNil(t, err)
}
