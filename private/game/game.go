package game

import (
	"image/color"
	"math/rand"
	"slices"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/zivlakmilos/tetris-go/private/constants"
	"github.com/zivlakmilos/tetris-go/private/object"
)

type Game struct {
	grid           *object.Grid
	blocks         []*object.Block
	currentBlock   *object.Block
	nextBlock      *object.Block
	lastUpdateTime float64
	gameOver       bool
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

	g.currentBlock = g.getRandomBlock()
	g.nextBlock = g.getRandomBlock()
}

func (g *Game) update() {
	if g.gameOver {
		if rl.GetKeyPressed() != 0 {
			g.reset()
		}
		return
	}

	currentTime := rl.GetTime()

	g.grid.Update()

	g.handleInput()

	if currentTime-g.lastUpdateTime >= constants.UpdateInterval {
		g.lastUpdateTime = currentTime
		g.moveBlockDown()
	}
}

func (g *Game) render() {
	rl.BeginDrawing()
	rl.ClearBackground(color.RGBA{44, 44, 127, 255})

	g.grid.Render()
	if !g.gameOver {
		g.currentBlock.Render()
	}

	rl.EndDrawing()
}

func (g *Game) getRandomBlock() *object.Block {
	if len(g.blocks) == 0 {
		g.blocks = g.getAllBlocks()
	}

	randomIdx := rand.Int() % len(g.blocks)
	block := g.blocks[randomIdx]
	g.blocks = slices.Delete(g.blocks, randomIdx, randomIdx+1)

	return block
}

func (g *Game) getAllBlocks() []*object.Block {
	res := []*object.Block{
		object.NewIBlock(),
		object.NewJBlock(),
		object.NewLBlock(),
		object.NewOBlock(),
		object.NewSBlock(),
		object.NewTBlock(),
		object.NewZBlock(),
	}

	for idx := range res {
		res[idx].Setup()
	}

	return res
}

func (g *Game) handleInput() {
	switch rl.GetKeyPressed() {
	case rl.KeyLeft:
		g.currentBlock.Move(-1, 0)
		if !g.isValidBlockPos() {
			g.currentBlock.Move(1, 0)
		}
	case rl.KeyRight:
		g.currentBlock.Move(1, 0)
		if !g.isValidBlockPos() {
			g.currentBlock.Move(-1, 0)
		}
	case rl.KeyDown:
		g.moveBlockDown()
	case rl.KeyUp:
		g.currentBlock.Rotate()
		if !g.isValidBlockPos() {
			g.currentBlock.UndoRotate()
		}
	}
}

func (g *Game) moveBlockDown() {
	g.currentBlock.Move(0, 1)
	if !g.isValidBlockPos() {
		g.currentBlock.Move(0, -1)
		g.lockBlock()
	}
}

func (g *Game) lockBlock() {
	tiles := g.currentBlock.GetCellPositions()
	for _, tile := range tiles {
		g.grid.Grid[tile.Y][tile.X] = g.currentBlock.Id
	}

	g.currentBlock = g.nextBlock
	if !g.isValidBlockPos() {
		g.gameOver = true
	}

	g.nextBlock = g.getRandomBlock()
	g.grid.ClearFullRows()
}

func (g *Game) isValidBlockPos() bool {
	tiles := g.currentBlock.GetCellPositions()
	for _, tile := range tiles {
		if !g.grid.IsValidCell(tile.X, tile.Y) {
			return false
		}
	}

	return true
}

func (g *Game) reset() {
	g.gameOver = false

	g.grid.Setup()
	g.blocks = g.getAllBlocks()
	g.currentBlock = g.getRandomBlock()
	g.nextBlock = g.getRandomBlock()
}
