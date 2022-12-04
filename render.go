package main

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func update() {
	updateMaps()

	exec = !rl.WindowShouldClose()
	y = 200
	secsSinceStart = int(math.Round(float64(loops / 60)))

	defaultMenu()
	mainMenuCooldown.Pressed = offPressed(mainMenuCooldown)
	searchCooldown.Pressed = offPressed(searchCooldown)
	addCooldown.Pressed = offPressed(addCooldown)
	delCooldown.Pressed = offPressed(delCooldown)
	modCooldown.Pressed = offPressed(modCooldown)

	searchCooldown = searchInputManager(searchCooldown)
	addCooldown = addInputManager(addCooldown)
	delCooldown = delInputManager(delCooldown)
	// checkInput(modCooldown)

	// fmt.Println(studentIdSlice)
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
			offMenus()
			resetInputBox()
			mainMenuCooldown = Cooldown{true, loops, true}
		}
	}

	// * Search ID
	if !button(570, 98, 775, 130) && !searchCooldown.Pressed {
		rl.DrawText("? Search ID", 575, 104, 32, rl.White)
	} else if button(570, 98, 775, 130) || searchCooldown.Pressed {
		rl.DrawText("? Search ID", 575, 104, 32, cPrimary)
		if !searchCooldown.Pressed && rl.IsMouseButtonPressed(0) {
			offMenus()
			resetInputBox()
			searchCooldown = Cooldown{true, loops, true}
		}
	}
	if searchCooldown.OnMenu {
		searchInput = inputBox(520, 400, 300, 50, 40, 7, false, true, "Enter ID", searchInput)
	}

	// * Add
	if !button(870, 98, 970, 130) && !addCooldown.Pressed {
		rl.DrawText("+ Add", 875, 104, 32, rl.White)
	} else if button(870, 98, 970, 130) || addCooldown.Pressed {
		rl.DrawText("+ Add", 875, 104, 32, cPrimary)
		if !addCooldown.Pressed && rl.IsMouseButtonPressed(0) {
			offMenus()
			resetInputBox()
			addCooldown = Cooldown{true, loops, true}
		}
	}
	if addCooldown.OnMenu {
		idAddInput = inputBox(320, 300, 300, 50, 40, 7, false, true, "Id", idAddInput)                    // Id
		nameAddInput = inputBox(320, 400, 300, 50, 40, 7, true, false, "Name", nameAddInput)              // Name
		lnameAddInput = inputBox(320, 500, 300, 50, 40, 7, true, false, "Surname", lnameAddInput)         // Last Name
		ageAddInput = inputBox(720, 300, 300, 50, 40, 2, false, true, "Age", ageAddInput)                 // Age
		gradeAddInput = inputBox(720, 400, 300, 50, 40, 2, false, true, "Grade", gradeAddInput)           // Grade
		citizenAddInput = inputBox(720, 500, 300, 50, 40, 5, true, false, "Citizenship", citizenAddInput) // Citizenship
	}

	// * Del
	if !button(990, 98, 1080, 130) && !delCooldown.Pressed {
		rl.DrawText("- Del", 1000, 104, 32, rl.White)
	} else if button(990, 98, 1080, 130) || delCooldown.Pressed {
		rl.DrawText("- Del", 1000, 104, 32, cPrimary)
		if !delCooldown.Pressed && rl.IsMouseButtonPressed(0) {
			offMenus()
			resetInputBox()
			delCooldown = Cooldown{true, loops, true}
		}
	}
	if delCooldown.OnMenu {
		delInput = inputBox(520, 400, 300, 50, 40, 7, false, true, "Enter ID", delInput)
	}

	// * Mod
	if !button(1115, 98, 1230, 130) && !modCooldown.Pressed {
		rl.DrawText("% Mod", 1125, 104, 32, rl.White)
	} else if button(1115, 98, 1230, 130) || modCooldown.Pressed {
		rl.DrawText("% Mod", 1125, 104, 32, cPrimary)
		if !modCooldown.Pressed && rl.IsMouseButtonPressed(0) {
			offMenus()
			resetInputBox()
			modCooldown = Cooldown{true, loops, true}
		}
	}
	if modCooldown.OnMenu {
		// inputBox(520, 400, 300, 50, 50, 7, true, true, inputBox())
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

	if mainMenuCooldown.OnMenu {
		studentIdSlice = drawColl(studentIdSlice, collsX[0])
		studentNameSlice = drawColl(studentNameSlice, collsX[1])
		studentLastNameSlice = drawColl(studentLastNameSlice, collsX[2])
		studentAgeSlice = drawColl(studentAgeSlice, collsX[3])
		studentGradeSlice = drawColl(studentGradeSlice, collsX[4])
		studentCitizenSlice = drawColl(studentCitizenSlice, collsX[5])
	}

}

