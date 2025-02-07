package game

import (
	"fmt"
	"math/rand"
	"slices"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/zivlakmilos/tetris-go/private/assets"
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
	font           rl.Font
	score          int

	music           rl.Music
	rotateSound     rl.Sound
	rotateSoundWave rl.Wave
	clearSound      rl.Sound
	clearSoundWave  rl.Wave
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

	g.cleanup()
}

func (g *Game) setup() {
	rl.InitWindow(500, 620, "Tetris")
	rl.SetTargetFPS(60)

	rl.InitAudioDevice()
	g.music = rl.LoadMusicStreamFromMemory(".mp3", assets.SoundMusic, int32(len(assets.SoundMusic)))
	g.rotateSoundWave = rl.LoadWaveFromMemory(".mp3", assets.SoundRotate, int32(len(assets.SoundRotate)))
	g.rotateSound = rl.LoadSoundFromWave(g.rotateSoundWave)
	g.clearSoundWave = rl.LoadWaveFromMemory(".mp3", assets.SoundClear, int32(len(assets.SoundClear)))
	g.clearSound = rl.LoadSoundFromWave(g.clearSoundWave)

	rl.PlayMusicStream(g.music)

	g.font = rl.LoadFontFromMemory(".ttf", assets.MonogramFont, 64, nil)

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
	rl.UpdateMusicStream(g.music)

	rl.BeginDrawing()
	rl.ClearBackground(constants.DarkBlue)

	g.grid.Render()
	if !g.gameOver {
		g.currentBlock.Render(11, 11)
	}

	g.renderScore()
	g.renderNextBlock()
	g.renderGameOver()

	rl.EndDrawing()
}

func (g *Game) cleanup() {
	rl.UnloadMusicStream(g.music)
	rl.UnloadSound(g.rotateSound)
	rl.UnloadSound(g.clearSound)
	rl.UnloadWave(g.rotateSoundWave)
	rl.UnloadWave(g.clearSoundWave)
	rl.CloseAudioDevice()
}

func (g *Game) renderScore() {
	rl.DrawTextEx(g.font, "Score", rl.Vector2{365, 15}, 38, 2, rl.White)
	rl.DrawRectangleRounded(rl.Rectangle{320, 55, 170, 60}, 0.3, 6, constants.LightBlue)

	scoreText := fmt.Sprintf("%d", g.score)
	scoreTextSize := rl.MeasureTextEx(g.font, scoreText, 32, 2)
	rl.DrawTextEx(g.font, scoreText, rl.Vector2{320 + (170-scoreTextSize.X)/2, 65}, 38, 2, rl.White)
}

func (g *Game) renderNextBlock() {
	rl.DrawTextEx(g.font, "Next", rl.Vector2{370, 175}, 38, 2, rl.White)
	rl.DrawRectangleRounded(rl.Rectangle{320, 215, 170, 180}, 0.3, 6, constants.LightBlue)

	switch g.nextBlock.Id {
	case 3:
		g.nextBlock.Render(255, 290)
	case 4:
		g.nextBlock.Render(255, 280)
	default:
		g.nextBlock.Render(270, 270)
	}
}

func (g *Game) renderGameOver() {
	if g.gameOver {
		rl.DrawTextEx(g.font, "GAME OVER", rl.Vector2{320, 450}, 38, 2, rl.White)
	}
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
		g.updateScore(0, 1)
	case rl.KeyUp:
		g.currentBlock.Rotate()
		if !g.isValidBlockPos() {
			g.currentBlock.UndoRotate()
		} else {
			rl.PlaySound(g.rotateSound)
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
	rowsClered := g.grid.ClearFullRows()
	if rowsClered > 0 {
		rl.PlaySound(g.clearSound)
		g.updateScore(rowsClered, 0)
	}
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

	g.score = 0
}

func (g *Game) updateScore(linesCleared, moveDownPointes int) {
	g.score += linesCleared*100 + moveDownPointes
}
