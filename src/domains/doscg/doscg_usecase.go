package doscg

import (
	"github.com/tkhamsila/backendtest/src/external/google"
)

type DoSCGUsecase struct {
	GoogleService google.Service
}

func NewDoSCGUsecase() DoSCGUsecase {
	return DoSCGUsecase{}
}

func (u DoSCGUsecase) FindXYZ() xyz {
	return xyz{5, 2, 3}
}

func (u DoSCGUsecase) FindBC(ans1 int, ans2 int) bc {
	A := 21
	B := ans1 - A
	C := ans2 - A
	return bc{B, C}
}

func (u DoSCGUsecase) FindBestWayToScgBangsue(start string) (*google.DirectionResp,error){
	destination := "SCG+สำนักงานใหญ่+บางซื่อ+1+ซอย+ปูนซีเมนต์ไทย+แขวง+บางซื่อ+เขตบางซื่อ+กรุงเทพมหานคร+10800"

	resp, err := u.GoogleService.GetDirectionAPI(start, destination)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
