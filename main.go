package main

import (
	"github.com/tkhamsila/backendtest/src/app"
	"github.com/tkhamsila/backendtest/src/configs"
	"github.com/tkhamsila/backendtest/src/domains/doscg"
)

func main() {
	cfg := configs.New()

	googleClient := app.SetupGoogle(cfg)

	e := app.SetupHttp(cfg)

	doscg.Init(e, googleClient)

	e.Logger.Fatal(e.Start(":" + cfg.PORT))
}
