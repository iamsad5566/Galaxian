package object

import (
	"gamedev/config"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Ship struct {
	image    *ebiten.Image
	Screen   *ebiten.Image
	LastShot time.Time
	Width    int
	Height   int
	X        float64
	Y        float64
}

func (ship *Ship) Draw(screen *ebiten.Image, cfg *config.Config) {
	ship.Screen = screen
	open := &ebiten.DrawImageOptions{}
	open.GeoM.Translate(ship.X, ship.Y)
	screen.DrawImage(ship.image, open)
}

func NewShip(screenWidth, screenHeight int) *Ship {
	img, _, err := ebitenutil.NewImageFromFile("./object/ship.png")
	if err != nil {
		log.Fatal(err)
	}

	width, height := img.Size()
	ship := &Ship{
		image:    img,
		Width:    width,
		Height:   height,
		X:        float64(screenWidth-width) / 2,
		Y:        float64(screenHeight - height),
		LastShot: time.Now().Add(-1 * time.Minute),
	}
	return ship
}
