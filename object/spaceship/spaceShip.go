package spaceship

import (
	"github.com/hajimehoshi/ebiten/v2"
)

var prefix string = "./object/spaceship/"

type SpaceShip interface {
	Draw(screen *ebiten.Image)
	ChangeShip(appearence *ebiten.Image)
}
