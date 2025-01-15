package main

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(300, 600, "Tetris")
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(color.RGBA{44, 44, 127, 255})
		rl.EndDrawing()
	}
}
