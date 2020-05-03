package google

import (
	"github.com/tkhamsila/backendtest/src/drivers/fastHttpDriver"
	"github.com/valyala/fasthttp"
)

//go:generate mockery -name=Service
type Service interface {
	GetDirectionAPI(start string, destination string) (*DirectionResp, error)
}

type GoogleClient struct {
	httpClient fastHttpDriver.FastHttpClient

	endpoint string
	key      string
	timeout  int
}

func New(key, endpoint string, timeout int) *GoogleClient {
	return &GoogleClient{
		endpoint:   endpoint,
		httpClient: &fasthttp.Client{},
		key:        key,
		timeout:    timeout,
	}
}

func (client *GoogleClient) setHttpClient(httpClient fastHttpDriver.FastHttpClient) *GoogleClient {
	client.httpClient = httpClient

	return client
}

func (client *GoogleClient) buildGetRequest(path string) (*fasthttp.Request, error) {
	req := &fasthttp.Request{}

	req.Header.SetMethod("GET")
	req.Header.SetContentType("application/json")

	fullEndpoint := client.endpoint + path + "&key=" + client.key

	req.SetRequestURI(fullEndpoint)

	return req, nil
}