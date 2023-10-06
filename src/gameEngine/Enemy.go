package gameEngine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

//----- Enemy structure -----//

type Enemy struct {
	name               string
	level              int
	MaxHealthPoint     int
	currentHealthPoint int

	Sprite        rl.Texture2D
	Sr_sprite     rl.Rectangle
	Dr_sprite     rl.Rectangle
	Vector_sprite rl.Vector2

	Frame_count_sprite int

	Enemy_attack1         bool
	Sprite_attack1        rl.Texture2D
	Sr_Sprite_attack1     rl.Rectangle
	Dr_Sprite_attack1     rl.Rectangle
	Vector_Sprite_attack1 rl.Vector2
	timer_attack          int
}

//---- Enemy Init -----//

func (e *Enemy) Init() {
	e.name = "Zeljko"
	e.level = 5
	e.MaxHealthPoint = 300
	e.currentHealthPoint = 300

	e.Sprite = rl.LoadTexture("assets/Tilesets/enemy.png")
	e.Sr_sprite = rl.NewRectangle(0, 0, 720, 720)
	e.Dr_sprite = rl.NewRectangle(1000, 350, 512, 512)
	e.Vector_sprite = rl.NewVector2(0, 0)

	e.Frame_count_sprite = 0

	e.Enemy_attack1 = false
	e.Sprite_attack1 = rl.LoadTexture("assets/Tilesets/eclair_rouge.png")
	e.Sr_Sprite_attack1 = rl.NewRectangle(0, 0, 112.6, 298)
	e.Dr_Sprite_attack1 = rl.NewRectangle(665, 400, 112.6, 298)
	e.Vector_Sprite_attack1 = rl.NewVector2(0, 0)
	e.timer_attack = 0
}