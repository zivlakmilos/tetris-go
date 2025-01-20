package game

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/zivlakmilos/tetris-go/private/object"
)

type Game struct {
	grid *object.Grid
}

func NewGame() *Game {
	return &Game{
		grid: object.NewGrid(),
	}
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

	g.grid.Setup()
}

func (g *Game) update() {
	g.grid.Update()
}

func (g *Game) render() {
	rl.BeginDrawing()
	rl.ClearBackground(color.RGBA{44, 44, 127, 255})

	g.grid.Render()

	rl.EndDrawing()
}
