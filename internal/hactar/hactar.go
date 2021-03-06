package hactar

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/NodeFactoryIo/hactar-daemon/pkg/util"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"net/url"
	"os"
)

const (
	mediaType    = "application/json"
	authEndpoint = "/user/daemon/login"
)

type Client struct {
	// HTTP client used to communicate
	client *http.Client
	// Base URL for API requests.
	BaseURL *url.URL
	// JWT token
	Token string
	// Services used for communicating with the API
	Nodes     NodesService
	DiskInfo  DiskInfoService
	Blocks    BlocksService
	Miner     MinerService
	PastDeals PastDealsService
}

type NodeInfo struct {
	Address string `json:"address"`
	Url     string `json:"url"`
}

func NewAuthClient(email string, password string) (*Client, error) {
	c := NewClient(nil)
	// call auth endpoint for jwt token
	token, err := c.Auth(email, password)
	if err != nil {
		// failed on auth
		return nil, err
	}
	c.Token = token
	return c, nil
}

func NewClient(token interface{}) *Client {
	httpClient := http.DefaultClient

	baseUrl := viper.GetString("hactar.api-url")
	baseURL, _ := url.Parse(baseUrl)

	c := &Client{client: httpClient, BaseURL: baseURL}

	c.Nodes = &nodesServices{client: c}
	c.DiskInfo = &diskInfoService{client: c}
	c.Blocks = &blocksService{client: c}
	c.Miner = &minerService{client: c}
	c.PastDeals = &pastDealsService{client: c}

	if token != nil {
		c.Token = util.String(token)
	}

	return c
}

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

func (c *Client) Auth(email string, password string) (string, error) {
	body := &TokenRequest{
		Email:    email,
		Password: password,
	}
	request, err := c.NewRequest(http.MethodPost, authEndpoint, body)

	if err != nil {
		return "", err
	}

	tokenResponse := new(TokenResponse)
	response, err := c.Do(request, tokenResponse)

	if err != nil {
		return "", err
	}

	if response != nil && !util.HttpResponseStatus2XX(response) {
		return "", errors.New(fmt.Sprintf("Unable to authorize, server returned http status %s", response.Status))
	}

	return tokenResponse.Token, err
}

// NewRequest creates an API request. A relative URL can be provided in urlStr, which will be resolved to the
// BaseURL of the Client.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	u, err := c.BaseURL.Parse(c.BaseURL.Path + urlStr)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	if body != nil {
		err = json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", mediaType)
	req.Header.Add("Accept", mediaType)
	if c.Token != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.Token))
	}
	return req, nil
}

// Do sends an API request and returns the API response. The API response is JSON decoded and stored in the value
// pointed to by v, or returned as an error if an API error has occurred. If v implements the io.Writer interface,
// the raw response will be written to v, without attempting to decode it.
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		if rerr := resp.Body.Close(); err == nil {
			err = rerr
		}
	}()

	err = CheckResponse(resp)
	if err != nil {
		return resp, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, resp.Body)
			if err != nil {
				return nil, err
			}
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
			if err != nil {
				return nil, err
			}
		}
	}

	return resp, err
}

// An ErrorResponse reports the error caused by an API request
type ErrorResponse struct {
	// HTTP response that caused this error
	Response *http.Response

	// Error message
	Message string `json:"message"`
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v",
		r.Response.Request.Method, r.Response.Request.URL, r.Response.StatusCode, r.Message)
}

func CheckResponse(r *http.Response) error {
	if util.HttpResponseStatus2XX(r) {
		return nil
	}
	// read error response
	errorBody := util.ReaderToString(r.Body)
	errorResponse := &ErrorResponse{Response: r, Message: errorBody}
	// stop daemon app
	if r.StatusCode == http.StatusNotFound {
		log.Error(errorResponse)
		fmt.Print("You unsubscribed from hactar service.")
		os.Exit(1)
	}
	return errorResponse
}
