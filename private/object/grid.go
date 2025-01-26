package object

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/zivlakmilos/tetris-go/private/constants"
)

type Grid struct {
	numRows  int
	numCols  int
	cellSize int
	colors   []color.RGBA

	Grid [20][10]int
}

func NewGrid() *Grid {
	return &Grid{
		numRows:  20,
		numCols:  10,
		cellSize: 30,
	}
}

func (o *Grid) Setup() {
	for i := 0; i < o.numRows; i++ {
		for j := 0; j < o.numCols; j++ {
			o.Grid[i][j] = 0
		}
	}
}

func (o *Grid) Update() {
}

func (o *Grid) Render() {
	for row := 0; row < o.numRows; row++ {
		for col := 0; col < o.numCols; col++ {
			cellValue := o.Grid[row][col]
			x := int32(col*o.cellSize + 11)
			y := int32(row*o.cellSize + 11)
			w := int32(o.cellSize - 1)
			h := int32(o.cellSize - 1)
			color := constants.Colors[cellValue]
			rl.DrawRectangle(x, y, w, h, color)
		}
	}
}

func (o *Grid) IsValidCell(x, y int) bool {
	if x < 0 || x >= o.numCols {
		return false
	}
	if y < 0 || y >= o.numRows {
		return false
	}

	if o.Grid[y][x] != 0 {
		return false
	}

	return true
}

func (o *Grid) ClearFullRows() {
	completed := 0
	for row := o.numRows - 1; row >= 0; row-- {
		if o.isRowFull(row) {
			completed++
		} else if completed > 0 {
			o.moveRowDown(row, completed)
		}
	}
}

func (o *Grid) isRowFull(row int) bool {
	for i := 0; i < o.numCols; i++ {
		if o.Grid[row][i] == 0 {
			return false
		}
	}

	return true
}

func (o *Grid) clearRow(row int) {
	for i := 0; i < o.numCols; i++ {
		o.Grid[row][i] = 0
	}
}

func (o *Grid) moveRowDown(row, numRows int) {
	for i := 0; i < o.numCols; i++ {
		o.Grid[row+numRows][i] = o.Grid[row][i]
		o.Grid[row][i] = 0
	}
}
