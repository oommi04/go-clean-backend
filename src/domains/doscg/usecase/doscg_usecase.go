package usecase

import (
	"github.com/tkhamsila/backendtest/src/domains/doscg"
	"time"
)

type DoSCGUsecase struct {
	MapService doscg.MapService
}

func NewDoSCGUsecase(m doscg.MapService) *DoSCGUsecase {
	return &DoSCGUsecase{MapService: m}
}

func (u DoSCGUsecase) FindXYZ() doscg.Xyz {
	z := 23 + 2 * (6 - 1)
	y := 5 - 2 * (2 - 1)
	x := y - 2 * (1 - 1)
	return doscg.Xyz{x, y, z}
}

func (u DoSCGUsecase) FindBC(ans1 int, ans2 int) doscg.Bc {
	A := 21
	B := ans1 - A
	C := ans2 - A
	return doscg.Bc{B, C}
}

func (u DoSCGUsecase) FindBestWayToScgBangsue(start string) (*doscg.DirectionResp, error) {
	destination := "SCG+สำนักงานใหญ่+บางซื่อ+1+ซอย+ปูนซีเมนต์ไทย+แขวง+บางซื่อ+เขตบางซื่อ+กรุงเทพมหานคร+10800"

	resp, err := u.MapService.GetDirectionAPI(start, destination)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (u DoSCGUsecase) AnswerCustomer() (string, time.Duration) {
	return doscg.AutomatedReplyMeassage, doscg.TimeOutWhenBotAnswerDelay
}
