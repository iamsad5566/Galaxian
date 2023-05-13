package attack

import (
	"gamedev/config"

	"github.com/hajimehoshi/ebiten/v2"
)

type bullet interface {
	Draw(screen *ebiten.Image, config config.Config)
}
