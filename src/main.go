package main

import (
	ge "game/gameEngine"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func drawScene() {
	var g ge.GameEngine
	rl.DrawRectangle(200, 200, 200, 50, rl.Black)
	rl.DrawRectangle(200, 200, 195, 45, rl.Red)
	rl.DrawText("Quitter", 240, 210, 30, rl.Black)
	rl.DrawTexture(g.Sprite.Bouton_start, 0, 0, rl.White)
	rl.DrawTexture(g.Sprite.Grass, 100, 100, rl.Green)
}

func input() {
	if rl.IsKeyDown(rl.KeyEscape) {
		quit()
	}
	x := rl.GetMouseX()
	y := rl.GetMouseY()
	if x > 200 && x < 400 && y > 200 && y < 250 {
		if rl.IsMouseButtonPressed(0) {
			quit()
		}
	}
}

func update() {
	var g ge.GameEngine
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

func init() {
	var g ge.GameEngine
	g.InitGameEngine(1920, 1080)
	rl.InitWindow(g.ScreenWidth, g.ScreenHeight, "Knight Fight")
	g.Sprite.Bouton_start = rl.LoadTexture("assets/Tilesets/bouton_start.png")
	g.Sprite.Grass = rl.LoadTexture("assets/Tilesets/Grass.png")

	rl.SetTargetFPS(60)
}

func main() {
	var g ge.GameEngine
	g.Running = true
	for g.Running {
		input()
		update()
		render()
	}
	quit()
}
