package http

import (
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/tkhamsila/backendtest/src/domains/doscg"
	"github.com/tkhamsila/backendtest/src/domains/doscg/mocks"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestDoSCGHandler_FindBCHandler(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/doscg/bc", strings.NewReader(`{"ans1": 23,"ans2": -21}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockUsecase := new(mocks.Usecase)
	h := DoSCGHandler{
		usecase: mockUsecase,
	}

	inputUsecase := bcInput{23, -21}
	expectRespUsecase := doscg.Bc{2, -42}

	mockUsecase.On("FindBC", inputUsecase.Ans1, inputUsecase.Ans2).Once().Return(expectRespUsecase)

	err := h.FindBCHandler(c)

	assert.NoError(t, err)
	mockUsecase.AssertExpectations(t)
}

func TestDoSCGHandler_FindBestWayToScgBangsueHandler(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/doscg/scgBangsueDirection", strings.NewReader(`{"start": "เซ็นทรัลเวิลด์+999%2F9+ถนน+พระรามที่+๑+แขวง+ปทุมวัน+เขตปทุมวัน+กรุงเทพมหานคร+10330"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockUsecase := new(mocks.Usecase)
	h := DoSCGHandler{
		usecase: mockUsecase,
	}

	inputUsecase := "เซ็นทรัลเวิลด์+999%2F9+ถนน+พระรามที่+๑+แขวง+ปทุมวัน+เขตปทุมวัน+กรุงเทพมหานคร+10330"
	expectRespUsecase := doscg.DirectionResp{}

	mockUsecase.On("FindBestWayToScgBangsue", inputUsecase).Once().Return(&expectRespUsecase, nil)

	err := h.FindBestWayToScgBangsueHandler(c)

	assert.NoError(t, err)
	mockUsecase.AssertExpectations(t)
}

func TestDoSCGHandler_FindXYZHandler(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/doscg/xyz", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockUsecase := new(mocks.Usecase)
	h := DoSCGHandler{
		usecase: mockUsecase,
	}

	expectRespUsecase := doscg.Xyz{}

	mockUsecase.On("FindXYZ").Once().Return(expectRespUsecase)

	err := h.FindXYZHandler(c)

	assert.NoError(t, err)
	mockUsecase.AssertExpectations(t)
}

func TestDoSCGHandler_AnswerCustomerHandler(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/doscg/answercustomer", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockUsecase := new(mocks.Usecase)
	mockBotService := new(mocks.BotService)
	h := DoSCGHandler{
		usecase: mockUsecase,
		bot: mockBotService,
	}

	expectRespUsecaseArg1 := doscg.AutomatedReplyMeassage
	expectRespUsecaseArg2 := doscg.TimeOutWhenBotAnswerDelay

	mockUsecase.On("AnswerCustomer").Once().Return(expectRespUsecaseArg1, expectRespUsecaseArg2)
	mockBotService.On("ReplyMessage", req, expectRespUsecaseArg1, expectRespUsecaseArg2).Once().Return(nil)

	err := h.AnswerCustomerHandler(c)

	assert.NoError(t, err)
	mockUsecase.AssertExpectations(t)
	mockBotService.AssertExpectations(t)
}