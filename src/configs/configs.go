package configs

import (
	"github.com/joho/godotenv"
	"os"
)

type Configs struct {
	GoogleMapKey         string
	PORT string
}

func New () *Configs {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	return &Configs{
		os.Getenv("GOOGLEMAPKEY"),
		os.Getenv("PORT"),
	}
}