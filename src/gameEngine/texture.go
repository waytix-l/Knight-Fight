package gameEngine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

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

	m.QuitButton = rl.LoadTexture("assets/Tilesets/bouton_quit3.png")
	m.QuitButtonOver = rl.LoadTexture("assets/Tilesets/bouton_quit_gris2.png")

	m.Bouton_X = 1200
	m.Bouton_Y = 600

	m.FrameCount = 0

	m.Lore_Display_Frame = 0

	//----- Menu Jeu -----//

	m.Montagne_Background = rl.LoadTexture("assets/Tilesets/fond_montagne.png")
	m.Sr_Montagne = rl.NewRectangle(0, 0, 1600, 800)
	m.Dr_Montagne = rl.NewRectangle(0, 0, 1920, 1080)
	m.Vector_montagne = rl.NewVector2(0, 0)

	m.Sol = rl.LoadTexture("assets/Tilesets/mapv1.0.png")
	m.Sr_sol = rl.NewRectangle(0, 700, 1500, 1080)
	m.Dr_sol = rl.NewRectangle(0, 400, 3500, 2000)
	m.Vector_sol = rl.NewVector2(0, 0)

	m.Ground_Pos = 840
	m.Gravity = 1

	m.Ath = rl.LoadTexture("assets/Tilesets/ath.png")
	m.Sr_Ath = rl.NewRectangle(0, 0, 500, 650)
	m.Dr_Ath = rl.NewRectangle(1360, -220, 1200, 1400)
	m.Vector_Ath = rl.NewVector2(0, 0)

	//----- Menu Donjon -----//

	m.Fond_Donjon = rl.LoadTexture("assets/Tilesets/dungeon1bg(1).gif")
	m.Sr_fond_donjon = rl.NewRectangle(0, 0, 1760, 1040)
	m.Dr_fond_donjon = rl.NewRectangle(0, 0, 1850, 995)
	m.Vector_fond_donjon = rl.NewVector2(0, 0)

	m.filtre_sombre = rl.LoadTexture("assets/Tilesets/filtre_sombre.png")

	m.Sol_donjon = rl.LoadTexture("assets/Tilesets/dungeon1.png")
	m.Sr_sol_donjon = rl.NewRectangle(0, 0, 970, 400)
	m.Dr_sol_donjon = rl.NewRectangle(0, 0, 1920, 1080)
	m.Vector_sol_donjon = rl.NewVector2(0, 0)

	m.Attack1Button = rl.LoadTexture("assets/Tilesets/buttonattack1.png")
	m.Attack2Button = rl.LoadTexture("assets/Tilesets/buttonattack2.png")
	m.Attack3Button = rl.LoadTexture("assets/Tilesets/buttonattack3.png")
	m.DodgeButton = rl.LoadTexture("assets/Tilesets/buttondodge.png")

	m.Attack1ButtonHover = rl.LoadTexture("assets/Tilesets/attack1.2.png")
	m.Attack2ButtonHover = rl.LoadTexture("assets/Tilesets/attack2.2.png")
	m.Attack3ButtonHover = rl.LoadTexture("assets/Tilesets/attack3.2.png")
	m.DodgeButtonHover = rl.LoadTexture("assets/Tilesets/dodge2.1.png")

	m.Fight = false
}