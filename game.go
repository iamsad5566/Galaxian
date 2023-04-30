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
}

func (game *Game) Update() error {
	game.keyPress.Update()
	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {
	screen.Fill(game.config.BgColor)
	game.ship.Draw(screen, game.config)
}

func (game *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return game.config.ScreenWidth, game.config.ScreenHeight
}

func NewGame() *Game {
	cfg := config.LoadConfig()
	ebiten.SetWindowSize(cfg.ScreenWidth, cfg.ScreenHeight)
	ebiten.SetWindowTitle(cfg.Title)

	return &Game{
		keyPress: &keypress.KeyPress{Msg: "Hello, World!"},
		config:   cfg,
		ship:     object.NewShip(),
	}
}
