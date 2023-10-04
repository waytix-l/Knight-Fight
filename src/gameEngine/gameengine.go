package gameEngine

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameEngine struct {
	ScreenWidth  int32
	ScreenHeight int32
	Title        string
	Running      bool
}

func (g *GameEngine) PrintScreenSize() {
	fmt.Println(g.ScreenWidth, "*", g.ScreenHeight)
}

func (g *GameEngine) InitGameEngine(x int32, y int32, title string) {
	g.ScreenWidth = x
	g.ScreenHeight = y
	g.Title = title
	rl.InitWindow(g.ScreenWidth, g.ScreenHeight, g.Title)
	rl.SetTargetFPS(60)
	rl.ToggleFullscreen()
}

func (g *GameEngine) RunningGameEngine(m *Menu) {
	rl.SetExitKey(0)

	m.Init_Menu()

	var perso Personnage
	inventaire := make(map[string]int)
	perso.Init("Lukas", Archer, 1, 100, 60, inventaire)
	var enemy Enemy
	enemy.Init()

	for !rl.WindowShouldClose() {
		switch m.menu {
		case 0:
			m.Afficher_Menu_Principal()

		case 1:
			m.Lore_Display()

		case 2:
			m.Afficher_Menu_Jeu(&perso)

		case 3:
			m.Afficher_Donjon(&perso, &enemy)

		}

	}

	rl.CloseWindow()

}

//----- Affichage scene de combat donjon -----//

