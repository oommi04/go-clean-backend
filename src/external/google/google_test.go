package google

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/tkhamsila/backendtest/src/drivers/fastHttpDriver/mocks"
)

type GoogleServiceSuite struct {
	suite.Suite

	fastHttp           *mocks.FastHttpClient
	service            *GoogleClient
	integrationService *GoogleClient
}

func Test_Google_Service_Suite(t *testing.T) {
	suite.Run(t, new(GoogleServiceSuite))
}

func (suite *GoogleServiceSuite) SetupTest() {
	suite.fastHttp = &mocks.FastHttpClient{}
	suite.service = New("key-test", "localhost", 5).setHttpClient(suite.fastHttp)
	suite.integrationService = New("key-test-staging", "https://maps.googleapis.com/maps/api", 5)
}
