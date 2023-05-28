package keypress

import (
	"gamedev/config"
	"gamedev/object/spaceship"
	"time"
)

type KeyPress struct {
	Timer time.Timer
}

func (keyPress *KeyPress) Update(ship spaceship.SpaceShip, config *config.Config) (bool, error) {
	move(ship, config)
	shot := isShoot(config)
	err := quit()
	return shot, err
}
