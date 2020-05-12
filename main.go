package main

import (
	"github.com/tkhamsila/backendtest/src/app"
	"github.com/tkhamsila/backendtest/src/configs"
	doscgHttp "github.com/tkhamsila/backendtest/src/domains/doscg/handler/http"
)

func main() {
	cfg := configs.New()

	googleClient := app.SetupGoogle(cfg)

	e := app.SetupHttp(cfg)

	lineBot := app.SetupLineBot(cfg)

	doscgHttp.Init(e, googleClient,lineBot)

	e.Logger.Fatal(e.Start(":" + cfg.PORT))
}