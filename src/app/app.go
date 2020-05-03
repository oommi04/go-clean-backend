package app

import (
	"github.com/labstack/echo"
	"github.com/tkhamsila/backendtest/src/configs"
	"github.com/tkhamsila/backendtest/src/drivers/echoDriver"
	"github.com/tkhamsila/backendtest/src/external/google"
	"net/http"
)

func SetupGoogle(c* configs.Configs) *google.GoogleClient{
	endPoint := "https://maps.googleapis.com/maps/api"
	googleClient := google.New(c.GoogleMapKey, endPoint, 5)
	return googleClient
}

func SetupHttp(c* configs.Configs) *echo.Echo{
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	middlewareEcho := echoDriver.InitMiddleware()
	e.Use(middlewareEcho.CORS)

	return e
}