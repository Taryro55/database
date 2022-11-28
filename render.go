package main

import (
	"math"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func update() {
	exec = !rl.WindowShouldClose()
	y = 200

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
	cBackground, _ := ParseHexColor("#121212")
	cBoxed, _ := ParseHexColor("#2c2c2c")
	rl.ClearBackground(cBackground)

	rl.DrawRectangleRounded(recForegound, 0.025, 0, cBoxed)

	rl.DrawText("Student Database DMCI", 65, 27, 50, rl.White)
	rl.DrawText("DATE", 1000, 27, 50, rl.White)

	rl.DrawLineEx(rl.Vector2{55, 140}, rl.Vector2{float32(width), 140}, 3, cBackground)

	rl.DrawText("Students", 65, 100, 37, rl.White)
	rl.DrawText("+ Add", 775, 104, 32, rl.White)
	rl.DrawText("- Del", 975, 104, 32, rl.White)
	rl.DrawText("% Mod", 1185, 104, 32, rl.White)

	rl.DrawLineEx(rl.Vector2{55, 190}, rl.Vector2{float32(width), 190}, 3, cBackground)

	sliceOfSecondLane := []string{"Id", "First Name", "Last Name", "Age", "Grade", "Citizenship"}
	collsX := getCollsX()

	for i, v := range sliceOfSecondLane {
		if i > len(sliceOfSecondLane) || i > len(collsX) {
			break
		}
		rl.DrawText(v, collsX[i], 153, 30, rl.White)
	}

	//
	for k, v := range readedMap {
		if len(p) == len(readedMap) {
			break
		}
		p = append(p, k)
		studentNameSlice = append(studentNameSlice, v.FName)
		studentLastNameSlice = append(studentLastNameSlice, v.LName)
		studentAgeSlice = append(studentAgeSlice, strconv.FormatInt(int64(v.Age), 10))
		studentGradeSlice = append(studentGradeSlice, strconv.FormatInt(int64(v.Grade), 10))
		studentCitizenSlice = append(studentCitizenSlice, strconv.FormatBool(v.Citizen))

		studentSliceOfSlices = append(studentSliceOfSlices, p, studentNameSlice, studentLastNameSlice, studentAgeSlice, studentGradeSlice, studentCitizenSlice)
	}

	//drawColl(studentNameSlice, 250)

	for i, slice := range studentSliceOfSlices {
		if i > len(slice) || i > len(collsX) {
			break
		}
		drawColl(slice, collsX[i])
	}

}

func getCollsX() []int32 {
	tSlice := make([]int32, 5)
	for x := 0; x > 6; x++ {
		y := int32(quadFunc(float64(x), 6.25, -73.519, 263.889, -42.811, 70.198))
		tSlice = append(tSlice, y)
	}
	return tSlice
}

/*
Returns the quadratic of the input
*/
func quadFunc(x, a, b, c, d, e float64) float64 {
	z := (a*math.Pow(x, 4) + b*math.Pow(x, 3) + c*math.Pow(x, 2) + d*x + e)
	return z
}

func drawColl(s []string, x int32) {
	y = 200

	for _, v := range s {
		rl.DrawText(v, x, y, 30, rl.White)
		y = y + 50
	}

	s = scroll(s)
}
