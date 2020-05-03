package doscg

import "github.com/tkhamsila/backendtest/src/external/google"

//go:generate mockery -name=Usecase
type Usecase interface {
	FindXYZ() xyz
	FindBC(ans1 int, ans2 int) bc
	FindBestWayToScgBangsue(start string) (*google.DirectionResp,error)
}
