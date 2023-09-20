package main

import (
	"main/gameEngine"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func drawScene() {
	rl.DrawText("aaa", 10, 10, 10, rl.Black)
}

func input() {
	if rl.IsKeyDown(rl.KeyEscape) {
		quit()
	}
}

func update() {
	var g gameEngine.GameEngine
	g.Running = !rl.WindowShouldClose()
}

func render() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.White)

	drawScene()

	rl.EndDrawing()
}

func quit() {
	rl.CloseWindow()
}

func main() {
	var g gameEngine.GameEngine
	g.InitGameEngine(1920, 1080)
	g.Running = true
	rl.InitWindow(g.ScreenWidth, g.ScreenHeight, "Knight Fighter")
	rl.SetTargetFPS(60)

	for g.Running {
		input()
		update()
		render()
	}
	quit()
}
