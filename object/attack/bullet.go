package attack

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type bullet struct {
	image  *ebiten.Image
	Width  int
	Height int
	X      float64
	Y      float64
}
