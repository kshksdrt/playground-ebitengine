package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct{}

// Declare a variable with the type of upCount and assign it to 0
var upCount int = 0
var upKeyPressed bool = false

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		upKeyPressed = true
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyUp) {
		upKeyPressed = false
		upCount = 0
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if upKeyPressed {
		upCount = upCount + 1
	}
	ebitenutil.DebugPrint(screen, fmt.Sprint(upCount))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Rise")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
