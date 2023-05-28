package keypress

import (
	"gamedev/config"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

var prevShot time.Time = time.Now().Add(-1 * time.Minute)

func isShoot(config *config.Config) bool {
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		if shoot := time.Now(); shoot.UnixMilli()-prevShot.UnixMilli() > config.BulletCooling {
			prevShot = shoot
			return true
		}
	}
	return false
}
