package keypress

import (
	"gamedev/config"
	"gamedev/object/spaceship"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

var prevShot time.Time = time.Now().Add(-1 * time.Minute)

func shoot(ship spaceship.SpaceShip, config config.Config) bool {
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		if shoot := time.Now(); shoot.UnixMilli()-prevShot.UnixMilli() > config.BulletCooling {
			prevShot = shoot
			return true
		}
	}
	return false
}
