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

	m.Init_Menu()

	var perso Personnage
	inventaire := make(map[string]int)
	perso.Init("Lucas", Archer, 1, 100, 50, inventaire)

	for !rl.WindowShouldClose() {
		switch m.menu {
		case 0:
			m.Afficher_Menu_Principal()

		case 1:
			m.Afficher_Menu_Jeu(&perso)
		}

	}

	rl.CloseWindow()

}

//--------------------------------------------

type Menu struct {
	menu int

	//----- Menu Principal -----//

	FrameCount int
	X_mouse    int32
	Y_mouse    int32

	Fond        rl.Texture2D
	Sr_fond     rl.Rectangle
	Dr_fond     rl.Rectangle
	Vector_fond rl.Vector2

	Title        rl.Texture2D
	Sr_title     rl.Rectangle
	Dr_title     rl.Rectangle
	Vector_title rl.Vector2

	Bouton_X int32
	Bouton_Y int32

	StartButton     rl.Texture2D
	StartButtonOver rl.Texture2D

	SettingsButton     rl.Texture2D
	SettingsButtonOver rl.Texture2D

	QuitButton     rl.Texture2D
	QuitButtonOver rl.Texture2D

	//----- Menu Jeu -----//

	Montagne_Background rl.Texture2D
	Sr_Montagne         rl.Rectangle
	Dr_Montagne         rl.Rectangle
	Vector_montagne     rl.Vector2

	Sol        rl.Texture2D
	Sr_sol     rl.Rectangle
	Dr_sol     rl.Rectangle
	Vector_sol rl.Vector2
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

	m.Montagne_Background = rl.LoadTexture("assets/Tilesets/fond_montagne.png")
	m.Sr_Montagne = rl.NewRectangle(0, 0, 1600, 800)
	m.Dr_Montagne = rl.NewRectangle(0, 0, 1920, 1080)
	m.Vector_montagne = rl.NewVector2(0, 0)

	m.Sol = rl.LoadTexture("assets/Tilesets/mapv0.4.png")
	m.Sr_sol = rl.NewRectangle(0, 700, 1500, 1080)
	m.Dr_sol = rl.NewRectangle(0, 400, 3500, 2000)
	m.Vector_sol = rl.NewVector2(0, 0)
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
	rl.DrawTexture(m.SettingsButton, m.Bouton_X, m.Bouton_Y+125, rl.RayWhite)
	rl.DrawTexture(m.QuitButton, m.Bouton_X, m.Bouton_Y+255, rl.RayWhite)

	x_mouse := rl.GetMouseX()
	y_mouse := rl.GetMouseY()
	if x_mouse > m.Bouton_X+285 && x_mouse < m.Bouton_X+525 && y_mouse > m.Bouton_Y+125 && y_mouse < m.Bouton_Y+235 {
		rl.DrawTexture(m.StartButtonOver, m.Bouton_X, m.Bouton_Y, rl.RayWhite)
		if rl.IsMouseButtonPressed(0) {
			fmt.Println("Start")
			m.menu = 1
		}
	}

	if x_mouse > m.Bouton_X+285 && x_mouse < m.Bouton_X+550 && y_mouse > m.Bouton_Y+125+130 && y_mouse < m.Bouton_Y+250+105 {
		rl.DrawTexture(m.SettingsButtonOver, m.Bouton_X, m.Bouton_Y+125, rl.RayWhite)
		if rl.IsMouseButtonPressed(0) {
			fmt.Println("Settings")
			m.menu = 2
		}
	}

	if x_mouse > m.Bouton_X+270 && x_mouse < m.Bouton_X+550 && y_mouse > m.Bouton_Y+125+130+135 && y_mouse < m.Bouton_Y+250+105+135 {
		rl.DrawTexture(m.QuitButtonOver, m.Bouton_X, m.Bouton_Y+255, rl.RayWhite)
		if rl.IsMouseButtonPressed(0) {
			rl.CloseWindow()
		}
	}

	rl.EndDrawing()
}

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


	if rl.IsKeyPressed(rl.KeyW) || rl.IsKeyPressed(rl.KeyUp) {
		perso.Dr_sprite.Y -= perso.Sprite_Speed * 20
		time.Sleep(time.Millisecond * 2)
		perso.Dr_sprite.Y += perso.Sprite_Speed * 20
	}
	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		perso.Dr_sprite.X -= perso.Sprite_Speed
		fmt.Println(perso.Dr_sprite.X)
	}
	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		perso.Dr_sprite.X += perso.Sprite_Speed
		fmt.Println(perso.Dr_sprite.X)
	}

	rl.DrawTexturePro(
		perso.sprite,
		perso.Sr_sprite,
		perso.Dr_sprite,
		perso.Vector_sprite,
		0,
		rl.RayWhite,
	)


	rl.EndDrawing()

	if rl.IsKeyPressed(rl.KeyEscape) {
		rl.CloseWindow()
	}

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

	sprite        rl.Texture2D
	Sr_sprite     rl.Rectangle
	Dr_sprite     rl.Rectangle
	Vector_sprite rl.Vector2
	Sprite_Speed  float32
}

func (p *Personnage) Init(Name string, Class ClassPerso, Level int, MaxHealthPoint int, CurrentHealthPoint int, Inventory map[string]int) {
	p.name = Name
	p.class = Class
	p.level = Level
	p.maxHealthPoint = MaxHealthPoint
	p.currentHealthPoint = CurrentHealthPoint
	p.inventory = Inventory

	p.sprite = rl.LoadTexture("assets/Tilesets/Run.png")
	p.Sr_sprite = rl.NewRectangle(0, 0, 128, 128)
	p.Dr_sprite = rl.NewRectangle(200, 840, 128, 128)
	p.Vector_sprite = rl.NewVector2(0, 0)
	p.Sprite_Speed = 3
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
