package object

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
	constants "github.com/zivlakmilos/tetris-go/private/contants"
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
			x := int32(col*o.cellSize + 1)
			y := int32(row*o.cellSize + 1)
			w := int32(o.cellSize - 1)
			h := int32(o.cellSize - 1)
			color := constants.Colors[cellValue]
			rl.DrawRectangle(x, y, w, h, color)
		}
	}
}
