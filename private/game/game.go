package game

import (
	"image/color"
	"math/rand"
	"slices"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/zivlakmilos/tetris-go/private/object"
)

type Game struct {
	grid         *object.Grid
	blocks       []*object.Block
	currentBlock *object.Block
	nextBlock    *object.Block
}

func NewGame() *Game {
	return &Game{
		grid: object.NewGrid(),
	}
}

func (g *Game) Run() {
	g.setup()

	block := object.NewTBlock()
	block.Setup()

	for !rl.WindowShouldClose() {
		g.update()
		g.render()
		block.Render()
	}
}

func (g *Game) setup() {
	rl.InitWindow(300, 600, "Tetris")
	rl.SetTargetFPS(60)

	g.grid.Setup()

	g.currentBlock = g.getRandomBlock()
	g.nextBlock = g.getRandomBlock()
}

func (g *Game) update() {
	g.grid.Update()
}

func (g *Game) render() {
	rl.BeginDrawing()
	rl.ClearBackground(color.RGBA{44, 44, 127, 255})

	g.grid.Render()
	g.currentBlock.Render()

	rl.EndDrawing()
}

func (g *Game) getRandomBlock() *object.Block {
	if len(g.blocks) == 0 {
		g.blocks = g.getAllBlocks()
		for idx := range g.blocks {
			g.blocks[idx].Setup()
		}
	}

	randomIdx := rand.Int() % len(g.blocks)
	block := g.blocks[randomIdx]
	g.blocks = slices.Delete(g.blocks, randomIdx, randomIdx+1)

	return block
}

func (g *Game) getAllBlocks() []*object.Block {
	return []*object.Block{
		object.NewIBlock(),
		object.NewJBlock(),
		object.NewLBlock(),
		object.NewOBlock(),
		object.NewSBlock(),
		object.NewTBlock(),
		object.NewZBlock(),
	}
}
