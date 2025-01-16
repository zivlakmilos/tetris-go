package game

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct{}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Run() {
	g.setup()

	for !rl.WindowShouldClose() {
		g.update()
		g.render()
	}
}

func (g *Game) setup() {
	rl.InitWindow(300, 600, "Tetris")
	rl.SetTargetFPS(60)
}

func (g *Game) update() {
}

func (g *Game) render() {
	rl.BeginDrawing()
	rl.ClearBackground(color.RGBA{44, 44, 127, 255})
	rl.EndDrawing()
}
