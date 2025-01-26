package object

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/zivlakmilos/tetris-go/private/constants"
)

type Position struct {
	X int
	Y int
}

type Block struct {
	Id            int
	x             int
	y             int
	cellSize      int
	rotationState int
	colors        []color.RGBA
	cells         [][]Position
}

func (o *Block) Setup() {
	o.cellSize = 30
	o.rotationState = 0
	o.colors = constants.Colors
}

func (o *Block) Update() {
}

func (o *Block) Render() {
	tiles := o.GetCellPositions()
	for _, tile := range tiles {
		x := int32(tile.X*o.cellSize + 11)
		y := int32(tile.Y*o.cellSize + 11)
		w := int32(o.cellSize - 1)
		h := int32(o.cellSize - 1)
		col := o.colors[o.Id]
		rl.DrawRectangle(x, y, w, h, col)
	}
}

func (o *Block) Move(x, y int) {
	o.x += x
	o.y += y
}

func (o *Block) Rotate() {
	o.rotationState++
	o.rotationState %= len(o.cells)
}

func (o *Block) UndoRotate() {
	o.rotationState--
	if o.rotationState < 0 {
		o.rotationState = len(o.cells) - 1
	}
}

func (o *Block) GetCellPositions() []Position {
	var res []Position

	tiles := o.cells[o.rotationState]
	for _, tile := range tiles {
		x := o.x + tile.X
		y := o.y + tile.Y
		res = append(res, Position{x, y})
	}

	return res
}

func NewLBlock() *Block {
	return &Block{
		Id: 1,
		x:  3,
		y:  0,
		cells: [][]Position{
			{{2, 0}, {0, 1}, {1, 1}, {2, 1}},
			{{1, 0}, {1, 1}, {1, 2}, {2, 2}},
			{{0, 1}, {1, 1}, {2, 1}, {0, 2}},
			{{0, 0}, {1, 0}, {1, 1}, {1, 2}},
		},
	}
}

func NewJBlock() *Block {
	return &Block{
		Id: 2,
		x:  3,
		y:  0,
		cells: [][]Position{
			{{0, 0}, {0, 1}, {1, 1}, {2, 1}},
			{{1, 0}, {2, 0}, {1, 1}, {1, 2}},
			{{0, 1}, {1, 1}, {2, 1}, {2, 2}},
			{{1, 0}, {1, 1}, {0, 2}, {1, 2}},
		},
	}
}

func NewIBlock() *Block {
	return &Block{
		Id: 3,
		x:  3,
		y:  -1,
		cells: [][]Position{
			{{0, 1}, {1, 1}, {2, 1}, {3, 1}},
			{{2, 0}, {2, 1}, {2, 2}, {2, 3}},
			{{0, 2}, {1, 2}, {2, 2}, {3, 2}},
			{{1, 0}, {1, 1}, {1, 2}, {1, 3}},
		},
	}
}

func NewOBlock() *Block {
	return &Block{
		Id: 4,
		x:  4,
		y:  0,
		cells: [][]Position{
			{{0, 0}, {1, 0}, {0, 1}, {1, 1}},
		},
	}
}

func NewSBlock() *Block {
	return &Block{
		Id: 5,
		x:  3,
		y:  0,
		cells: [][]Position{
			{{1, 0}, {2, 0}, {0, 1}, {1, 1}},
			{{1, 0}, {1, 1}, {2, 1}, {2, 2}},
			{{1, 1}, {2, 1}, {0, 2}, {1, 2}},
			{{0, 0}, {0, 1}, {1, 1}, {1, 2}},
		},
	}
}

func NewTBlock() *Block {
	return &Block{
		Id: 6,
		x:  3,
		y:  0,
		cells: [][]Position{
			{{1, 0}, {0, 1}, {1, 1}, {2, 1}},
			{{1, 0}, {1, 1}, {2, 1}, {1, 2}},
			{{0, 1}, {1, 1}, {2, 1}, {1, 2}},
			{{1, 0}, {0, 1}, {1, 1}, {1, 2}},
		},
	}
}

func NewZBlock() *Block {
	return &Block{
		Id: 7,
		x:  3,
		y:  0,
		cells: [][]Position{
			{{0, 0}, {1, 0}, {1, 1}, {2, 1}},
			{{2, 0}, {1, 1}, {2, 1}, {1, 2}},
			{{0, 1}, {1, 1}, {1, 2}, {2, 2}},
			{{1, 0}, {0, 1}, {1, 1}, {0, 2}},
		},
	}
}
