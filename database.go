package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func init() {
	var PATH, _ = os.Getwd()
	path := path.Join(PATH, fPath)

	// Creates a file if it doesnt exist
	if _, err := os.Stat(path); err != nil {
		_, err := os.Create(path)
		if err != nil {
			// failed to create/open the file
			log.Fatal(err)
		}
	}

	rl.InitWindow(width, height, WINDOW_TITLE)
	rl.SetTargetFPS(60)
	rl.SetMouseScale(1.0, 1.0)

}

func quit() {
	write(studentMap)
	rl.CloseWindow()
}

/*
Creates maps for all info. Key = id && Value = info
Recives a slice of Student
! It also recives a temp rand.Rand to generate ids. Create a better sys
*/
func createMaps(s []Student, ids []int) {
	studentMap, studentNameMap, studentLastNameMap, studentAgeMap, studentCitizenMap, studentGradeMap = make(map[string]Student), make(map[int]string), make(map[int]string), make(map[int]int), make(map[int]bool), make(map[int]int)
	for x, y := range s {
		// id := r.Intn(9_999_999-1_000_000) + 1_000_000 // Only 8 digit ids
		studentMap[strconv.FormatInt(int64(ids[x]), 10)] = y
		studentNameMap[ids[x]] = s[x].FName
		studentLastNameMap[ids[x]] = s[x].LName
		studentAgeMap[ids[x]] = s[x].Age
		studentGradeMap[ids[x]] = s[x].Grade
		studentCitizenMap[ids[x]] = s[x].Citizen
	}
}
func createSlices() {
	for k, v := range readedMap {
		if len(studentIdSlice) == len(readedMap) {
			break
		}
		fmt.Println(k)
		studentIdSlice = append(studentIdSlice, k)
		studentNameSlice = append(studentNameSlice, v.FName)
		studentLastNameSlice = append(studentLastNameSlice, v.LName)
		studentAgeSlice = append(studentAgeSlice, strconv.FormatInt(int64(v.Age), 10))
		studentGradeSlice = append(studentGradeSlice, strconv.FormatInt(int64(v.Grade), 10))
		studentCitizenSlice = append(studentCitizenSlice, strconv.FormatBool(v.Citizen))
	}
}

func main() {
	fmt.Println("\033[H\033[2J")

	studs, ids := []Student{}, []int{}
	readedMap = read()
	for k, v := range readedMap {
		id, _ := strconv.Atoi(k)
		studs = append(studs, v)
		ids = append(ids, id)
	}

	createMaps(studs, ids)
	createSlices()

	for exec {
		// fmt.Println(rl.GetMouseX(), rl.GetMouseY())

		// fmt.Println(studentAgeMap, studentLastNameMap, studentNameMap, studentGradeMap, studentCitizenMap)
		update()
		render()
		loops++
	}
	updateMainMap()
	quit()
}