// add a parameter to specify where to store the value. The Menus that only use one input box can store it on the general, the others can have their own dedicated variables to store declarated on the declarations.go
func inputBox(x, y, w, h float32, font, maxChar int, letters, numbers bool, text string, input Input) Input {
	maxChar--
	if input.InputText == nil {
		input.InputText = make([]int, 0)
	}
	textBox := rl.Rectangle{x, y, w, h}
	maxLettUnix, minLettUnix, maxNumbUnix, minNumbUnix := 0, 0, 0, 0

	if rl.CheckCollisionPointRec(rl.GetMousePosition(), textBox) {
		input.OnInputBox = true
	} else if !rl.CheckCollisionPointRec(rl.GetMousePosition(), textBox) {
		input.OnInputBox = false
	}

	if letters {
		maxLettUnix = 90
		minLettUnix = 65
	}
	if numbers {
		maxNumbUnix = 58
		minNumbUnix = 48
	}

	rl.DrawRectangleRec(textBox, rl.LightGray)
	rl.DrawText(unixSliceToStr(input.InputText), textBox.ToInt32().X, textBox.ToInt32().Y, int32(font), cPrimary)

	if input.OnInputBox {
		input.FramesCounter++
		rl.SetMouseCursor(2)
		key := rl.GetKeyPressed()

		for key > 0 { // Manages addition of text
			if (inBetween(int(key), minLettUnix, maxLettUnix) || inBetween(int(key), minNumbUnix, maxNumbUnix)) && (input.LetterCount <= maxChar) {
				input.LetterCount++
				input.InputText = append(input.InputText, int(key))
			}
			key = rl.GetKeyPressed()
		}

		if rl.IsKeyPressed(rl.KeyBackspace) { // Manages deletion of text
			if input.LetterCount <= 9 && input.LetterCount >= 0 && len(input.InputText) != 0 {
				input.LetterCount--
				input.InputText = input.InputText[:len(input.InputText)-1]
			}
		}

		// Manages the Drawing while mouse on the textbos
		rl.DrawRectangleLines(int32(textBox.X), int32(textBox.Y), int32(textBox.Width), int32(textBox.Height), rl.Red)
		if input.LetterCount <= maxChar {
			if input.FramesCounter/20%2 == 0 {
				rl.DrawText("_", textBox.ToInt32().X+8+rl.MeasureText(unixSliceToStr(input.InputText), int32(font)), textBox.ToInt32().Y+12, int32(font), rl.Maroon)
			}
		}

	} else {
		if input.LetterCount == 0 {
			rl.DrawText(text, textBox.ToInt32().X+8+rl.MeasureText(unixSliceToStr(input.InputText), int32(font)), textBox.ToInt32().Y+12, int32(font)-5, rl.White)
		}

		rl.SetMouseCursor(0)
		input.FramesCounter = 0
	}
	return input
}

func searchInputManager(c Cooldown) Cooldown {
	if c.OnMenu && rl.IsKeyPressed(rl.KeyEnter) {
		c.OnMenu = false
		searchFor, _ := strconv.Atoi(unixSliceToStr(searchInput.InputText))
		moveSliceToTop(search(searchFor))
		resetInputBox()
	}
	return c
}

