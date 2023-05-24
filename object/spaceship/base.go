package spaceship

import (
	"gamedev/config"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var LastShot time.Time = time.Now().Add(-3 * time.Minute)

type Base struct {
	appearance *ebiten.Image
	Width      int
	Height     int
	PositionX  float64
	PositionY  float64
}

var _ SpaceShip = (*Base)(nil)

func (base *Base) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(base.PositionX, base.PositionY)
	screen.DrawImage(base.appearance, op)
}

func (base *Base) ChangeShip(appearance *ebiten.Image) {
	base.appearance = appearance
}

func NewBaseShip(config config.Config) *Base {
	img, _, err := ebitenutil.NewImageFromFile(prefix + "base.png")
	if err != nil {
		log.Fatal(err)
	}
	width, height := img.Size()
	return &Base{
		appearance: img,
		Width:      width,
		Height:     height,
		PositionX:  float64(config.ScreenWidth-width) / 2,
		PositionY:  float64(config.ScreenHeight - height),
	}
}
