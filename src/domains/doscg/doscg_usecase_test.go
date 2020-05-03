package doscg

import (
	"github.com/stretchr/testify/assert"
	"github.com/tkhamsila/backendtest/src/external/google"
	"github.com/tkhamsila/backendtest/src/external/google/mocks"
	"testing"
)

func TestDoSCGUsecase_FindBC(t *testing.T) {
	u := DoSCGUsecase{}
	resp := u.FindBC(23,-21)
	expectResp := bc{2, -42}
	assert.Equal(t, expectResp, resp)
}

func TestDoSCGUsecase_FindBestWayToScgBangsue(t *testing.T) {
	mockGoogleClient := new(mocks.Service)
	u := DoSCGUsecase{
		mockGoogleClient,
	}

	expectRespGoogle := google.DirectionResp{}

	t.Run("success", func(t *testing.T){
		start := "cmu"
		destination := "SCG+สำนักงานใหญ่+บางซื่อ+1+ซอย+ปูนซีเมนต์ไทย+แขวง+บางซื่อ+เขตบางซื่อ+กรุงเทพมหานคร+10800"

		mockGoogleClient.On("GetDirectionAPI", start, destination).Once().Return(&expectRespGoogle, nil)

		_, err := u.FindBestWayToScgBangsue(start)

		assert.NoError(t,err)
	})

	t.Run("error", func(t *testing.T){
		start := "cmu"
		destination := "SCG+สำนักงานใหญ่+บางซื่อ+1+ซอย+ปูนซีเมนต์ไทย+แขวง+บางซื่อ+เขตบางซื่อ+กรุงเทพมหานคร+10800"

		mockGoogleClient.On("GetDirectionAPI", start, destination).Once().Return(nil, google.ErrorUnableFindDirectionStartAndDestination)

		_, err := u.FindBestWayToScgBangsue(start)

		assert.Error(t, err)
		assert.Equal(t, google.ErrorUnableFindDirectionStartAndDestination, err)
	})

}