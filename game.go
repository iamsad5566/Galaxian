package main

import (
	"gamedev/config"
	"gamedev/keypress"
	"gamedev/object"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	keyPress *keypress.KeyPress
	config   *config.Config
	ship     *object.Ship
	bullets  *object.Bullets
}

func (game *Game) Update() error {
	event, err := game.keyPress.Update(game.ship, game.config)
	game.bullets.Update(event, game.config, game.ship)
	return err
}

func (game *Game) Draw(screen *ebiten.Image) {
	screen.Fill(game.config.BgColor)
	game.ship.Draw(screen, game.config)
	game.bullets.Draw(screen, game.config)
}

func (game *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return game.config.ScreenWidth, game.config.ScreenHeight
}

func NewGame() *Game {
	cfg := config.LoadConfig()
	ebiten.SetWindowSize(cfg.ScreenWidth, cfg.ScreenHeight)
	ebiten.SetWindowTitle(cfg.Title)
	ebiten.SetTPS(cfg.FPS)

	return &Game{
		keyPress: &keypress.KeyPress{},
		config:   cfg,
		ship:     object.NewShip(cfg.ScreenWidth, cfg.ScreenHeight),
		bullets: &object.Bullets{
			BulletList: make([]*object.Bullet, 0),
		},
	}
}
