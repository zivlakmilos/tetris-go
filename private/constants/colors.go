package constants

import "image/color"

var (
	DarKrey = color.RGBA{26, 31, 40, 255}
	Green   = color.RGBA{47, 230, 23, 255}
	Red     = color.RGBA{232, 18, 18, 255}
	Orange  = color.RGBA{226, 116, 17, 255}
	Yellow  = color.RGBA{237, 234, 4, 255}
	Purple  = color.RGBA{116, 0, 247, 255}
	Cyan    = color.RGBA{21, 204, 209, 255}
	Blue    = color.RGBA{13, 64, 216, 255}

	LightBlue = color.RGBA{59, 85, 162, 255}
	DarkBlue  = color.RGBA{44, 44, 127, 255}
)

var Colors = []color.RGBA{
	DarKrey,
	Green,
	Red,
	Orange,
	Yellow,
	Purple,
	Cyan,
	Blue,
}

const UpdateInterval = 0.2
