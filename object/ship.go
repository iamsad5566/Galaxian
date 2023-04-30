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
	width  int
	height int
}

func (ship *Ship) Draw(screen *ebiten.Image, cfg *config.Config) {
	open := &ebiten.DrawImageOptions{}
	open.GeoM.Translate(float64(cfg.ScreenWidth-ship.width)/2, float64(cfg.ScreenHeight-ship.height))
	screen.DrawImage(ship.image, open)
}

func NewShip() *Ship {
	img, _, err := ebitenutil.NewImageFromFile("./object/ship.bmp")
	if err != nil {
		log.Fatal(err)
	}

	width, height := img.Size()
	ship := &Ship{
		image:  img,
		width:  width,
		height: height,
	}
	return ship
}
