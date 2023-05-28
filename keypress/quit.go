package keypress

import "github.com/hajimehoshi/ebiten/v2"

func quit() error {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return ebiten.Termination
	}
	return nil
}
