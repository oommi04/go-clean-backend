package google

func (suite *GoogleServiceSuite) TestGoogleClient_Integration_directionAPI_Success() {
	centralWorld := "เซ็นทรัลเวิลด์+999%2F9+ถนน+พระรามที่+๑+แขวง+ปทุมวัน+เขตปทุมวัน+กรุงเทพมหานคร+10330"
	scgBangsue := "SCG+สำนักงานใหญ่+บางซื่อ+1+ซอย+ปูนซีเมนต์ไทย+แขวง+บางซื่อ+เขตบางซื่อ+กรุงเทพมหานคร+10800"
	_, err := suite.integrationService.GetDirectionAPI(centralWorld, scgBangsue)
	suite.NoError(err)
	suite.fastHttp.AssertExpectations(suite.T())
}

func (suite *GoogleServiceSuite) TestGoogleClient_Integration_directionAPI_ErrorDirectionStartAndDestination() {
	centralWorld := "เซ็นทรัลเวิลด์+999%2F9+ถนน+พระรามที่+๑+แขวง+ปทุมวัน+เขตปทุมวัน+กรุงเทพมหานคร+10330"
	scgBangsue := "gggg"
	_, err := suite.integrationService.GetDirectionAPI(centralWorld, scgBangsue)
	suite.Error(err)
	suite.Equal(ErrorUnableFindDirectionStartAndDestination, err)
	suite.fastHttp.AssertExpectations(suite.T())
}