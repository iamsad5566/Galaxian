package object

import (
	"gamedev/config"
	"gamedev/object/spaceship"
	"gamedev/util"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Bullets struct {
	BulletList []*Bullet
}

type Bullet struct {
	image  *ebiten.Image
	Width  int
	Height int
	X      float64
	Y      float64
}

func (bullets *Bullets) Update(shot bool, config *config.Config, ship spaceship.SpaceShip) {
	if shot {
		bullets.BulletList = append(bullets.BulletList, NewBullet(ship))
	}

	for i := 0; i < len(bullets.BulletList); i++ {
		if !util.IsBulletExist(float64(config.ScreenHeight), bullets.BulletList[i].Y) {
			tmp := bullets.BulletList[i:]
			bullets.BulletList = bullets.BulletList[:i-1]
			bullets.BulletList = append(bullets.BulletList, tmp...)
		} else {
			bullets.BulletList[i].Y -= config.BulletSpeedFactor
		}
	}
}

func (bullets *Bullets) Draw(screen *ebiten.Image, config *config.Config) {
	for _, bullet := range bullets.BulletList {
		bullet.draw(screen, config)
	}
}

func (bullet *Bullet) draw(screen *ebiten.Image, config *config.Config) {
	open := &ebiten.DrawImageOptions{}
	open.GeoM.Translate(bullet.X, bullet.Y)
	screen.DrawImage(bullet.image, open)
}

func NewBullet(ship spaceship.SpaceShip) *Bullet {
	img, _, err := ebitenutil.NewImageFromFile("./object/bullet.png")
	if err != nil {
		log.Fatal(err)
	}
	width, height := img.Size()
	bullet := &Bullet{
		image:  img,
		Width:  width,
		Height: height,
		X:      ship.GetAxisX() + float64(ship.GetWidth()-width)/2,
		Y:      ship.GetAxisY() - 15,
	}
	return bullet
}
