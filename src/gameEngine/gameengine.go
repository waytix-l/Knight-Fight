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

	sr = rl.NewRectangle(0, 0, 700, 500)
	dr = rl.NewRectangle(0, 0, 1920, 1080)
	vecteur := rl.NewVector2(0, 0)

	
	x := int32(rl.GetMonitorWidth(rl.GetCurrentMonitor()))
	y := int32(rl.GetMonitorHeight(rl.GetCurrentMonitor()))
	fmt.Print(x, y)
	startButton := rl.LoadTexture("assets/Tilesets/bouton_start2.png")
	quitButton := rl.LoadTexture("assets/Tilesets/bouton_quit2.png")
	settingsButton := rl.LoadTexture("assets/Tilesets/bouton_settings2.png")
	startButtonOver := rl.LoadTexture("assets/Tilesets/bouton_start_gris.png")
	quitButtonOver := rl.LoadTexture("assets/Tilesets/bouton_quit_gris.png")
	settingsButtonOver := rl.LoadTexture("assets/Tilesets/bouton_settings_gris.png")
	fond := rl.LoadTexture("assets/Tilesets/Tittle_Screen_jeu.png")
	bouton_x := 1200
	bouton_y := 400
	menu := 0
	color_black := rl.ColorAlpha(rl.Black, 0.5)
	color_gray := rl.ColorAlpha(rl.Gray, 0.5)

	for !rl.WindowShouldClose() {
		switch menu {
		case 0 :
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

			rl.DrawTexture(startButton, int32(bouton_x), int32(bouton_y), rl.RayWhite)
			rl.DrawTexture(settingsButton, int32(bouton_x), int32(bouton_y)+150, rl.RayWhite)
			rl.DrawTexture(quitButton, int32(bouton_x), int32(bouton_y)+300, rl.RayWhite)

			x_mouse := rl.GetMouseX()
			y_mouse := rl.GetMouseY()
			if x_mouse > int32(bouton_x)+270 && x_mouse < int32(bouton_x)+550 && y_mouse > int32(bouton_y)+130 && y_mouse < int32(bouton_y)+250 {
				rl.DrawTexture(startButtonOver, int32(bouton_x), int32(bouton_y), rl.RayWhite)
				if rl.IsMouseButtonPressed(0) {
					fmt.Println("Start")
					menu = 1
				}
			}

			if x_mouse > int32(bouton_x)+270 && x_mouse < int32(bouton_x)+550 && y_mouse > int32(bouton_y)+130+150 && y_mouse < int32(bouton_y)+250+150 {
				rl.DrawTexture(settingsButtonOver, int32(bouton_x), int32(bouton_y)+150, rl.RayWhite)
				if rl.IsMouseButtonPressed(0) {
					fmt.Println("Settings")
					menu = 2
				}
			}

			if x_mouse > int32(bouton_x)+270 && x_mouse < int32(bouton_x)+550 && y_mouse > int32(bouton_y)+130+150+150 && y_mouse < int32(bouton_y)+250+150+150 {
				rl.DrawTexture(quitButtonOver, int32(bouton_x), int32(bouton_y)+150+150, rl.RayWhite)
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

type FrameDisposal int 
const (
	FrameDisposalNone FrameDisposal = iota
	FrameDisposalDontDispose
	FrameDisposalRestoreBackground
	FrameDisposalRestorePrevious
)

type GifImage struct {
	Texture rl.Texture2D
	Width int
	Height int
	Frames int
	Timing []int
	Disposal []FrameDisposal
}


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