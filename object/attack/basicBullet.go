package attack

import (
	"gamedev/config"
	"gamedev/object/spaceship"
	"gamedev/util"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const basicBulletPixel string = "./object/bullet.png"

type BasicBarrage struct {
	Barrage []bullet
}

func (basicBarrage *BasicBarrage) Update(shot bool, config *config.Config, ship spaceship.SpaceShip) {
	if shot {
		basicBarrage.Barrage = append(basicBarrage.Barrage, *basicBarrage.newBullet(ship))
	}

	for i := 0; i < len(basicBarrage.Barrage); i++ {
		if !util.IsBulletExist(float64(config.ScreenHeight), basicBarrage.Barrage[i].Y) {
			tmp := basicBarrage.Barrage[i:]
			basicBarrage.Barrage = basicBarrage.Barrage[:i-1]
			basicBarrage.Barrage = append(basicBarrage.Barrage, tmp...)
		} else {
			basicBarrage.Barrage[i].Y -= config.BulletSpeedFactor
		}
	}
}

func (basicBarrage *BasicBarrage) Draw(screen *ebiten.Image, config *config.Config) {
	for _, barrage := range basicBarrage.Barrage {
		barrage.draw(screen, config)
	}
}

func (BasicBarrage *BasicBarrage) newBullet(ship spaceship.SpaceShip) *bullet {
	img, _, err := ebitenutil.NewImageFromFile(basicBulletPixel)
	if err != nil {
		log.Fatal(err)
	}
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()
	b := &bullet{
		image:  img,
		Width:  width,
		Height: height,
		X:      util.MediumPosition(ship.GetAxisX(), ship.GetWidth(), width),
		Y:      ship.GetAxisY() - 15,
	}
	return b
}

func NewBasicBarrage() BasicBarrage {
	basicBarrage := BasicBarrage{
		Barrage: make([]bullet, 0),
	}
	return basicBarrage
}
