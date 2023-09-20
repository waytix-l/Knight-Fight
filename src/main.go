package main

import (
	"fmt"
	ge "game/gameEngine"

	rl "github.com/gen2brain/raylib-go/raylib"
	
	//rlplus "github.com/lachee/raylib-goplus/raylib"
)


func main() {
	var g ge.GameEngine
	g.InitGameEngine(100, 800, "Knight Fight")
	rl.ToggleFullscreen()
	x:= int32(rl.GetMonitorWidth(rl.GetCurrentMonitor()))
	y := int32(rl.GetMonitorHeight(rl.GetCurrentMonitor()))
	fmt.Print(x,y)
	startButton := rl.LoadTexture("assets/Tilesets/bouton_start2.png")
	quitButton := rl.LoadTexture("assets/Tilesets/bouton_quit2.png")
	settingsButton := rl.LoadTexture("assets/Tilesets/bouton_settings2.png")

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.White)
		rl.DrawText("aaaaaaaaaa", 10, 10, 20, rl.Black)
		rl.DrawTexture(startButton, x/2-x/5, y/8, rl.RayWhite)
		rl.DrawTexture(settingsButton, x/2-x/5, y/4, rl.RayWhite)
		rl.DrawTexture(quitButton, x/2-x/5, y/2 - 120, rl.RayWhite)
		


		rl.EndDrawing()
	}
	
	rl.CloseWindow()
}