func (m *Menu) Afficher_Donjon(perso *Personnage, enemy *Enemy) {

	rl.BeginDrawing()
	rl.ClearBackground(rl.White)

	rl.DrawTexturePro(
		m.Fond_Donjon,
		m.Sr_fond_donjon,
		m.Dr_fond_donjon,
		m.Vector_fond_donjon,
		0,
		rl.RayWhite,
	)

	rl.DrawTexture(m.filtre_sombre, 0, 0, rl.Black)

	rl.DrawTexturePro(
		m.Sol_donjon,
		m.Sr_sol_donjon,
		m.Dr_sol_donjon,
		m.Vector_sol_donjon,
		0,
		rl.RayWhite,
	)

	x_mouse := rl.GetMouseX()
	y_mouse := rl.GetMouseY()

	rl.DrawTexture(m.Attack1Button, 220, 900, rl.White)
	if x_mouse > 220 && x_mouse < 470 && y_mouse > 900 && y_mouse < 1030 {
		rl.DrawTexture(m.Attack1ButtonHover, 220, 900, rl.White)
		if rl.IsMouseButtonPressed(0) {
			perso.sprite = rl.LoadTexture("assets/Tilesets/Attack_1.png")
			perso.Sr_sprite.X = 0
			perso.attack1 = true
		}
	}
	rl.DrawTexture(m.Attack2Button, 620, 900, rl.White)
	if x_mouse > 620 && x_mouse < 870 && y_mouse > 900 && y_mouse < 1030 {
		rl.DrawTexture(m.Attack2ButtonHover, 620, 900, rl.White)
		if rl.IsMouseButtonPressed(0) {
			perso.attack2 = true
		}
	}
	rl.DrawTexture(m.Attack3Button, 1020, 900, rl.White)
	if x_mouse > 1020 && x_mouse < 1270 && y_mouse > 900 && y_mouse < 1030 {
		rl.DrawTexture(m.Attack3ButtonHover, 1020, 900, rl.White)
		if rl.IsMouseButtonPressed(0) {
			perso.sprite = rl.LoadTexture("assets/Tilesets/Attack_3.png")
			perso.Sr_sprite.X = 0
			perso.attack3 = true
		}
	}

	rl.DrawTexture(m.DodgeButton, 1420, 900, rl.White)
	if x_mouse > 1420 && x_mouse < 1670 && y_mouse > 900 && y_mouse < 1030 {
		rl.DrawTexture(m.DodgeButtonHover, 1420, 900, rl.White)
		if rl.IsMouseButtonPressed(0) {
			perso.dodge = true
		}
	}

	currentHealthPoint := fmt.Sprint(perso.currentHealthPoint)
	maxHealthPoint := fmt.Sprint(perso.maxHealthPoint)
	level := fmt.Sprint(perso.level)

	couleur_vie_perso := rl.Green
	if perso.currentHealthPoint <= perso.maxHealthPoint/10 {
		couleur_vie_perso = rl.Red
	} else if perso.currentHealthPoint <= perso.maxHealthPoint/2 {
		couleur_vie_perso = rl.Yellow
	}

	rl.DrawText(perso.name, int32(715-10*len(perso.name)), 600, 20, rl.White)
	rl.DrawText(currentHealthPoint, 690, 625, 20, couleur_vie_perso)
	rl.DrawText("/", 720, 625, 20, couleur_vie_perso)
	rl.DrawText(maxHealthPoint, 740, 625, 20, couleur_vie_perso)
	rl.DrawText("Lvl. ", int32(730-10*len(perso.name)+12*len(perso.name)), 600, 20, rl.White)
	rl.DrawText(level, int32(730-10*len(perso.name)+12*len(perso.name)+40), 600, 20, rl.White)
	//rl.DrawText("'G' : Take Potion", 20, 1050, 20, rl.RayWhite)

	currentHealthPointEnemy := fmt.Sprint(enemy.currentHealthPoint)
	maxHealthPointEnemy := fmt.Sprint(enemy.MaxHealthPoint)
	level_enemy := fmt.Sprint(enemy.level)

	couleur_vie_enemy := rl.Green
	if enemy.currentHealthPoint <= enemy.MaxHealthPoint/10 {
		couleur_vie_enemy = rl.Red
	} else if enemy.currentHealthPoint <= enemy.MaxHealthPoint/2 {
		couleur_vie_enemy = rl.Yellow
	}

	rl.DrawText(enemy.name, 1200, 380, 20, rl.White)
	rl.DrawText(currentHealthPointEnemy, 1225, 405, 20, couleur_vie_enemy)
	rl.DrawText("/", 1270, 405, 20, couleur_vie_enemy)
	rl.DrawText(maxHealthPointEnemy, 1295, 405, 20, couleur_vie_enemy)
	rl.DrawText("Lvl. ", 1280, 380, 20, rl.White)
	rl.DrawText(level_enemy, 1320, 380, 20, rl.White)

	if enemy.Frame_count_sprite == 6 && enemy.Sr_sprite.X == 9360 {
		enemy.Frame_count_sprite = 0
		enemy.Sr_sprite.X = 0
	} else if enemy.Frame_count_sprite == 6 {
		enemy.Frame_count_sprite = 0
		enemy.Sr_sprite.X += 720
	}

	rl.DrawTexturePro(
		enemy.Sprite,
		enemy.Sr_sprite,
		enemy.Dr_sprite,
		enemy.Vector_sprite,
		0,
		rl.RayWhite,
	)

	perso.Dr_sprite.Width = 256
	perso.Dr_sprite.Height = 256

	rl.DrawTexturePro(
		perso.sprite,
		perso.Sr_sprite,
		perso.Dr_sprite,
		perso.Vector_sprite,
		0,
		rl.RayWhite,
	)

	enemy.Frame_count_sprite++

	if rl.IsKeyPressed(rl.KeyH) {
		enemy.Enemy_attack1 = true
	}

	if perso.attack1 {
		perso.timer_attack++
		if perso.timer_attack < 20 {
			perso.Dr_sprite.X += perso.Sprite_Speed * 7.2
		} else if perso.timer_attack <= 60 && perso.timer_attack >= 20 {
			if perso.timer_attack%4 == 0 {
				perso.Sr_sprite.X += 128
			}
		} else if perso.timer_attack <= 80 && perso.timer_attack > 60 {
			perso.Dr_sprite.X -= 20.6
		} else {
			enemy.currentHealthPoint -= 50
			perso.timer_attack = 0
			perso.attack1 = false
			perso.sprite = rl.LoadTexture("assets/Tilesets/Idle.png")
			perso.Sr_sprite.X = 0
			enemy.Enemy_attack1 = true
		}
	}

	if perso.attack2 {
		perso.timer_attack++
		if perso.timer_attack < 20 {
			perso.Dr_sprite.X += perso.Sprite_Speed * 7
		} else if perso.timer_attack <= 56 && perso.timer_attack >= 20 {
			rl.DrawTexturePro(
				perso.attack2_anim,
				perso.Sr_attack2,
				perso.Dr_attack2,
				perso.Vector_attack2,
				0,
				rl.White,
			)
			if perso.timer_attack%3 == 0 {
				perso.Sr_attack2.X += 240
			}
		} else if perso.timer_attack <= 76 && perso.timer_attack > 56 {
			perso.Dr_sprite.X -= 20
		} else {
			enemy.currentHealthPoint -= 150
			perso.timer_attack = 0
			perso.attack2 = false
			perso.Sr_attack2.X = 0
			enemy.Enemy_attack1 = true
		}
	}

	if perso.attack3 {
		perso.timer_attack++
		if perso.timer_attack < 20 {
			perso.Dr_sprite.X += 20
		} else if perso.timer_attack <= 40 && perso.timer_attack >= 20 {
			perso.Dr_sprite.Y -= 8
		} else if perso.timer_attack <= 68 && perso.timer_attack > 40 {
			if perso.timer_attack%4 == 0 {
				perso.Sr_sprite.X += 128
			}
		} else if perso.timer_attack <= 88 && perso.timer_attack > 68 {
			perso.Dr_sprite.X -= 18.9
			perso.Dr_sprite.Y += 8.4
		} else {
			enemy.currentHealthPoint -= 100
			perso.timer_attack = 0
			perso.attack3 = false
			perso.sprite = rl.LoadTexture("assets/Tilesets/Idle.png")
			perso.Sr_sprite.X = 0
			enemy.Enemy_attack1 = true
		}

	}

	if perso.dodge {
		perso.timer_attack++
		if perso.timer_attack < 56 {
			rl.DrawTexturePro(
				perso.Dodge_effect,
				perso.sr_dodge,
				perso.dr_dodge,
				perso.Vector_dodge,
				0,
				rl.White,
			)
			if perso.timer_attack%4 == 0 {
				perso.sr_dodge.X += 248
			}

		} else {
			perso.dodge = false
			perso.timer_attack = 0
			perso.sr_dodge.X = 0
		}
	}

	if enemy.Enemy_attack1 {
		rl.DrawTexturePro(
			enemy.Sprite_attack1,
			enemy.Sr_Sprite_attack1,
			enemy.Dr_Sprite_attack1,
			enemy.Vector_Sprite_attack1,
			0,
			rl.White,
		)
		enemy.timer_attack++
		if enemy.timer_attack <= 20 {
			if enemy.timer_attack%4 == 0 {
				enemy.Sr_Sprite_attack1.X += 112.6
			}
		} else {
			perso.currentHealthPoint -= 20
			enemy.Sr_Sprite_attack1.X = 0
			enemy.Enemy_attack1 = false
			enemy.timer_attack = 0
		}
	}

	if enemy.currentHealthPoint <= 0 {
		perso.timer_attack++
		if perso.timer_attack > 40 {
			m.menu = 2
			perso.Dr_sprite.X = 950
			perso.Dr_sprite.Y = 840
			perso.Dr_sprite.Width = 128
			perso.Dr_sprite.Height = 128
			perso.timer_attack = 0
		}
	}

	if perso.currentHealthPoint <= 0 {

		rl.DrawTexturePro(
			perso.dead_screen,
			perso.sr_dead,
			perso.dr_dead,
			perso.vector_dead,
			0,
			rl.Gray,
		)
		if rl.IsKeyPressed(rl.KeyEnter) {
			m.menu = 2
		perso.Dr_sprite.X = 950
		perso.Dr_sprite.Y = 840
		perso.Dr_sprite.Width = 128
		perso.Dr_sprite.Height = 128
		perso.currentHealthPoint = perso.maxHealthPoint / 2
		perso.timer_attack = 0
		}
	}

	rl.EndDrawing()

	if rl.IsKeyPressed(rl.KeyEscape) {
		rl.CloseWindow()
	}

	if rl.IsKeyPressed(rl.KeyC) {
		m.menu = 2
		perso.Dr_sprite.X = 950
		perso.Dr_sprite.Y = 840
		perso.Dr_sprite.Width = 128
		perso.Dr_sprite.Height = 128
	}

}

//----- Fonction Combat -----//

func Fight(perso *Personnage, enemy *Enemy) {
	round := 0
	for perso.currentHealthPoint == 0 || enemy.currentHealthPoint == 0 {
		round++
	}
}
