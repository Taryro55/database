package main

import (
	// "fmt"
	"fmt"
	"math"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func update() {
	exec = !rl.WindowShouldClose()
	y = 200
	secsSinceStart = int(math.Round(float64(loops / 60)))

	if rl.CheckCollisionPointRec(rl.GetMousePosition(), textBox) {
		onInputBox = true
	} else if !rl.CheckCollisionPointRec(rl.GetMousePosition(), textBox) {
		onInputBox = false
	}

	if searchCooldown.Pressed && searchCooldown.Loops+7 == loops {
		searchCooldown.Pressed = false
	}
	if addCooldown.Pressed && addCooldown.Loops+7 == loops {
		addCooldown.Pressed = false
	}
	if delCooldown.Pressed && delCooldown.Loops+7 == loops {
		delCooldown.Pressed = false
	}
	if modCooldown.Pressed && modCooldown.Loops+7 == loops {
		modCooldown.Pressed = false
	}

	if showInputBox && rl.IsKeyPressed(rl.KeyEnter) {
		showInputBox = false
		search(intUnixSliceToInt(inputText))
		letterCount = 0
		inputText = []int{}
		rl.SetMouseCursor(0)
	}
	fmt.Println(inputText, letterCount)

	// fmt.Println(intUnixSliceToInt(inputText), letterCount)


	recBackground = rl.Rectangle{
		X:      float32(offsetX),
		Y:      float32(offsetY),
		Width:  float32(width - (2 * offsetX)),
		Height: float32(height - (2 * offsetY)),
	}

	recForegound = rl.Rectangle{
		X:      float32(55),
		Y:      float32(90),
		Width:  float32(width),
		Height: float32(height),
	}
}

func render() {
	rl.BeginDrawing()
	defer rl.EndDrawing()

	rl.ClearBackground(cBackground)

	rl.DrawRectangleRounded(recForegound, 0.025, 0, cBoxed)

	rl.DrawText("Student Database DMCI", 65, 27, 50, rl.White)
	rl.DrawText("DATE", 1000, 27, 50, rl.White)

	firstRow()
	otherRows()
}

func firstRow() {
	rl.DrawText("Students", 65, 100, 37, rl.White)

	// * Search ID
	if !button(570, 98, 730, 130) && !searchCooldown.Pressed {
		rl.DrawText("? Search ID", 575, 104, 32, rl.White)
	} else if button(570, 98, 730, 130) || searchCooldown.Pressed {
		rl.DrawText("? Search ID", 575, 104, 32, cPrimary)

		if !searchCooldown.Pressed {
			searchCooldown = Cooldown{true, secsSinceStart, loops}

			showInputBox = true

		}
	}
	if showInputBox {

		if onInputBox {
			rl.SetMouseCursor(2)

			key := rl.GetKeyPressed()
			for key > 0 {

				if (key >= 32) && (key <= 125) && (letterCount < 9) {
					letterCount++
					inputText = append(inputText, int(key))
				}
				key = rl.GetKeyPressed()
			}
			if rl.IsKeyPressed(rl.KeyBackspace) {
				if letterCount <= 9 && letterCount >= 0 && len(inputText) != 0 {
					letterCount--
					inputText = inputText[:len(inputText)-1]
				}
			}

		} else {
			rl.SetMouseCursor(0)
		}
		if onInputBox {
			framesCounter++
		} else {
			framesCounter = 0
		}

		rl.DrawRectangleRec(textBox, rl.LightGray)
		a := strconv.FormatInt(int64(intUnixSliceToInt(inputText)), 10)
		rl.DrawText(a, 300, 200, 30, rl.White)

		if onInputBox {
			rl.DrawRectangleLines(int32(textBox.X), int32(textBox.Y), int32(textBox.Width), int32(textBox.Height), rl.Red)
		}
	}

	// * Add
	if !button(870, 98, 970, 130) && !addCooldown.Pressed {
		rl.DrawText("+ Add", 875, 104, 32, rl.White)
	} else if button(870, 98, 970, 130) || addCooldown.Pressed {
		rl.DrawText("+ Add", 875, 104, 32, cPrimary)
		if !addCooldown.Pressed {
			addCooldown = Cooldown{true, secsSinceStart, loops}
		}
	}

	// * Del
	if !button(990, 98, 1080, 130) && !delCooldown.Pressed {
		rl.DrawText("- Del", 1000, 104, 32, rl.White)
	} else if button(990, 98, 1080, 130) || delCooldown.Pressed {
		rl.DrawText("- Del", 1000, 104, 32, cPrimary)
		if !delCooldown.Pressed {
			delCooldown = Cooldown{true, secsSinceStart, loops}
		}
	}

	// * Mod
	if !button(1115, 98, 1230, 130) && !modCooldown.Pressed {
		rl.DrawText("% Mod", 1125, 104, 32, rl.White)
	} else if button(1115, 98, 1230, 130) || modCooldown.Pressed {
		if !modCooldown.Pressed {
			modCooldown = Cooldown{true, secsSinceStart, loops}
		}
		rl.DrawText("% Mod", 1125, 104, 32, cPrimary)
	}
}

func otherRows() {
	rl.DrawLineEx(rl.Vector2{55, 140}, rl.Vector2{float32(width), 140}, 3, cBackground)

	collsX := getCollsX()
	sliceOfSecondLane := []string{"Id", "First Name", "Last Name", "Age", "Grade", "Citizenship"}
	for i, v := range sliceOfSecondLane {
		rl.DrawText(v, collsX[i], 153, 30, rl.White)
	}

	rl.DrawLineEx(rl.Vector2{55, 190}, rl.Vector2{float32(width), 190}, 3, cBackground)

	if !showInputBox {
		studentIdSlice = drawColl(studentIdSlice, collsX[0])
		studentNameSlice = drawColl(studentNameSlice, collsX[1])
		studentLastNameSlice = drawColl(studentLastNameSlice, collsX[2])
		studentAgeSlice = drawColl(studentAgeSlice, collsX[3])
		studentGradeSlice = drawColl(studentGradeSlice, collsX[4])
		studentCitizenSlice = drawColl(studentCitizenSlice, collsX[5])
	}

}
