package usecase

import (
	"github.com/stretchr/testify/assert"
	"github.com/tkhamsila/backendtest/src/domains/doscg"
	"github.com/tkhamsila/backendtest/src/domains/doscg/mocks"
	"testing"
)

func TestDoSCGUsecase_FindXYZ(t *testing.T) {
	u := DoSCGUsecase{}
	resp := u.FindXYZ()
	expectResp := doscg.Xyz{3,3,33}
	assert.Equal(t, expectResp, resp)
}

func TestDoSCGUsecase_FindBC(t *testing.T) {
	u := DoSCGUsecase{}
	resp := u.FindBC(23, -21)
	expectResp := doscg.Bc{2, -42}
	assert.Equal(t, expectResp, resp)
}

func TestDoSCGUsecase_FindBestWayToScgBangsue(t *testing.T) {
	mockMapService := new(mocks.MapService)
	u := DoSCGUsecase{
		mockMapService,
	}

	expectRespGoogle := doscg.DirectionResp{}

	start := "cmu"
	destination := "SCG+สำนักงานใหญ่+บางซื่อ+1+ซอย+ปูนซีเมนต์ไทย+แขวง+บางซื่อ+เขตบางซื่อ+กรุงเทพมหานคร+10800"

	t.Run("success", func(t *testing.T) {
		mockMapService.On("GetDirectionAPI", start, destination).Once().Return(&expectRespGoogle, nil)

		_, err := u.FindBestWayToScgBangsue(start)

		assert.NoError(t, err)
		mockMapService.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockMapService.On("GetDirectionAPI", start, destination).Once().Return(nil, doscg.ErrorUnableFindDirectionStartAndDestination)

		_, err := u.FindBestWayToScgBangsue(start)

		assert.Error(t, err)
		assert.Equal(t, doscg.ErrorUnableFindDirectionStartAndDestination, err)
		mockMapService.AssertExpectations(t)
	})
}
