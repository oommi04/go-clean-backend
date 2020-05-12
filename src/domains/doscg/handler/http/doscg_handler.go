package http

import (
	"github.com/labstack/echo"
	"github.com/tkhamsila/backendtest/src/domains/doscg"
	"github.com/tkhamsila/backendtest/src/domains/doscg/usecase"
	"github.com/tkhamsila/backendtest/src/utils/errorStatus"
	"net/http"
)

type bcInput struct {
	Ans1 int `json:"ans1"`
	Ans2 int `json:"ans2"`
}

type bestWayToScgBangsueInput struct {
	Start string `json:"start"`
}

type responseError struct {
	Message string `json:"message"`
}

type DoSCGHandler struct {
	usecase doscg.Usecase
	bot     doscg.BotService
}

func Init(e *echo.Echo, m doscg.MapService, b doscg.BotService) {
	u := usecase.DoSCGUsecase{
		MapService: m,
	}
	handler := &DoSCGHandler{
		usecase: u,
		bot:     b,
	}
	e.GET("/doscg/xyz", handler.FindXYZHandler)
	e.POST("/doscg/bc", handler.FindBCHandler)
	e.POST("/doscg/scgBangsueDirection", handler.FindBestWayToScgBangsueHandler)
	e.POST("/doscg/answercustomer", handler.AnswerCustomerHandler)
}

func (h *DoSCGHandler) FindXYZHandler(c echo.Context) error {
	resp := h.usecase.FindXYZ()

	return c.JSON(http.StatusOK, resp)
}

func (h *DoSCGHandler) FindBCHandler(c echo.Context) error {
	body := bcInput{}

	if err := c.Bind(&body); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	resp := h.usecase.FindBC(body.Ans1, body.Ans2)

	return c.JSON(http.StatusOK, resp)
}

func (h *DoSCGHandler) FindBestWayToScgBangsueHandler(c echo.Context) error {
	body := bestWayToScgBangsueInput{}

	if err := c.Bind(&body); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	resp, err := h.usecase.FindBestWayToScgBangsue(body.Start)

	if err != nil {
		return c.JSON(errorStatus.GetStatusCode(err), responseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *DoSCGHandler) AnswerCustomerHandler(c echo.Context) error {
	replyMessage, timeout := h.usecase.AnswerCustomer()
	err := h.bot.ReplyMessage(c.Request(), replyMessage, timeout)

	if err != nil {
		return c.JSON(errorStatus.GetStatusCode(err), responseError{Message: err.Error()})
	}

	return c.JSON(200, "OK")
}
