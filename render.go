package main

import (
	"math"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func update() {
	exec = !rl.WindowShouldClose()
	y = 200
	secsSinceStart = int(math.Round(float64(loops / 60)))

	

	if mainMenuCooldown.Pressed && mainMenuCooldown.Loops+7 == loops {
		mainMenuCooldown.Pressed = false
		searchMenu = false
		// letterCount = 0
		inputText = []int{}
		rl.SetMouseCursor(0)
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

	if searchMenu && rl.IsKeyPressed(rl.KeyEnter) {
		searchMenu = false

		// if conversion to int != err, means its int, so it can be searched
		searchFor, _ := strconv.Atoi(unixSliceToStr(inputText))
		search(searchFor)
		// search(unixSliceToStr(strToIntSlice(inputText)))

		// letterCount = 0
		inputText = []int{}
		rl.SetMouseCursor(0)
	}

	// fmt.Println(inputText, letterCount, strconv.FormatInt(int64(unixSliceToStr(inputText)), 10))

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

	if !button(60, 98, 230, 130) && !mainMenuCooldown.Pressed {
		rl.DrawText("Students", 65, 100, 37, rl.White)
	} else if button(60, 98, 230, 130) || mainMenuCooldown.Pressed {
		rl.DrawText("Students", 65, 100, 37, cPrimary)

		if !mainMenuCooldown.Pressed && rl.IsMouseButtonPressed(0) {
			mainMenuCooldown = Cooldown{true, secsSinceStart, loops}
		}
	}

	// * Search ID
	if !button(570, 98, 775, 130) && !searchCooldown.Pressed {
		rl.DrawText("? Search ID", 575, 104, 32, rl.White)
	} else if button(570, 98, 775, 130) || searchCooldown.Pressed {
		rl.DrawText("? Search ID", 575, 104, 32, cPrimary)

		if !searchCooldown.Pressed && rl.IsMouseButtonPressed(0) {
			searchCooldown = Cooldown{true, secsSinceStart, loops}
			searchMenu = true
		}
		inputBox(searchMenu, 520, 400, 300, 50, 40, 7, false, true)
	}

	// * Add
	if !button(870, 98, 970, 130) && !addCooldown.Pressed {
		rl.DrawText("+ Add", 875, 104, 32, rl.White)
	} else if button(870, 98, 970, 130) || addCooldown.Pressed {
		rl.DrawText("+ Add", 875, 104, 32, cPrimary)
		if !addCooldown.Pressed && rl.IsMouseButtonPressed(0) {
			addCooldown = Cooldown{true, secsSinceStart, loops}
		}
	}

	// * Del
	if !button(990, 98, 1080, 130) && !delCooldown.Pressed {
		rl.DrawText("- Del", 1000, 104, 32, rl.White)
	} else if button(990, 98, 1080, 130) || delCooldown.Pressed {
		rl.DrawText("- Del", 1000, 104, 32, cPrimary)
		if !delCooldown.Pressed && rl.IsMouseButtonPressed(0) {
			delCooldown = Cooldown{true, secsSinceStart, loops}
		}
	}

	// * Mod
	if !button(1115, 98, 1230, 130) && !modCooldown.Pressed {
		rl.DrawText("% Mod", 1125, 104, 32, rl.White)
	} else if button(1115, 98, 1230, 130) || modCooldown.Pressed {
		rl.DrawText("% Mod", 1125, 104, 32, cPrimary)
		if !modCooldown.Pressed && rl.IsMouseButtonPressed(0) {
			modCooldown = Cooldown{true, secsSinceStart, loops}
		}
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

	if !searchMenu {
		studentIdSlice = drawColl(studentIdSlice, collsX[0])
		studentNameSlice = drawColl(studentNameSlice, collsX[1])
		studentLastNameSlice = drawColl(studentLastNameSlice, collsX[2])
		studentAgeSlice = drawColl(studentAgeSlice, collsX[3])
		studentGradeSlice = drawColl(studentGradeSlice, collsX[4])
		studentCitizenSlice = drawColl(studentCitizenSlice, collsX[5])
	}

}

func inputBox(menu bool, x, y, w, h float32, font, maxChar int, letters, numbers bool) {
	textBox := rl.Rectangle{x, y, w, h}
	onInputBox := false
	maxLettUnix, minLettUnix, maxNumbUnix, minNumbUnix := 0, 0, 0, 0
	letterCount    := 0
	framesCounter  := 0


	if rl.CheckCollisionPointRec(rl.GetMousePosition(), textBox) {
		onInputBox = true
	} else if !rl.CheckCollisionPointRec(rl.GetMousePosition(), textBox) {
		onInputBox = false
	}

	if letters {
		maxLettUnix = 90
		minLettUnix = 65
	}
	if numbers {
		maxNumbUnix = 58
		minNumbUnix = 48
	}

	if menu {
		rl.DrawRectangleRec(textBox, rl.LightGray)
		rl.DrawText(unixSliceToStr(inputText), textBox.ToInt32().X, textBox.ToInt32().Y, int32(font), rl.Maroon)

		if onInputBox {
			framesCounter++
			rl.SetMouseCursor(2)
			key := rl.GetKeyPressed()

			for key > 0 { // Manages addition of text
				if (inBetween(int(key), minLettUnix, maxLettUnix) || inBetween(int(key), minNumbUnix, maxNumbUnix)) && (letterCount <= maxChar) {
					letterCount++
					inputText = append(inputText, int(key))
				}
				key = rl.GetKeyPressed()
			}

			if rl.IsKeyPressed(rl.KeyBackspace) { 	// Manages deletion of text
				if letterCount <= 9 && letterCount >= 0 && len(inputText) != 0 {
					letterCount--
					inputText = inputText[:len(inputText)-1]
				}
			}

			// Manages the Drawing while mouse on the textbos
			rl.DrawRectangleLines(int32(textBox.X), int32(textBox.Y), int32(textBox.Width), int32(textBox.Height), rl.Red)
			if letterCount <= maxChar {
				if framesCounter/20%2 == 0 {
					rl.DrawText("_", textBox.ToInt32().X + 8 + rl.MeasureText(unixSliceToStr(inputText), int32(font)), textBox.ToInt32().Y + 12, 40, rl.Maroon)
				}
			}

		} else {
			rl.SetMouseCursor(0)
			framesCounter = 0
		}
	}
}
