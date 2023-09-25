package gameEngine

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"

	"time"
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

	sourceMontagne := rl.NewRectangle(0, 0, 1600, 800)
	destMontagne := rl.NewRectangle(0, 0, 1920, 1080)

	sourceSol := rl.NewRectangle(0, 700, 1500, 1080)
	destSol := rl.NewRectangle(0, 400, 3500, 2000)

	x := int32(rl.GetMonitorWidth(rl.GetCurrentMonitor()))
	y := int32(rl.GetMonitorHeight(rl.GetCurrentMonitor()))
	fmt.Print(x, y)
	montagne := rl.LoadTexture("assets/Tilesets/fond_montagne.png")
	sol := rl.LoadTexture("assets/Tilesets/mapv0.9.png")
	frame_count_sword := 0
	frame_count_eclair := 0

	player := rl.LoadTexture("assets/Tilesets/Run.png")
	playerSrc := rl.NewRectangle(0, 0, 128, 128)
	playerDest := rl.NewRectangle(200, 840, 128, 128)
	playerSpeed := float32(3)

	test_sword := rl.LoadTexture("assets/Tilesets/spritesheet_animatedsword.png")
	swordSrc := rl.NewRectangle(0, 0, 240, 196)
	swordDest := rl.NewRectangle(800, 400, 240, 196)

	//eclair_bleu := rl.LoadTexture("assets/Tilesets/eclair_bleu.png")
	//eclairSrc := rl.NewRectangle(0, 0, 450, 300)
	//eclairDest := rl.NewRectangle(100, 100, 400, 240)

	m.Init_Menu()

	for !rl.WindowShouldClose() {
		switch m.menu {
		case 0:
			m.Afficher_Menu_Principal()

		case 1:
			frame_count_sword++
			frame_count_eclair++

			rl.BeginDrawing()
			rl.ClearBackground(rl.White)
			rl.DrawTexturePro(
				montagne,
				sourceMontagne,
				destMontagne,
				rl.NewVector2(0,0),
				0,
				rl.White,
			)
			rl.DrawTexturePro(
				sol,
				sourceSol,
				destSol,
				rl.NewVector2(0,0),
				0,
				rl.White,
			)

			rl.DrawTexturePro(
				player,
				playerSrc,
				playerDest,
				rl.NewVector2(0,0),
				0,
				rl.White,
			)


			if rl.IsKeyPressed(rl.KeyW) || rl.IsKeyPressed(rl.KeyUp) {
				playerDest.Y -= playerSpeed * 20
				time.Sleep(time.Millisecond * 2)
				playerDest.Y += playerSpeed * 20
			}
			if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
				playerDest.X -= playerSpeed
			} 
			if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
				playerDest.X += playerSpeed
			} 

			if frame_count_eclair == 3 && playerSrc.X == 896 {
				playerSrc.X = 0
				frame_count_eclair = 0
			} else if frame_count_eclair == 3 {
				playerSrc.X += 128
				frame_count_eclair = 0
			}

			if frame_count_eclair == 3 && playerSrc.X == 896 {
				playerSrc.X = 0
				frame_count_eclair = 0
			} else if frame_count_eclair == 3 {
				playerSrc.X += 128
				frame_count_eclair = 0
			}



			rl.DrawTexturePro(
				test_sword,
				swordSrc,
				swordDest,
				rl.NewVector2(0,0),
				0,
				rl.RayWhite,
			)

			if swordSrc.X == 2880 && frame_count_sword == 3 {
				swordSrc.X = 0
				frame_count_sword = 0
			} else if frame_count_sword == 3 {
				swordSrc.X += 240
				frame_count_sword = 0
			}

			if rl.IsKeyPressed(rl.KeyEscape) {
				rl.CloseWindow()
			}

			rl.EndDrawing()


		}

	}

	rl.CloseWindow()

}



//--------------------------------------------


type Menu struct {
	
	menu int

	//----- Menu Principal -----//

	FrameCount int
	X_mouse int32
	Y_mouse int32

	Fond rl.Texture2D
	Sr_fond rl.Rectangle
	Dr_fond rl.Rectangle
	Vector_fond rl.Vector2

	Title rl.Texture2D
	Sr_title rl.Rectangle
	Dr_title rl.Rectangle
	Vector_title rl.Vector2

	Bouton_X int32
	Bouton_Y int32

	StartButton rl.Texture2D
	StartButtonOver rl.Texture2D

	SettingsButton rl.Texture2D
	SettingsButtonOver rl.Texture2D

	QuitButton rl.Texture2D
	QuitButtonOver rl.Texture2D

	//----- Menu Jeu -----///

}

