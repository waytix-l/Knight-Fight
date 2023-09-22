package main

import (
	ge "game/gameEngine"

	//rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	var g ge.GameEngine
	g.InitGameEngine(1920, 1080, "Knight Fight")
	g.RunningGameEngine()
}
