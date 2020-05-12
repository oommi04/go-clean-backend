package errorStatus

import (
	"github.com/tkhamsila/backendtest/src/domains/doscg"
	"net/http"
)

func GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	switch err {
	case doscg.ErrorUnableCreateRequest:
	case doscg.ErrorAPIKeyInvalid:
	case doscg.ErrorUnableRequestGoogleDirection:
		return http.StatusInternalServerError
	case doscg.ErrorUnableFindDirectionStartAndDestination:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
	return http.StatusInternalServerError
}
