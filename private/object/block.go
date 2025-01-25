package object

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
	constants "github.com/zivlakmilos/tetris-go/private/contants"
)

type position struct {
	x int
	y int
}

type Block struct {
	id            int
	cellSize      int
	rotationState int
	colors        []color.RGBA
	cells         [][]position
}

func (o *Block) Setup() {
	o.cellSize = 30
	o.rotationState = 0
	o.colors = constants.Colors
}

func (o *Block) Update() {
}

func (o *Block) Render() {
	tiles := o.cells[o.rotationState]
	for _, tile := range tiles {
		x := int32(tile.x*o.cellSize + 1)
		y := int32(tile.y*o.cellSize + 1)
		w := int32(o.cellSize - 1)
		h := int32(o.cellSize - 1)
		col := o.colors[o.id]
		rl.DrawRectangle(x, y, w, h, col)
	}
}

func NewLBlock() *Block {
	return &Block{
		id: 1,
		cells: [][]position{
			{{2, 0}, {0, 1}, {1, 1}, {2, 1}},
			{{1, 0}, {1, 1}, {1, 2}, {2, 2}},
			{{0, 1}, {1, 1}, {2, 1}, {0, 2}},
			{{0, 0}, {1, 0}, {1, 1}, {1, 2}},
		},
	}
}

func NewJBlock() *Block {
	return &Block{
		id: 2,
		cells: [][]position{
			{{0, 0}, {0, 1}, {1, 1}, {2, 1}},
			{{1, 0}, {2, 0}, {1, 1}, {1, 2}},
			{{0, 1}, {1, 1}, {2, 1}, {2, 2}},
			{{1, 0}, {1, 1}, {0, 2}, {1, 2}},
		},
	}
}

func NewIBlock() *Block {
	return &Block{
		id: 3,
		cells: [][]position{
			{{0, 1}, {1, 1}, {2, 1}, {3, 1}},
			{{2, 0}, {2, 1}, {2, 2}, {2, 3}},
			{{0, 2}, {1, 2}, {2, 2}, {3, 2}},
			{{1, 0}, {1, 1}, {1, 2}, {1, 3}},
		},
	}
}

func NewOBlock() *Block {
	return &Block{
		id: 3,
		cells: [][]position{
			{{0, 0}, {1, 0}, {0, 1}, {1, 1}},
			{{0, 0}, {1, 0}, {0, 1}, {1, 1}},
			{{0, 0}, {1, 0}, {0, 1}, {1, 1}},
			{{0, 0}, {1, 0}, {0, 1}, {1, 1}},
		},
	}
}

func NewSBlock() *Block {
	return &Block{
		id: 3,
		cells: [][]position{
			{{1, 0}, {2, 0}, {0, 1}, {1, 1}},
			{{1, 0}, {1, 1}, {2, 1}, {2, 2}},
			{{1, 1}, {2, 1}, {0, 2}, {2, 1}},
			{{0, 0}, {0, 1}, {1, 1}, {1, 2}},
		},
	}
}

func NewTBlock() *Block {
	return &Block{
		id: 3,
		cells: [][]position{
			{{1, 0}, {0, 1}, {1, 1}, {2, 1}},
			{{0, 1}, {1, 1}, {2, 1}, {1, 2}},
			{{0, 1}, {1, 1}, {2, 1}, {1, 2}},
			{{1, 0}, {1, 0}, {1, 1}, {2, 1}},
		},
	}
}

func NewZBlock() *Block {
	return &Block{
		id: 3,
		cells: [][]position{
			{{0, 0}, {1, 0}, {1, 1}, {2, 1}},
			{{2, 0}, {1, 1}, {2, 1}, {1, 2}},
			{{0, 1}, {1, 1}, {1, 2}, {2, 2}},
			{{1, 0}, {0, 1}, {1, 1}, {0, 2}},
		},
	}
}
