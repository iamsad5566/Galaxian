package object

import (
	"gamedev/config"
	"log"

	_ "golang.org/x/image/bmp"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Ship struct {
	image  *ebiten.Image
	Width  int
	Height int
	X      float64
	Y      float64
}

func (ship *Ship) Draw(screen *ebiten.Image, cfg *config.Config) {
	open := &ebiten.DrawImageOptions{}
	open.GeoM.Translate(ship.X, ship.Y)
	screen.DrawImage(ship.image, open)
}

func NewShip(screenWidth, screenHeight int) *Ship {
	img, _, err := ebitenutil.NewImageFromFile("./object/ship.bmp")
	if err != nil {
		log.Fatal(err)
	}

	width, height := img.Size()
	ship := &Ship{
		image:  img,
		Width:  width,
		Height: height,
		X:      float64(screenWidth-width) / 2,
		Y:      float64(screenHeight - width),
	}
	return ship
}
