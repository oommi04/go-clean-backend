package google

import (
	"encoding/json"
	"github.com/stretchr/testify/mock"
	"github.com/tkhamsila/backendtest/src/domains/doscg"
	"github.com/valyala/fasthttp"
	"net/http"
	"time"
)

func (suite *GoogleServiceSuite) TestGoogleClient_directionAPI_Success() {
	centralWorld := "เซ็นทรัลเวิลด์+999%2F9+ถนน+พระรามที่+๑+แขวง+ปทุมวัน+เขตปทุมวัน+กรุงเทพมหานคร+10330"
	scgBangsue := "SCG+สำนักงานใหญ่+บางซื่อ+1+ซอย+ปูนซีเมนต์ไทย+แขวง+บางซื่อ+เขตบางซื่อ+กรุงเทพมหานคร+10800"
	fullPath := googleDirectionPath(centralWorld, scgBangsue)

	req, _ := suite.service.buildGetRequest(fullPath)
	resp := &fasthttp.Response{}

	suite.fastHttp.On("DoTimeout", req, resp, 5*time.Second).Once().Run(func(args mock.Arguments) {
		resp := args[1].(*fasthttp.Response)
		resp.SetStatusCode(http.StatusOK)

		respBody := doscg.DirectionResp{Status: "OK", Routes: []*doscg.Routes{
			{
				Summary: "ทางพิเศษศรีรัช",
				OverviewPolyline: doscg.OverviewPolyline{
					Points: "iz{rAeysdRvDZjAFrADVETuBRwB`@yEr@sHn@mGf@sF^uDj@mFl@}FBYsBRyDRaA@oCDc@B_@EMImER{DXiCLqBJuBN_AF_CZsA`@s@XIDYNqAn@q@f@q@r@u@~@m@|@q@lAmAjBc@n@KHo@l@aBdAcChAEBOFk@Pg@ZuBbA_Af@c@Vi@X[VSPCBi@`@{@f@YJs@L_@?c@G}@][WMSYs@GMCY@q@DYDQZs@t@s@`CaBp@a@pBoAv@q@JM\\\\m@Hu@Cq@CUGKSU]a@w@Wa@G[A]Bi@La@JmAl@g@P_Ar@mErD{@p@eA|@{FbFwDjDWX[l@{@hBwBlG}@hBo@lAiAzA}A|Ag@`@aBfAwAv@yBlAu@b@_Ar@_A~@y@dAi@|@KRCBuAxC_ArBwBdFgCrF_@x@ORUf@{@~A_@`@c@\\\\QHWLg@Lk@F}@BOA}@OmCs@aB[UEqAKgACwCJaABgB@gCE_AGiBYy@Se@GUAmCy@cA[qA]qGeBqBo@_@CyEqAgD_AK?S?]DUJUPe@j@QPq@Xe@D[C_@MSOOUKUE]?OFs@ZqAMEMEaAUqDgAmE}A[KqBs@u@QiDcAoDgAcAYy@US?KCI~@KzA_@bDs@pEMnASWyAY}Aa@kCy@sFmBgCu@eBm@{@g@a@YcA_@OG{@]mPwEqK_Dm@MSDs@vA[jAEIKGa@MyCdGANw@xASYa@QSGOTQLW?cBk@Ty@g@Q"}}}}
		creatorJSON, _ := json.Marshal(respBody)
		resp.SetBody(creatorJSON)
	}).Return(nil)

	_, err := suite.service.GetDirectionAPI(centralWorld, scgBangsue)

	suite.NoError(err)
	suite.fastHttp.AssertExpectations(suite.T())
}

func (suite *GoogleServiceSuite) TestGoogleClient_directionAPI_Error() {
	centralWorld := "เซ็นทรัลเวิลด์+999%2F9+ถนน+พระรามที่+๑+แขวง+ปทุมวัน+เขตปทุมวัน+กรุงเทพมหานคร+10330"
	scgBangsue := "SCG+สำนักงานใหญ่+บางซื่อ+1+ซอย+ปูนซีเมนต์ไทย+แขวง+บางซื่อ+เขตบางซื่อ+กรุงเทพมหานคร+10800"
	fullPath := googleDirectionPath(centralWorld, scgBangsue)

	req, _ := suite.service.buildGetRequest(fullPath)
	resp := &fasthttp.Response{}

	suite.fastHttp.On("DoTimeout", req, resp, 5*time.Second).Once().Run(func(args mock.Arguments) {
		resp := args[1].(*fasthttp.Response)
		resp.SetStatusCode(http.StatusBadRequest)
	}).Return(nil)

	_, err := suite.service.GetDirectionAPI(centralWorld, scgBangsue)

	suite.Error(err)
	suite.Equal(doscg.ErrorUnableRequestGoogleDirection, err)
	suite.fastHttp.AssertExpectations(suite.T())
}
