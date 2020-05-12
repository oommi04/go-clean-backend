package google

import (
	"encoding/json"
	"github.com/tkhamsila/backendtest/src/domains/doscg"
	"github.com/valyala/fasthttp"
	"net/http"
	"time"
)

func googleDirectionPath(start, destination string) string {
	return "/directions/json?origin=" + start + "&destination=" + destination + "&mode=transit"
}

func (client *GoogleClient) GetDirectionAPI(start, destination string) (*doscg.DirectionResp, error) {
	fullPath := googleDirectionPath(start, destination)
	req, err := client.buildGetRequest(fullPath)

	if err != nil {
		return nil, doscg.ErrorUnableCreateRequest
	}

	resp := &fasthttp.Response{}

	err = client.httpClient.DoTimeout(req, resp, time.Duration(client.timeout)*time.Second)
	if err != nil {
		return nil, doscg.ErrorUnableRequestGoogleDirection
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, doscg.ErrorUnableRequestGoogleDirection
	}
	respBody := doscg.DirectionResp{}
	_ = json.Unmarshal(resp.Body(), &respBody)

	if respBody.Status == "NOT_FOUND" {
		return nil, doscg.ErrorUnableFindDirectionStartAndDestination
	}

	if respBody.Status == "REQUEST_DENIED" {
		return nil, doscg.ErrorAPIKeyInvalid
	}

	return &respBody, nil
}
