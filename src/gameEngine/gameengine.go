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
	Sprite       SpriteStruct
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
	perso.Init("Lukas", Archer, 1, 100, 100, inventaire)
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

		case 4:
			m.Afficher_Menu_Jeu_porte2(&perso)

		}

	}

	rl.CloseWindow()

}

//--------------------------------------------




//----- Affichage Menu Principal -----//



//----- Affichage Lore -----//



//----- Affichage Jeu Scene Montagne -----//

func (m *Menu) Afficher_Menu_Jeu(perso *Personnage) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.White)

	rl.DrawTexturePro(
		m.Montagne_Background,
		m.Sr_Montagne,
		m.Dr_Montagne,
		m.Vector_montagne,
		0,
		rl.White,
	)

	rl.DrawTexturePro(
		m.Sol,
		m.Sr_sol,
		m.Dr_sol,
		m.Vector_sol,
		0,
		rl.White,
	)

	rl.DrawTexturePro(
		m.Ath,
		m.Sr_Ath,
		m.Dr_Ath,
		m.Vector_Ath,
		0,
		rl.White,
	)

	if (rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft)) && perso.Dr_sprite.X > -40 {
		perso.FrameCount++
		perso.Dr_sprite.X -= perso.Sprite_Speed
		if perso.FrameCount == 3 {
			perso.Sr_sprite.X += 128
			perso.FrameCount = 0
		}
		perso.sprite = rl.LoadTexture("assets/Tilesets/runleft.png")
		if rl.IsKeyDown(rl.KeySpace) && perso.Dr_sprite.Y >= float32(m.Ground_Pos) {
			perso.jump = true
		}

	} else if (rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight)) && perso.Dr_sprite.X < 1825 {
		perso.FrameCount++
		perso.Dr_sprite.X += perso.Sprite_Speed
		if perso.FrameCount == 3 {
			perso.Sr_sprite.X += 128
			perso.FrameCount = 0
		}
		perso.sprite = rl.LoadTexture("assets/Tilesets/Run.png")
		if rl.IsKeyDown(rl.KeySpace) && perso.Dr_sprite.Y >= float32(m.Ground_Pos) {
			perso.jump = true
		}

	} else {
		perso.sprite = rl.LoadTexture("assets/Tilesets/Idle.png")
	}

	if rl.IsKeyDown(rl.KeySpace) && perso.Dr_sprite.Y >= float32(m.Ground_Pos) {
		perso.jump = true
	}

	if perso.jump {
		perso.jump_timer++
		if perso.jump_timer < 20 {
			perso.Dr_sprite.Y -= perso.Sprite_Speed * 1.3
		} else if perso.jump_timer < 25 {
			perso.Dr_sprite.Y -= perso.Sprite_Speed * 0.8
		} else if perso.jump_timer < 30 {

		} else if perso.Dr_sprite.Y != float32(m.Ground_Pos) {
			perso.Dr_sprite.Y += perso.Sprite_Speed * 1.5
		}
	}

	if perso.Dr_sprite.Y >= float32(m.Ground_Pos) {
		perso.jump = false
		perso.jump_timer = 0

		perso.Dr_sprite.Y = float32(m.Ground_Pos)
	}

	rl.DrawTexturePro(
		perso.sprite,
		perso.Sr_sprite,
		perso.Dr_sprite,
		perso.Vector_sprite,
		0,
		rl.RayWhite,
	)

	currentHealthPoint := fmt.Sprint(perso.currentHealthPoint)
	maxHealthPoint := fmt.Sprint(perso.maxHealthPoint)
	level := fmt.Sprint(perso.level)
	class := ""
	if perso.class == 0 {
		class = "Archer"
	} else if perso.class == 1 {
		class = "Warrior"
	}

	couleur_vie := rl.Green
	if perso.currentHealthPoint <= perso.maxHealthPoint/10 {
		couleur_vie = rl.Red
	} else if perso.currentHealthPoint <= perso.maxHealthPoint/2 {
		couleur_vie = rl.Yellow
	}

	rl.DrawText(perso.name, int32(1785-10*len(perso.name)), 35, 30, rl.Black)
	rl.DrawText(currentHealthPoint, 1700, 80, 30, couleur_vie)
	rl.DrawText("/", 1760, 80, 30, couleur_vie)
	rl.DrawText(maxHealthPoint, 1800, 80, 30, couleur_vie)
	rl.DrawText("Lvl.", 1795, 120, 30, rl.Black)
	rl.DrawText(level, 1850, 120, 30, rl.Black)
	rl.DrawText("'G' : Take Potion", 20, 1050, 20, rl.RayWhite)
	rl.DrawText(class, 1670, 120, 30, rl.Black)

	if rl.IsKeyPressed(rl.KeyG) {
		if perso.currentHealthPoint < perso.maxHealthPoint {
			perso.currentHealthPoint -= 10
		}
		perso.level++
	}

	if perso.Dr_sprite.X > 440 && perso.Dr_sprite.X < 560 {
		rl.DrawText("APPUYEZ SUR 'E' POUR", 470, 840, 15, rl.Red)
		rl.DrawText("DÉCOUVRIR LE SYSTÈME DE COMBAT", 430, 870, 15, rl.Red)
		if rl.IsKeyDown(rl.KeyE) {
			rl.DrawText("SYSTÈME DE COMBAT", 730, 200, 40, rl.DarkGray)
		}
	}

	if perso.Dr_sprite.X > 1000 && perso.Dr_sprite.X < 1060 && perso.Donjon < 1 {
		rl.DrawText("APPUYEZ SUR 'C' POUR", 1010, 750, 15, rl.Red)
		rl.DrawText("ENTRER DANS LE DONJON", 1000, 780, 15, rl.Red)
		if rl.IsKeyPressed(rl.KeyC) {
			m.menu = 3
			perso.Dr_sprite.Y = 565
			perso.sprite = rl.LoadTexture("assets/Tilesets/Idle.png")
			perso.Dr_sprite.X = 600
			perso.Donjon = 1
		}
	}

	rl.EndDrawing()

	if rl.IsKeyPressed(rl.KeyEscape) {
		rl.CloseWindow()
	}

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
		//if rl.IsMouseButtonPressed(0){
			//perso.sprite = rl.LoadTexture(aaa)
			//perso.Sr_sprite.X = 0
			//perso.dodgegif = true
		//}
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

//----- Affichage scene 2eme donjon dévérouillé -----//

func (m *Menu) Afficher_Menu_Jeu_porte2(perso *Personnage) {

}

//----- Type, Structure et Méthodes -----//

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
	//dodgegif 	rl.Texture2D
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
	//p.dodgegif = rl.LoadTexture(asset/)
}

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

func (e *Enemy) Init() {
	e.name = "Zeljko"
	e.level = 20
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

type SpriteStruct struct {
	Bouton_start   rl.Texture2D
	Bouton_quit    rl.Texture2D
	Bouton_setting rl.Texture2D
	Grass          rl.Texture2D
}

type SourceRectangle struct {
	X      float32
	Y      float32
	Width  float32
	Height float32
}

type DestRectangle struct {
	X      float32
	Y      float32
	Width  float32
	Height float32
}

func (r *SourceRectangle) InitSourceRectangle(x float32, y float32, width float32, height float32) {
	r.X = x
	r.Y = y
	r.Width = width
	r.Height = height
}

func (r *DestRectangle) InitDestRectangle(x float32, y float32, width float32, height float32) {
	r.X = x
	r.Y = y
	r.Width = width
	r.Height = height
}
