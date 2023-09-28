package gameEngine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

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

	Lore_Display_Frame int

	//----- Menu Jeu -----//

	Montagne_Background rl.Texture2D
	Sr_Montagne         rl.Rectangle
	Dr_Montagne         rl.Rectangle
	Vector_montagne     rl.Vector2

	Sol        rl.Texture2D
	Sr_sol     rl.Rectangle
	Dr_sol     rl.Rectangle
	Vector_sol rl.Vector2

	Ground_Pos int
	Gravity    int

	Ath        rl.Texture2D
	Sr_Ath     rl.Rectangle
	Dr_Ath     rl.Rectangle
	Vector_Ath rl.Vector2

	//----- Menu Donjon -----//

	Fond_Donjon        rl.Texture2D
	Sr_fond_donjon     rl.Rectangle
	Dr_fond_donjon     rl.Rectangle
	Vector_fond_donjon rl.Vector2

	filtre_sombre rl.Texture2D

	Sol_donjon        rl.Texture2D
	Sr_sol_donjon     rl.Rectangle
	Dr_sol_donjon     rl.Rectangle
	Vector_sol_donjon rl.Vector2

	Attack1Button rl.Texture2D
	Attack1ButtonHover rl.Texture2D
	Attack2Button rl.Texture2D
	Attack2ButtonHover rl.Texture2D
	Attack3Button rl.Texture2D
	Attack3ButtonHover rl.Texture2D
	DodgeButton   rl.Texture2D
	DodgeButtonHover rl.Texture2D

	Fight bool
}