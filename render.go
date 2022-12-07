package main

import (
	"fmt"
	"math"
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

	// idCooldown.Pressed = offPressed(idCooldown)

	searchCooldown = searchInputManager(searchCooldown)
	addCooldown = addInputManager(addCooldown)
	delCooldown = delInputManager(delCooldown)
	modCooldown = modInputManager(modCooldown)

	// errorSlice = append(errorSlice, "TEST", "TEST2")
	if errorSlice != nil {
		if !errorLooped {
			for _, v := range errorSlice {
				errorText = errorText + "\n" + v
			}
			errorLooped = true
		}
		rl.DrawText(errorText, 750, -20, 25, cError)

	}
}

func render() {
	rl.BeginDrawing()
	defer rl.EndDrawing()

	rl.ClearBackground(cBackground)

	rl.DrawRectangleRounded(recForegound, 0.025, 0, cBoxed)

	rl.DrawText("Student Database DMCI", 65, 27, 50, rl.White)

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
		nameAddInput = inputBox(320, 400, 300, 50, 40, 10, true, false, "Name", nameAddInput)             // Name
		lnameAddInput = inputBox(320, 500, 300, 50, 40, 10, true, false, "Surname", lnameAddInput)        // Last Name
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
		oldIdModInput = inputBox(520, 200, 300, 50, 40, 7, false, true, "Id To Mod", oldIdModInput)

		idModInput = inputBox(320, 300, 300, 50, 40, 7, false, true, "Id", idModInput)                    // Id
		nameModInput = inputBox(320, 400, 300, 50, 40, 10, true, false, "Name", nameModInput)             // Name
		lnameModInput = inputBox(320, 500, 300, 50, 40, 10, true, false, "Surname", lnameModInput)        // Last Name
		ageModInput = inputBox(720, 300, 300, 50, 40, 2, false, true, "Age", ageModInput)                 // Age
		gradeModInput = inputBox(720, 400, 300, 50, 40, 2, false, true, "Grade", gradeModInput)           // Grade
		citizenModInput = inputBox(720, 500, 300, 50, 40, 5, true, false, "Citizenship", citizenModInput) // Citizenship
	}
}

func updateCooldown(from int, to Cooldown) {
	switch from {
	case 0:
		idCooldown.OnMenu = true
	case 1:
		nameCooldown.OnMenu = true
	case 2:
		lnameCooldown.OnMenu = true
	case 3:
		ageCooldown.OnMenu = true
	case 4:
		gradeCooldown.OnMenu = true
	case 5:
		citizenCooldown.OnMenu = true
	}
}

