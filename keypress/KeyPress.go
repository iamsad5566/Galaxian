package keypress

import (
	"gamedev/config"
	"gamedev/object/spaceship"
	"time"
)

type KeyPress struct {
	Timer time.Timer
}

func (keyPress *KeyPress) Update(ship *spaceship.SpaceShip, config *config.Config) (int, error) {
	err := move(ship, config)
}
