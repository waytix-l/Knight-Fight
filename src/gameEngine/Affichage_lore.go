package gameEngine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (m *Menu) Lore_Display() {
	m.Lore_Display_Frame++
	rl.BeginDrawing()
	rl.ClearBackground(rl.DarkGray)
	if m.Lore_Display_Frame < 120 {
		rl.DrawText("Press Enter to skip", 50, 1000, 20, rl.RayWhite)
	} else if m.Lore_Display_Frame < 300 {
		rl.DrawText("As you step into this breathtaking realm, a world of unparalleled beauty unfolds before your eyes.", 90, 200, 35, rl.RayWhite)
		rl.DrawText("Press Enter to skip", 50, 1000, 20, rl.RayWhite)
	} else if m.Lore_Display_Frame < 600 && m.Lore_Display_Frame >= 300 {
		rl.DrawText("As you step into this breathtaking realm, a world of unparalleled beauty unfolds before your eyes.", 90, 200, 35, rl.RayWhite)
		rl.DrawText("It is a place where nature's wonders intertwine with the ethereal,", 400, 300, 35, rl.RayWhite)
		rl.DrawText("Press Enter to skip", 50, 1000, 20, rl.RayWhite)
	} else if m.Lore_Display_Frame < 900 && m.Lore_Display_Frame >= 600 {
		rl.DrawText("As you step into this breathtaking realm, a world of unparalleled beauty unfolds before your eyes.", 90, 200, 35, rl.RayWhite)
		rl.DrawText("It is a place where nature's wonders intertwine with the ethereal,", 400, 300, 35, rl.RayWhite)
		rl.DrawText("where luminescent flora illuminates the landscape with an otherworldly glow.", 300, 400, 35, rl.RayWhite)
		rl.DrawText("Press Enter to skip", 50, 1000, 20, rl.RayWhite)
	} else if m.Lore_Display_Frame < 1200 && m.Lore_Display_Frame >= 900 {
		rl.DrawText("As you step into this breathtaking realm, a world of unparalleled beauty unfolds before your eyes.", 90, 200, 35, rl.RayWhite)
		rl.DrawText("It is a place where nature's wonders intertwine with the ethereal,", 400, 300, 35, rl.RayWhite)
		rl.DrawText("where luminescent flora illuminates the landscape with an otherworldly glow.", 300, 400, 35, rl.RayWhite)
		rl.DrawText("Towering mountains pierce the heavens, and crystalline rivers serenade the ears with their gentle melodies.", 10, 500, 35, rl.RayWhite)
		rl.DrawText("Press Enter to skip", 50, 1000, 20, rl.RayWhite)
	} else if m.Lore_Display_Frame <= 1400 && m.Lore_Display_Frame >= 1200 {
		rl.DrawText("Press Enter to skip", 50, 1000, 20, rl.RayWhite)
	} else if m.Lore_Display_Frame > 1400 {
		m.Lore_Display_Frame = 0
		m.menu = 2
	}

	if rl.IsKeyPressed(rl.KeyEnter) {
		m.Lore_Display_Frame = 0
		m.menu = 2
	}

	rl.EndDrawing()
}