package keypress

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type KeyPress struct {
	Msg string
}

func (keyPress *KeyPress) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		fmt.Println("<----")
		keyPress.Msg = "left pressed"
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		fmt.Println("---->")
		keyPress.Msg = "right pressed"
	} else if ebiten.IsKeyPressed(ebiten.KeySpace) {
		fmt.Println("=====")
		keyPress.Msg = "space pressed"
	}

}
