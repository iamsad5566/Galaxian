package attack

import (
	"gamedev/config"

	"github.com/hajimehoshi/ebiten/v2"
)

type bullet struct {
	image  *ebiten.Image
	Width  int
	Height int
	X      float64
	Y      float64
}

func (bullet *bullet) draw(screen *ebiten.Image, config *config.Config) {
	option := &ebiten.DrawImageOptions{}
	option.GeoM.Translate(bullet.X, bullet.Y)
	screen.DrawImage(bullet.image, option)
}
