package google

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"net/http"
	"time"
)

func googleDirectionPath(start, destination string) string{
	return "/directions/json?origin=" + start + "&destination=" + destination + "&mode=transit"
}

type OverviewPolyline struct {
	Points string `json:"points"`
}

type Routes struct {
	OverviewPolyline OverviewPolyline `json:"overview_polyline"`
	Summary       string        `json:"summary"`
}

type DirectionResp struct {
	ErrorMessage string `json:"error_message"`
	Routes []*Routes`json:"routes"`
	Status string `json:"status"`
}


func (client *GoogleClient) GetDirectionAPI(start, destination string) (*DirectionResp, error){
	fullPath := googleDirectionPath(start, destination)
	req, err := client.buildGetRequest(fullPath)

	if err != nil {
		return nil, ErrorUnableCreateRequest
	}

	resp := &fasthttp.Response{}

	err = client.httpClient.DoTimeout(req, resp, time.Duration(client.timeout)*time.Second)
	if err != nil {
		return nil, ErrorUnableRequestGoogleDirection
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, ErrorUnableRequestGoogleDirection
	}
	respBody := DirectionResp{}
	_ = json.Unmarshal(resp.Body(), &respBody)

	if respBody.Status == "NOT_FOUND" {
		return nil, ErrorUnableFindDirectionStartAndDestination
	}

	if respBody.Status == "REQUEST_DENIED" {
		return nil, ErrorAPIKeyInvalid
	}

	return &respBody, nil
}