func (m *Menu) Init_Menu() {

	m.menu = 0

	//----- Menu Principal -----//

	m.Fond = rl.LoadTexture("assets/Tilesets/Fond_anime.png")
	m.Sr_fond = rl.NewRectangle(0, 0, 800, 500)
	m.Dr_fond = rl.NewRectangle(0, 0, 1920, 1080)
	m.Vector_fond = rl.NewVector2(0, 0)

	m.Title = rl.LoadTexture("assets/Tilesets/knight_fight_title.png")
	m.Sr_title = rl.NewRectangle(0, 0, 800, 500)
	m.Dr_title = rl.NewRectangle(470, -140, 1000, 650)
	m.Vector_title = rl.NewVector2(0, 0)
	
	m.StartButton = rl.LoadTexture("assets/Tilesets/bouton_start3.png")
	m.StartButtonOver = rl.LoadTexture("assets/Tilesets/bouton_start_gris2.png")

	m.SettingsButton = rl.LoadTexture("assets/Tilesets/bouton_settings3.png")
	m.SettingsButtonOver = rl.LoadTexture("assets/Tilesets/bouton_settings_gris2.png")

	m.QuitButton = rl.LoadTexture("assets/Tilesets/bouton_quit3.png")
	m.QuitButtonOver = rl.LoadTexture("assets/Tilesets/bouton_quit_gris2.png")

	m.Bouton_X = 1200
	m.Bouton_Y = 400

	m.FrameCount = 0

	//----- Menu Jeu -----//

}


func (m *Menu) Afficher_Menu_Principal() {
	
	m.FrameCount++

	rl.BeginDrawing()

	rl.ClearBackground(rl.White)
			rl.DrawTexturePro(
				m.Fond,
				m.Sr_fond,
				m.Dr_fond,
				m.Vector_fond,
				0,
				rl.RayWhite,
			)
			if m.Sr_fond.X == 5600 && m.FrameCount == 8 {
				m.Sr_fond.X = 0
				m.FrameCount = 0
			} else if m.FrameCount == 8 {
				m.Sr_fond.X += 800
				m.FrameCount = 0
			}

			rl.DrawTexturePro(
				m.Title,
				m.Sr_title,
				m.Dr_title,
				m.Vector_title,
				0,
				rl.White,
			)
			rl.DrawTexture(m.StartButton, m.Bouton_X, m.Bouton_Y, rl.RayWhite)
			rl.DrawTexture(m.SettingsButton, m.Bouton_X, m.Bouton_Y + 125, rl.RayWhite)
			rl.DrawTexture(m.QuitButton, m.Bouton_X, m.Bouton_Y+255, rl.RayWhite)

			x_mouse := rl.GetMouseX()
			y_mouse := rl.GetMouseY()
			if x_mouse > m.Bouton_X + 285 && x_mouse < m.Bouton_X + 525 && y_mouse > m.Bouton_Y + 125 && y_mouse < m.Bouton_Y + 235 {
				rl.DrawTexture(m.StartButtonOver, m.Bouton_X, m.Bouton_Y, rl.RayWhite)
				if rl.IsMouseButtonPressed(0) {
					fmt.Println("Start")
					m.menu = 1
				}
			}

			if x_mouse > m.Bouton_X + 285 && x_mouse < m.Bouton_X + 550 && y_mouse > m.Bouton_Y + 125 + 130 && y_mouse < m.Bouton_Y + 250 + 105 {
				rl.DrawTexture(m.SettingsButtonOver, m.Bouton_X, m.Bouton_Y + 125, rl.RayWhite)
				if rl.IsMouseButtonPressed(0) {
					fmt.Println("Settings")
					m.menu = 2
				}
			}

			if x_mouse > m.Bouton_X + 270 && x_mouse < m.Bouton_X + 550 && y_mouse > m.Bouton_Y + 125 + 130 + 135 && y_mouse < m.Bouton_Y + 250 + 105 + 135 {
				rl.DrawTexture(m.QuitButtonOver, m.Bouton_X, m.Bouton_Y + 255, rl.RayWhite)
				if rl.IsMouseButtonPressed(0) {
					rl.CloseWindow()
				}
			}



	rl.EndDrawing()
}

func (m *Menu) Afficher_Menu_Jeu() {
	
}

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
