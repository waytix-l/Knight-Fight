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

func (g *GameEngine) RunningGameEngine() {
	rl.SetExitKey(0)
	var sr rl.Rectangle
	var dr rl.Rectangle

	sr = rl.NewRectangle(0, 0, 800, 500)
	dr = rl.NewRectangle(0, 0, 1920, 1080)
	vecteur := rl.NewVector2(0, 0)

	sourceMontagne := rl.NewRectangle(0, 0, 1600, 800)
	destMontagne := rl.NewRectangle(0, 0, 1920, 1080)

	sourceSol := rl.NewRectangle(0, 700, 1500, 1080)
	destSol := rl.NewRectangle(0, 400, 3500, 2000)

	x := int32(rl.GetMonitorWidth(rl.GetCurrentMonitor()))
	y := int32(rl.GetMonitorHeight(rl.GetCurrentMonitor()))
	fmt.Print(x, y)
	startButton := rl.LoadTexture("assets/Tilesets/bouton_start3.png")
	quitButton := rl.LoadTexture("assets/Tilesets/bouton_quit3.png")
	settingsButton := rl.LoadTexture("assets/Tilesets/bouton_settings3.png")
	startButtonOver := rl.LoadTexture("assets/Tilesets/bouton_start_gris2.png")
	quitButtonOver := rl.LoadTexture("assets/Tilesets/bouton_quit_gris2.png")
	settingsButtonOver := rl.LoadTexture("assets/Tilesets/bouton_settings_gris2.png")
	fond := rl.LoadTexture("assets/Tilesets/Fond_anime.png")
	montagne := rl.LoadTexture("assets/Tilesets/fond_montagne.png")
	sol := rl.LoadTexture("assets/Tilesets/mapv0.1.png")
	title := rl.LoadTexture("assets/Tilesets/knight_fight_title.png")
	bouton_x := 1200
	bouton_y := 400
	menu := 0
	color_black := rl.ColorAlpha(rl.Black, 0.5)
	color_gray := rl.ColorAlpha(rl.Gray, 0.5)
	frame_count := 0
	frame_count_sword := 0
	frame_count_eclair := 0

	player := rl.LoadTexture("assets/Tilesets/bouton_quit_gris2.png")
	playerSrc := rl.NewRectangle(0, 0, 800, 300)
	playerDest := rl.NewRectangle(200, 864, 200, 130)
	playerSpeed := float32(3)

	test_sword := rl.LoadTexture("assets/Tilesets/spritesheet_animatedsword.png")
	swordSrc := rl.NewRectangle(0, 0, 240, 196)
	swordDest := rl.NewRectangle(800, 400, 240, 196)

	test_eclair := rl.LoadTexture("assets/Tilesets/spritesheet_eclair_test.png")
	eclairSrc := rl.NewRectangle(0, 0, 800, 600)
	eclairDest := rl.NewRectangle(600, 400, 240, 196)

	for !rl.WindowShouldClose() {
		switch menu {
		case 0:
			frame_count++
			rl.BeginDrawing()
			rl.ClearBackground(rl.White)
			rl.DrawTexturePro(
				fond,
				sr,
				dr,
				vecteur,
				0,
				rl.RayWhite,
			)
			if sr.Width == 5600 && frame_count == 8 {
				sr.X = 0
				frame_count = 0
			} else if frame_count == 8 {
				sr.X += 800
				frame_count = 0
			}

			rl.DrawTexturePro(
				title,
				rl.NewRectangle(0, 0, 800, 500),
				rl.NewRectangle(470, -140, 1000, 650),
				rl.NewVector2(0, 0),
				0,
				rl.White,
			)
			rl.DrawTexture(startButton, int32(bouton_x), int32(bouton_y), rl.RayWhite)
			rl.DrawTexture(settingsButton, int32(bouton_x), int32(bouton_y)+125, rl.RayWhite)
			rl.DrawTexture(quitButton, int32(bouton_x), int32(bouton_y)+255, rl.RayWhite)

			x_mouse := rl.GetMouseX()
			y_mouse := rl.GetMouseY()
			if x_mouse > int32(bouton_x)+285 && x_mouse < int32(bouton_x)+525 && y_mouse > int32(bouton_y)+125 && y_mouse < int32(bouton_y)+235 {
				rl.DrawTexture(startButtonOver, int32(bouton_x), int32(bouton_y), rl.RayWhite)
				if rl.IsMouseButtonPressed(0) {
					fmt.Println("Start")
					menu = 1
				}
			}

			if x_mouse > int32(bouton_x)+285 && x_mouse < int32(bouton_x)+550 && y_mouse > int32(bouton_y)+125+130 && y_mouse < int32(bouton_y)+250+105 {
				rl.DrawTexture(settingsButtonOver, int32(bouton_x), int32(bouton_y)+125, rl.RayWhite)
				if rl.IsMouseButtonPressed(0) {
					fmt.Println("Settings")
					menu = 2
				}
			}

			if x_mouse > int32(bouton_x)+270 && x_mouse < int32(bouton_x)+550 && y_mouse > int32(bouton_y)+125+130+135 && y_mouse < int32(bouton_y)+250+105+135 {
				rl.DrawTexture(quitButtonOver, int32(bouton_x), int32(bouton_y)+255, rl.RayWhite)
				if rl.IsMouseButtonPressed(0) {
					rl.CloseWindow()
				}
			}

			rl.EndDrawing()

		case 1:
			frame_count_sword++
			frame_count_eclair++
			if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
				playerDest.Y -= playerSpeed
				time.Sleep(time.Millisecond * 20)
				playerDest.Y += playerSpeed
			}
			if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
				playerDest.X -= playerSpeed
			}
			if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
				playerDest.X += playerSpeed
			}

			rl.BeginDrawing()
			rl.ClearBackground(rl.White)
			rl.DrawTexturePro(
				montagne,
				sourceMontagne,
				destMontagne,
				vecteur,
				0,
				rl.White,
			)
			rl.DrawTexturePro(
				sol,
				sourceSol,
				destSol,
				vecteur,
				0,
				rl.White,
			)

			rl.DrawTexturePro(
				player,
				playerSrc,
				playerDest,
				vecteur,
				0,
				rl.White,
			)

			rl.DrawTexturePro(
				test_sword,
				swordSrc,
				swordDest,
				vecteur,
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

			rl.DrawTexturePro(
				test_eclair,
				eclairSrc,
				eclairDest,
				vecteur,
				0,
				rl.White,
			)

			//if eclairSrc.X == 23200 && frame_count_eclair == 100 {
			//	eclairSrc.X = 0
			//	frame_count_eclair = 0
			//} else if frame_count_eclair == 100 {
			//	eclairSrc.X += 800
			//	frame_count_eclair = 0
			//}

			rl.EndDrawing()

		case 2:
			rl.BeginDrawing()
			rl.ClearBackground(rl.White)

			rl.DrawTexturePro(
				fond,
				sr,
				dr,
				vecteur,
				0,
				rl.RayWhite,
			)

			rl.DrawRectangle(50, 50, 1830, 970, color_black)
			rl.DrawRectangle(60, 60, 1810, 950, color_gray)

			rl.DrawTexture(quitButton, int32(bouton_x)+100, int32(bouton_y)+330, rl.RayWhite)
			x_mouse := rl.GetMouseX()
			y_mouse := rl.GetMouseY()
			if x_mouse > int32(bouton_x)+385 && x_mouse < int32(bouton_x)+630 && y_mouse > int32(bouton_y)+450 && y_mouse < int32(bouton_y)+570 {
				rl.DrawTexture(quitButtonOver, int32(bouton_x)+100, int32(bouton_y)+330, rl.RayWhite)
				if rl.IsMouseButtonPressed(0) {
					menu = 0
				}
			}

			rl.DrawText("Settings", 100, 100, 40, rl.Black)
			rl.DrawText("Settings", 150, 150, 40, rl.Black)

			rl.EndDrawing()

		}

	}

	rl.CloseWindow()

}

//func menu_principal() {

//}

//--------------------------------------------

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
