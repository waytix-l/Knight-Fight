package main

import (
	"fmt"
	ge "game/gameEngine"

	rl "github.com/gen2brain/raylib-go/raylib"
)


func main() {
	var g ge.GameEngine
	g.InitGameEngine(100, 800, "Knight Fight")
	rl.ToggleFullscreen()
	x:= int32(rl.GetMonitorWidth(1))
	y := int32(rl.GetMonitorHeight(1))
	fmt.Print(x,y)
	quitButton := rl.LoadTexture("assets/Tilesets/bouton_quit2.png")
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.White)
		rl.DrawText("aaaaaaaaaa", 10, 10, 20, rl.Black)
		rl.DrawTexture(quitButton, x/2, y/3, rl.RayWhite)
		rl.EndDrawing()
	}
	rl.CloseWindow()
}
