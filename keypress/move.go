package keypress

import (
	"gamedev/config"
	"gamedev/object/spaceship"

	"github.com/hajimehoshi/ebiten/v2"
)

func move(ship spaceship.SpaceShip, config *config.Config) {
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
}

func processMap(key ebiten.Key, ship spaceship.SpaceShip, config *config.Config) {
	switch key {
	case ebiten.KeyLeft:
		if ship.GetAxisX() <= 0 {
			return
		}
		ship.SetAxisX(ship.GetAxisX() - config.ShipSpeedFactor)
	case ebiten.KeyRight:
		if ship.GetAxisX()+float64(ship.GetWidth()) >= float64(config.ScreenWidth) {
			return
		}
		ship.SetAxisX(ship.GetAxisX() + config.ShipSpeedFactor)
	case ebiten.KeyUp:
		if ship.GetAxisY() <= 0 {
			return
		}
		ship.SetAxisY(ship.GetAxisY() - config.ShipSpeedFactor)
	case ebiten.KeyDown:
		if ship.GetAxisY()+float64(ship.GetHeight()) >= float64(config.ScreenHeight) {
			return
		}
		ship.SetAxisY(ship.GetAxisY() + config.ShipSpeedFactor)
	}
}