func otherRows() {
	rl.DrawLineEx(rl.Vector2{55, 140}, rl.Vector2{float32(width), 140}, 3, cBackground)

	collsX := getCollsX()
	sliceOfSecondLane := []string{"Id", "First Name", "Last Name", "Age", "Grade", "Citizenship"}
	ys := []int32{100, 380, 625, 780, 1020, 1240}
	cools := []Cooldown{idCooldown, nameCooldown, lnameCooldown, ageCooldown, gradeCooldown, citizenCooldown}

	for i, v := range collsX {
		if !button(v, 150, ys[i], 180) && !cools[i].Pressed {
			rl.DrawText(sliceOfSecondLane[i], v, 153, 32, rl.White)
		} else if button(v, 98, ys[i], 180) || cools[i].Pressed {
			rl.DrawText(sliceOfSecondLane[i], v, 153, 32, cPrimary)
			if !cools[i].Pressed && rl.IsMouseButtonPressed(0) {
				offMenus()
				old := make([]string, 0)
				for _, v := range studentIdSlice {
					old = append(old, v)
				}
				cools[i] = Cooldown{true, loops, true}
				updateCooldown(i, idCooldown)

				if idCooldown.OnMenu {
					studentIdSlice = intToStrSlice(bubbleSort(strToIntSlice(studentIdSlice)))

					studentNameSlice = resortString(old, studentIdSlice, studentNameSlice)
					studentLastNameSlice = resortString(old, studentIdSlice, studentLastNameSlice)
					studentAgeSlice = resortString(old, studentIdSlice, studentAgeSlice)
					studentGradeSlice = resortString(old, studentIdSlice, studentGradeSlice)
					studentCitizenSlice = resortString(old, studentIdSlice, studentCitizenSlice)
					idCooldown.OnMenu = false
				} else if nameCooldown.OnMenu {
					save()
					studentNameSlice, studentIdSlice = bSortString(studentNameMap).key, bSortString(studentNameMap).val

					studentLastNameSlice = resortString(old, studentIdSlice, studentLastNameSlice)
					studentAgeSlice = resortString(old, studentIdSlice, studentAgeSlice)
					studentGradeSlice = resortString(old, studentIdSlice, studentGradeSlice)
					studentCitizenSlice = resortString(old, studentIdSlice, studentCitizenSlice)
					nameCooldown.OnMenu = false
				} else if lnameCooldown.OnMenu {
					save()
					studentLastNameSlice, studentIdSlice = bSortString(studentLastNameMap).key, bSortString(studentLastNameMap).val

					studentNameSlice = resortString(old, studentIdSlice, studentNameSlice)
					studentAgeSlice = resortString(old, studentIdSlice, studentAgeSlice)
					studentGradeSlice = resortString(old, studentIdSlice, studentGradeSlice)
					studentCitizenSlice = resortString(old, studentIdSlice, studentCitizenSlice)
					lnameCooldown.OnMenu = false
				} else if ageCooldown.OnMenu {
					save()
					studentAgeSlice, studentIdSlice = intToStrSlice(bSortInt(studentAgeMap).key), intToStrSlice(bSortInt(studentAgeMap).val)

					studentNameSlice = resortString(old, studentIdSlice, studentNameSlice)
					studentLastNameSlice = resortString(old, studentIdSlice, studentLastNameSlice)
					studentGradeSlice = resortString(old, studentIdSlice, studentGradeSlice)
					studentCitizenSlice = resortString(old, studentIdSlice, studentCitizenSlice)
					ageCooldown.OnMenu = false
				} else if gradeCooldown.OnMenu {
					save()
					studentGradeSlice, studentIdSlice = intToStrSlice(bSortInt(studentGradeMap).key), intToStrSlice(bSortInt(studentGradeMap).val)

					studentNameSlice = resortString(old, studentIdSlice, studentNameSlice)
					studentLastNameSlice = resortString(old, studentIdSlice, studentLastNameSlice)
					studentAgeSlice = resortString(old, studentIdSlice, studentAgeSlice)
					studentCitizenSlice = resortString(old, studentIdSlice, studentCitizenSlice)
					gradeCooldown.OnMenu = false
				} else if citizenCooldown.OnMenu {
					save()
					studentCitizenSlice, studentIdSlice = boolSliceToString(binaryToBool(bSortBool(studentCitizenMap).val)), intToStrSlice(bSortBool(studentCitizenMap).key)

					studentNameSlice = resortString(old, studentIdSlice, studentNameSlice)
					studentLastNameSlice = resortString(old, studentIdSlice, studentLastNameSlice)
					studentAgeSlice = resortString(old, studentIdSlice, studentAgeSlice)
					studentGradeSlice = resortString(old, studentIdSlice, studentGradeSlice)
					citizenCooldown.OnMenu = false
				}
			}
		}
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
			if input.LetterCount <= maxChar+1 && input.LetterCount >= 0 && len(input.InputText) != 0 {
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
		errorSlice = []string{}
		errorLooped = false
		errorText = ""
		c.OnMenu = false
		searchFor, _ := strconv.Atoi(unixSliceToStr(searchInput.InputText))
		index := search(searchFor, true)
		if index != -1 {
			moveSliceToTop(index, studentIdSlice)
			moveSliceToTop(index, studentNameSlice)
			moveSliceToTop(index, studentLastNameSlice)
			moveSliceToTop(index, studentAgeSlice)
			moveSliceToTop(index, studentGradeSlice)
			moveSliceToTop(index, studentCitizenSlice)
		} else {
			errorSlice = append(errorSlice, "Cannot find that ID!")
		}
		resetInputBox()

	}
	return c
}

func addInputManager(c Cooldown) Cooldown {
	if c.OnMenu && rl.IsKeyPressed(rl.KeyEnter) {
		errorSlice = []string{}
		errorLooped = false
		errorText = ""
		c.OnMenu = false

		name := unixSliceToStr(nameAddInput.InputText)
		lname := unixSliceToStr(lnameAddInput.InputText)
		age, _ := strconv.Atoi(unixSliceToStr(ageAddInput.InputText))
		citizen := strings.ToLower(unixSliceToStr(citizenAddInput.InputText))
		grade, _ := strconv.Atoi(unixSliceToStr(gradeAddInput.InputText))

		if !sliceContains(unixSliceToStr(idAddInput.InputText), studentIdSlice) {
			idAddInput.Valid = true
		}
		if !(lname == "") || !(name == "") {
			nameAddInput.Valid = true
			lnameAddInput.Valid = true
		}
		if inBetween(age, 1, 99) {
			ageAddInput.Valid = true
		}
		if inBetween(grade, 1, 12) {
			gradeAddInput.Valid = true
		}
		if citizen == "true" || citizen == "false" {
			citizenAddInput.Valid = true
		}

		if idAddInput.Valid && nameAddInput.Valid && lnameAddInput.Valid && ageAddInput.Valid && gradeAddInput.Valid && citizenAddInput.Valid {
			studentIdSlice = append(studentIdSlice, unixSliceToStr(idAddInput.InputText))
			studentNameSlice = append(studentNameSlice, unixSliceToStr(nameAddInput.InputText))
			studentLastNameSlice = append(studentLastNameSlice, unixSliceToStr(lnameAddInput.InputText))
			studentAgeSlice = append(studentAgeSlice, unixSliceToStr(ageAddInput.InputText))
			studentGradeSlice = append(studentGradeSlice, unixSliceToStr(gradeAddInput.InputText))
			studentCitizenSlice = append(studentCitizenSlice, citizen)
		}

		if !idAddInput.Valid {
			errorSlice = append(errorSlice, "Id's cant be repeated.")
		}

		if !ageAddInput.Valid || !nameAddInput.Valid || !lnameAddInput.Valid {
			errorSlice = append(errorSlice, "Please fill out every input.")
		}

		if !gradeAddInput.Valid {
			errorSlice = append(errorSlice, "Grade must be between 1-12, try again.")
			// rl.DrawText("Grade must be between 1-12, try again.", 1000, 30, 20, cError)
		}
		if !citizenAddInput.Valid {
			errorSlice = append(errorSlice, "Citizenship can only be true or false.")
			// fmt.Println("Citizenship can only be true or false")
		}

		resetInputBox()
	}
	return c
}

func delInputManager(c Cooldown) Cooldown {
	if c.OnMenu && rl.IsKeyPressed(rl.KeyEnter) {
		errorSlice = []string{}
		errorLooped = false
		errorText = ""
		c.OnMenu = false
		searchFor, _ := strconv.Atoi(unixSliceToStr(delInput.InputText))
		index := search(searchFor, true)
		// fmt.Println(studentAgeSlice)
		if index == -1 {
			errorSlice = append(errorSlice, "Cannot find that ID!")
			resetInputBox()
			return c
		} else if index != -1 {
			studentIdSlice = sliceStrDelete(index, studentIdSlice)
			studentAgeSlice = sliceStrDelete(index, studentAgeSlice)
			studentCitizenSlice = sliceStrDelete(index, studentCitizenSlice)
			studentGradeSlice = sliceStrDelete(index, studentGradeSlice)
			studentNameSlice = sliceStrDelete(index, studentNameSlice)
			studentLastNameSlice = sliceStrDelete(index, studentLastNameSlice)
		}
		// fmt.Println(studentAgeSlice)
		resetInputBox()
		save()
	}
	return c
}

func modInputManager(c Cooldown) Cooldown {
	if c.OnMenu && rl.IsKeyPressed(rl.KeyEnter) {
		errorSlice = []string{}
		errorLooped = false
		errorText = ""
		c.OnMenu = false
		searchFor, _ := strconv.Atoi(unixSliceToStr(oldIdModInput.InputText))
		index := search(searchFor, true)
		save()

		fmt.Println(unixSliceToStr(idModInput.InputText))
		if unixSliceToStr(idModInput.InputText) != "" {
			studentIdSlice[index] = unixSliceToStr(idModInput.InputText)
		}
		if unixSliceToStr(nameModInput.InputText) != "" {
			studentNameSlice[index] = unixSliceToStr(nameModInput.InputText)
		}
		if unixSliceToStr(lnameModInput.InputText) != "" {
			studentLastNameSlice[index] = unixSliceToStr(lnameModInput.InputText)
		}
		if unixSliceToStr(ageModInput.InputText) != "" {
			studentAgeSlice[index] = unixSliceToStr(ageModInput.InputText)
		}
		if unixSliceToStr(gradeModInput.InputText) != "" {
			studentGradeSlice[index] = unixSliceToStr(gradeModInput.InputText)
		}
		if unixSliceToStr(citizenModInput.InputText) != "" {
			studentCitizenSlice[index] = unixSliceToStr(citizenModInput.InputText)
		}
		save()

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
	if !mainMenuCooldown.OnMenu && !searchCooldown.OnMenu && !addCooldown.OnMenu && !delCooldown.OnMenu && !modCooldown.OnMenu && !idCooldown.OnMenu && !nameCooldown.OnMenu && !lnameCooldown.OnMenu && !ageCooldown.OnMenu && !gradeCooldown.OnMenu && !citizenCooldown.OnMenu {
		mainMenuCooldown.OnMenu = true
	}
}

func offMenus() {
	mainMenuCooldown.OnMenu, searchCooldown.OnMenu, addCooldown.OnMenu, delCooldown.OnMenu, modCooldown.OnMenu = false, false, false, false, false
	idCooldown.OnMenu, nameCooldown.OnMenu, lnameCooldown.OnMenu, ageCooldown.OnMenu, gradeCooldown.OnMenu, citizenCooldown.OnMenu = false, false, false, false, false, false
}

func offPressed(c Cooldown) bool {
	if c.Pressed && c.Loops+7 == loops {
		c.Pressed = false
	}
	return c.Pressed
}
