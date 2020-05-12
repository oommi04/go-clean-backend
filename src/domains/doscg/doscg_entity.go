package doscg

import "time"

type Xyz struct {
	X int `json:"x"`
	Y int `json:"y"`
	Z int `json:"z"`
}

type Bc struct {
	B int `json:"B"`
	C int `json:"C"`
}

type OverviewPolyline struct {
	Points string `json:"points"`
}

type Routes struct {
	OverviewPolyline OverviewPolyline `json:"overview_polyline"`
	Summary          string           `json:"summary"`
}

type DirectionResp struct {
	ErrorMessage string    `json:"error_message"`
	Routes       []*Routes `json:"routes"`
	Status       string    `json:"status"`
}

var (
	AutomatedReplyMeassage = "ขอบคุณสำหรับข้อความ ขณะนี้ทางเรายังไม่สามารถติดต่อได้ในขณะนี้ โปรดฝากคำถามหรือข้อสงสัย แล้วทางเราจะตอบกลับภายหลัง"
)

var (
	TimeOutWhenBotAnswerDelay = 3 * time.Second
)
