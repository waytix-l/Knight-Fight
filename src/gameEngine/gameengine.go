package gameEngine

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"

)

type GameEngine struct {
	ScreenWidth int32
	ScreenHeight int32
	Running bool
	Sprite SpriteStruct
}

func (g *GameEngine) PrintScreenSize() {
	fmt.Println(g.ScreenWidth, "*", g.ScreenHeight)
}

func (g *GameEngine) InitGameEngine(x, y int32) {
	g.ScreenWidth = x
	g.ScreenHeight = y
}




//--------------------------------------------


type ClassPerso int

const (
	Archer ClassPerso = iota
	Warrior ClassPerso = iota
)

type Personnage struct {
	name               string
	class              ClassPerso
	level              int
	maxHealthPoint     int
	currentHealthPoint int
	inventory          map[string]int
}


func (p *Personnage) Init(Name string, Class ClassPerso, Level int, MaxHealthPoint int, CurrentHealthPoint int, Inventory map[string]int) {
	p.name = Name
	p.class = Class
	p.level = Level
	p.maxHealthPoint = MaxHealthPoint
	p.currentHealthPoint = CurrentHealthPoint
	p.inventory = Inventory
}

type SpriteStruct struct {
	Bouton_start rl.Texture2D
	Bouton_quit rl.Texture2D
	Bouton_setting rl.Texture2D
	Grass rl.Texture2D
}


