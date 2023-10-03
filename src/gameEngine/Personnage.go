package gameEngine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)


type ClassPerso int

const (
	Archer  ClassPerso = iota
	Warrior ClassPerso = iota
)

type Personnage struct {
	name               string
	class              ClassPerso
	level              int
	maxHealthPoint     int
	currentHealthPoint int
	inventory          map[string]int

	sprite         rl.Texture2D
	Sr_sprite      rl.Rectangle
	Dr_sprite      rl.Rectangle
	Vector_sprite  rl.Vector2
	FrameCount     int
	Sprite_Speed   float32
	Running        bool
	Donjon         int
	jump           bool
	jump_timer     int
	attack1        bool
	attack2        bool
	attack2_anim   rl.Texture2D
	Sr_attack2     rl.Rectangle
	Dr_attack2     rl.Rectangle
	Vector_attack2 rl.Vector2

	attack3      bool
	dodge        bool
	Dodge_effect 	rl.Texture2D
	timer_attack int
}

func (p *Personnage) Init(Name string, Class ClassPerso, Level int, MaxHealthPoint int, CurrentHealthPoint int, Inventory map[string]int) {
	p.name = Name
	p.class = Class
	p.level = Level
	p.maxHealthPoint = MaxHealthPoint
	p.currentHealthPoint = CurrentHealthPoint
	p.inventory = Inventory

	p.sprite = rl.LoadTexture("assets/Tilesets/Idle.png")
	p.Sr_sprite = rl.NewRectangle(0, 0, 128, 128)
	p.Dr_sprite = rl.NewRectangle(200, 840, 128, 128)
	p.Vector_sprite = rl.NewVector2(0, 0)
	p.FrameCount = 0
	p.Sprite_Speed = 3
	p.Running = false
	p.Donjon = 0
	p.jump = false
	p.jump_timer = 0
	p.attack1 = false
	p.attack2 = false

	p.attack2_anim = rl.LoadTexture("assets/Tilesets/spritesheet_animatedsword.png")
	p.Sr_attack2 = rl.NewRectangle(0, 0, 240, 196)
	p.Dr_attack2 = rl.NewRectangle(1100, 600, 240, 196)
	p.Vector_attack2 = rl.NewVector2(0, 0)

	p.attack3 = false
	p.dodge = false
	p.timer_attack = 0
	
	p.Dodge_effect = rl.LoadTexture("assets/Tilesets/Dodge_effect.png")
}