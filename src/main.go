package main

import (
	ge "game/gameEngine"

	rl "github.com/gen2brain/raylib-go/raylib"
)


func main() {
	var g ge.GameEngine
	g.InitGameEngine(1920, 1080, "Knight Fight")
	caca := rl.LoadTexture("assets/Tilesets/bouton_start.png")
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.White)
		rl.DrawText("aaaaaaaaaa", 10, 10, 20, rl.Black)
		rl.DrawTexture(caca, 20, 20, rl.RayWhite)
		rl.EndDrawing()
	}
	rl.CloseWindow()
}
