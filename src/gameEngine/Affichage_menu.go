package gameEngine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

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

	rl.DrawTexture(m.QuitButton, m.Bouton_X, m.Bouton_Y+125, rl.RayWhite)

	x_mouse := rl.GetMouseX()
	y_mouse := rl.GetMouseY()
	if x_mouse > m.Bouton_X+285 && x_mouse < m.Bouton_X+525 && y_mouse > m.Bouton_Y+125 && y_mouse < m.Bouton_Y+235 {
		rl.DrawTexture(m.StartButtonOver, m.Bouton_X, m.Bouton_Y, rl.RayWhite)
		if rl.IsMouseButtonPressed(0) {
			rl.EndDrawing()
			m.menu = 1
		}
	}

	if x_mouse > m.Bouton_X+270 && x_mouse < m.Bouton_X+550 && y_mouse > m.Bouton_Y+125+125 && y_mouse < m.Bouton_Y+235+125 {
		rl.DrawTexture(m.QuitButtonOver, m.Bouton_X, m.Bouton_Y+125, rl.RayWhite)
		if rl.IsMouseButtonPressed(0) {
			rl.CloseWindow()
		}
	}

	rl.EndDrawing()
}