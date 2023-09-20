package main

import (
	"fmt"
	ge "game/gameEngine"

	rl "github.com/gen2brain/raylib-go/raylib"
	
	//rlplus "github.com/lachee/raylib-goplus/raylib"
)


func main() {
	var g ge.GameEngine
	g.InitGameEngine(1920, 1080, "Knight Fight")
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
		rl.DrawTexture(startButton, 1200, y/8 + 150, rl.RayWhite)
		rl.DrawTexture(settingsButton, 1200, y/4 + 150, rl.RayWhite)
		rl.DrawTexture(quitButton, 1200, y/2 + 15, rl.RayWhite)
		


		rl.EndDrawing()
	}
	
	rl.CloseWindow()
}
