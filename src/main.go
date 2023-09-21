package main

import (
	ge "game/gameEngine"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	var g ge.GameEngine
	g.InitGameEngine(1920, 1080, "Knight Fight")
	g.RunningGameEngine()
}

func game() {
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.White)
		rl.DrawText("Champignon", 100, 100, 30, rl.Black)

		rl.EndDrawing()
	}
}
