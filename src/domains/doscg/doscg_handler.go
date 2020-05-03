package doscg

import (
	"github.com/labstack/echo"
	"github.com/tkhamsila/backendtest/src/external/google"
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
	usecase Usecase
}

func Init(e *echo.Echo, g google.Service) {
	u := DoSCGUsecase{
		 GoogleService: g,
	}
	handler := &DoSCGHandler{
		usecase: u,
	}
	e.GET("/doscg/xyz", handler.FindXYZHandler)
	e.POST("/doscg/bc",handler.FindBCHandler)
	e.POST("/doscg/scgBangsueDirection", handler.FindBestWayToScgBangsueHandler)
}

func(h *DoSCGHandler) FindXYZHandler(c echo.Context) error{
	resp := h.usecase.FindXYZ()

	return c.JSON(http.StatusOK, resp)
}

func (h *DoSCGHandler) FindBCHandler(c echo.Context) error {
	body := bcInput{}

	if err := c.Bind(&body); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	resp := h.usecase.FindBC(body.Ans1,body.Ans2)

	return  c.JSON(http.StatusOK, resp)
}

func (h *DoSCGHandler) FindBestWayToScgBangsueHandler(c echo.Context) error{
	body := bestWayToScgBangsueInput{}

	if err := c.Bind(&body); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	resp, err := h.usecase.FindBestWayToScgBangsue(body.Start)

	if err != nil {
		return c.JSON(http.StatusBadRequest, responseError{Message: err.Error()})
	}


	return  c.JSON(http.StatusOK, resp)
}
