package gameEngine

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"

)

type GameEngine struct {
	ScreenWidth int32
	ScreenHeight int32
	Title string
	Running bool
	Sprite SpriteStruct
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
	var sr rl.Rectangle
	var dr rl.Rectangle

	sr = rl.NewRectangle(0, 0, 800, 500)
	dr = rl.NewRectangle(0, 0, 1920, 1080)
	vecteur := rl.NewVector2(0, 0)

	
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
	title := rl.LoadTexture("assets/Tilesets/knight_fight_title.png")
	bouton_x := 1200
	bouton_y := 400
	menu := 0
	color_black := rl.ColorAlpha(rl.Black, 0.5)
	color_gray := rl.ColorAlpha(rl.Gray, 0.5)
	frame_count := 0

	for !rl.WindowShouldClose() {
		switch menu {
		case 0 :
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
				rl.NewVector2(0,0),
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
			rl.BeginDrawing()
			rl.ClearBackground(rl.White)

			rl.DrawText("Champi", 100, 100, 30, rl.Black)

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


			rl.DrawText("Settings", 70, 70, 30, rl.Black)

		
			rl.EndDrawing()
		
			}
		
			
		
	}

	rl.CloseWindow()

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


type SourceRectangle struct {
	X float32
	Y float32
	Width float32
	Height float32
}

type DestRectangle struct {
	X float32
	Y float32
	Width float32
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