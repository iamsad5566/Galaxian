package main

import (
	"gamedev/config"
	"gamedev/keypress"
	"gamedev/object/attack"
	"gamedev/object/spaceship"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	keyPress *keypress.KeyPress
	config   *config.Config
	ship     spaceship.SpaceShip
	bullets  *attack.BasicBarrage
}

// behaviors in each updating frame
func (game *Game) Update() error {
	event, err := game.keyPress.Update(game.ship, game.config)
	game.bullets.Update(event, game.config, game.ship)
	return err
}

// draw the objects in each frame
func (game *Game) Draw(screen *ebiten.Image) {
	screen.Fill(game.config.BgColor)
	game.ship.Draw(screen)
	game.bullets.Draw(screen, game.config)
}

// define the size of the windows
func (game *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return game.config.ScreenWidth, game.config.ScreenHeight
}

func NewGame() *Game {
	// load the config first and set the basic parameters from the loaded configuration
	cfg := config.LoadConfig()
	ebiten.SetWindowSize(cfg.ScreenWidth, cfg.ScreenHeight)
	ebiten.SetWindowTitle(cfg.Title)
	ebiten.SetTPS(cfg.FPS)

	return &Game{
		keyPress: &keypress.KeyPress{},
		config:   cfg,
		ship:     spaceship.NewBaseShip(cfg),
		bullets:  attack.NewBasicBarrage(),
	}
}
