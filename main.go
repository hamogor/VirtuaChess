package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

type Game struct {
	Match Match
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func NewGame() *Game {
	g := &Game{}
	g.Match.Player1 = CreateGoh()
	g.Match.Player2 = CreateGoh()
	g.Start()
	return g
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

}

func main() {
	g := NewGame()
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowTitle("VirtuaChess")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
