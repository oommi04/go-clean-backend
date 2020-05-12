package doscg

import (
	"net/http"
	"time"
)

//go:generate mockery -name=Usecase
type Usecase interface {
	FindXYZ() Xyz
	FindBC(ans1 int, ans2 int) Bc
	FindBestWayToScgBangsue(start string) (*DirectionResp, error)
	AnswerCustomer() (string, time.Duration)
}

//go:generate mockery -name=MapService
type MapService interface {
	GetDirectionAPI(start string, destination string) (*DirectionResp, error)
}

//go:generate mockery -name=BotService
type BotService interface {
	ReplyMessage(r *http.Request, messageReply string, timeout time.Duration) error
}
