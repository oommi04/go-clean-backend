package doscg

import "errors"

var (
	ErrorUnableCreateRequest                    = errors.New("unable create request from Path")
	ErrorUnableRequestGoogleDirection           = errors.New("unable request google direction")
	ErrorUnableFindDirectionStartAndDestination = errors.New("unable find direction between start and destination")
	ErrorAPIKeyInvalid                          = errors.New("The provided API key is invalid")
)
