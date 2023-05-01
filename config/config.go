package config

import (
	"encoding/json"
	"image/color"
	"log"
	"os"
)

type Config struct {
	ScreenHeight      int        `json:"screenHeight"`
	ScreenWidth       int        `json:"screenWidth"`
	Title             string     `json:"title"`
	BgColor           color.RGBA `json:"bgColor"`
	ShipSpeedFactor   float64    `json:"shipSpeedFactor"`
	BulletSpeedFactor float64    `json:"bulletSpeedFactor"`
	FPS               int        `json:"fps"`
	BulletCooling     int64      `json:"bulletCooling"`
}

func LoadConfig() *Config {
	f, err := os.Open("./config/config.json")
	if err != nil {
		log.Fatalf("os.Open failed: %v\n", err)
	}

	var cfg Config
	err = json.NewDecoder(f).Decode(&cfg)
	if err != nil {
		log.Fatalf("json.Decode failed: %v\n", err)
	}

	return &cfg
}
