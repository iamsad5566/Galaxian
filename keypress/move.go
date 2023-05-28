package keypress

import (
	"gamedev/config"
	"gamedev/object/spaceship"

	"github.com/hajimehoshi/ebiten/v2"
)

func move(ship spaceship.SpaceShip, config *config.Config) error {
	keyMap := make(map[ebiten.Key]bool)
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		keyMap[ebiten.KeyLeft] = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		keyMap[ebiten.KeyRight] = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		keyMap[ebiten.KeyUp] = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		keyMap[ebiten.KeyDown] = true
	}
	for k, v := range keyMap {
		if v {
			// goroutine here would induce wired shaking
			processMap(k, ship, config)
		}
	}
	return nil
}

func processMap(key ebiten.Key, ship spaceship.SpaceShip, config *config.Config) {
	switch key {
	case ebiten.KeyLeft:
		if ship.X <= 0 {
			return
		}
		ship.X -= config.ShipSpeedFactor
	case ebiten.KeyRight:
		if ship.X+float64(ship) >= float64(config.ScreenWidth) {
			return
		}
		ship.X += config.ShipSpeedFactor
	case ebiten.KeyUp:
		if ship.Y <= 0 {
			return
		}
		ship.Y -= config.ShipSpeedFactor
	case ebiten.KeyDown:
		if ship.Y+float64(ship.Height) >= float64(config.ScreenHeight) {
			return
		}
		ship.Y += config.ShipSpeedFactor
	}
}
