package main

import (
	ge "game/gameEngine"
)

//----- Main Function -----//

func main() {
	var g ge.GameEngine
	var m ge.Menu
	g.InitGameEngine(1920, 1080, "Knight Fight")
	g.RunningGameEngine(&m)
}
