package spaceship

import (
	"github.com/hajimehoshi/ebiten/v2"
)

var prefix string = "./object/spaceship/"

type SpaceShip interface {
	Draw(screen *ebiten.Image)
	ChangeShip(appearence *ebiten.Image)
	GetAxisX() float64
	GetAxisY() float64
	GetWidth() int
	GetHeight() int
	SetAxisX(x float64)
	SetAxisY(y float64)
}
