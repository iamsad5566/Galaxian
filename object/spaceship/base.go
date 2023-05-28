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

func (base *Base) GetAxisX() float64 {
	return base.PositionX
}

func (base *Base) GetAxisY() float64 {
	return base.PositionY
}

func (base *Base) GetWidth() int {
	return base.Width
}

func (base *Base) GetHeight() int {
	return base.Height
}

func (base *Base) SetAxisX(x float64) {
	base.PositionX = x
}

func (base *Base) SetAxisY(y float64) {
	base.PositionY = y
}

func NewBaseShip(config *config.Config) *Base {
	img, _, err := ebitenutil.NewImageFromFile(prefix + "base.png")
	if err != nil {
		log.Fatal(err)
	}
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()
	return &Base{
		appearance: img,
		Width:      width,
		Height:     height,
		PositionX:  float64(config.ScreenWidth-width) / 2,
		PositionY:  float64(config.ScreenHeight - height),
	}
}
