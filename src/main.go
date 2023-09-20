package main

import (
	ge "game/gameEngine"

	rl "github.com/gen2brain/raylib-go/raylib"
)


func main() {
	var g ge.GameEngine
	g.InitGameEngine(1920, 1080, "Knight Fight")
	rl.DrawRectangle(200, 200, 200, 50, rl.Black)
	rl.DrawRectangle(200, 200, 195, 45, rl.Red)
	rl.DrawText("Quitter", 240, 210, 30, rl.Black)
	rl.DrawTexture(g.Sprite.Bouton_start, 0, 0, rl.White)
	rl.DrawTexture(g.Sprite.Grass, 100, 100, rl.Green)
}
