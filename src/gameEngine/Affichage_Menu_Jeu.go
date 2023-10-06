package gameEngine

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

//---- Main game scene gestion & display -----//

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
		if perso.currentHealthPoint <= perso.maxHealthPoint - 30 {
			perso.currentHealthPoint += 30
		} else if perso.currentHealthPoint < perso.maxHealthPoint {
			perso.currentHealthPoint = perso.maxHealthPoint
		}
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
