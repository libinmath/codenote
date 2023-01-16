package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// https://mp.weixin.qq.com/s/5HfZ2TrnUl2pfBft5-CAJg

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("外星人入侵")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