func addInputManager(c Cooldown) Cooldown {
	if c.OnMenu && rl.IsKeyPressed(rl.KeyEnter) {
		c.OnMenu = false

		// Checks if the input values are whats expected.
		validId := false
		validAge := false
		validGrade := false
		validCitizen := false

		id, _ := strconv.Atoi(unixSliceToStr(idAddInput.InputText))
		age, _ := strconv.Atoi(unixSliceToStr(ageAddInput.InputText))
		citizen := strings.ToLower(unixSliceToStr(citizenAddInput.InputText))
		grade, _ := strconv.Atoi(unixSliceToStr(gradeAddInput.InputText))

		if fmt.Sprint(reflect.TypeOf(id/1_000_000)) == "int" && !sliceContains(unixSliceToStr(idAddInput.InputText), studentIdSlice) {
			validId = true
		}

		if inBetween(age, 0, 99) {
			validAge = true
		}
		if inBetween(grade, 0, 12) {
			validGrade = true
		}
		if citizen == "true" || citizen == "false" {
			validCitizen = true
		}

		if validId && validAge && validGrade && validCitizen {
			studentIdSlice = append(studentIdSlice, unixSliceToStr(idAddInput.InputText))
			studentNameSlice = append(studentNameSlice, unixSliceToStr(nameAddInput.InputText))
			studentLastNameSlice = append(studentLastNameSlice, unixSliceToStr(lnameAddInput.InputText))
			studentAgeSlice = append(studentAgeSlice, unixSliceToStr(ageAddInput.InputText))
			studentGradeSlice = append(studentGradeSlice, unixSliceToStr(gradeAddInput.InputText))
			studentCitizenSlice = append(studentCitizenSlice, strings.ToLower(unixSliceToStr(citizenAddInput.InputText)))
		} else {
			// TODO print warning that data type was incorrect
		}

		resetInputBox()
	}
	return c
}

func delInputManager(c Cooldown) Cooldown {
	if c.OnMenu && rl.IsKeyPressed(rl.KeyEnter) {
		c.OnMenu = false
		searchFor, _ := strconv.Atoi(unixSliceToStr(delInput.InputText))
		index := search(searchFor)
		studentIdSlice = sliceStrDelete(index, studentIdSlice)
		studentAgeSlice = sliceStrDelete(index, studentAgeSlice)
		studentCitizenSlice = sliceStrDelete(index, studentCitizenSlice)
		studentGradeSlice = sliceStrDelete(index, studentGradeSlice)
		studentNameSlice = sliceStrDelete(index, studentNameSlice)
		studentLastNameSlice = sliceStrDelete(index, studentLastNameSlice)
		resetInputBox()
	}
	return c
}

func resetInputBox() {

	searchInput = resetInput(searchInput)
	delInput = resetInput(delInput)

	ageAddInput = resetInput(ageAddInput)
	citizenAddInput = resetInput(citizenAddInput)
	gradeAddInput = resetInput(gradeAddInput)
	idAddInput = resetInput(idAddInput)
	nameAddInput = resetInput(nameAddInput)
	lnameAddInput = resetInput(lnameAddInput)

	rl.SetMouseCursor(0)
}

func defaultMenu() {
	if !mainMenuCooldown.OnMenu && !searchCooldown.OnMenu && !addCooldown.OnMenu && !delCooldown.OnMenu && !modCooldown.OnMenu {
		mainMenuCooldown.OnMenu = true
	}
}

func offMenus() {
	mainMenuCooldown.OnMenu, searchCooldown.OnMenu, addCooldown.OnMenu, delCooldown.OnMenu, modCooldown.OnMenu = false, false, false, false, false
}

func offPressed(c Cooldown) bool {
	if c.Pressed && c.Loops+7 == loops {
		c.Pressed = false
	}
	return c.Pressed
}
