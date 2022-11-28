package main

import (
	"fmt"
	"log"
	"math/rand"
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
	// open()
	// if err := file.Close(); err != nil {
	// 	// failed to close the file
	// 	log.Fatal(err)
	// }
}

/*
Creates maps for all info. Key = id && Value = info
Recives a slice of Student
! It also recives a temp rand.Rand to generate ids. Create a better sys
*/
func createMaps(s []Student, r *rand.Rand) {
	studentMap, studentNameMap, studentLastNameMap, studentAgeMap, studentCitizenMap, studentGradeMap = make(map[string]Student), make(map[int]string), make(map[int]string), make(map[int]int), make(map[int]bool), make(map[int]int)
	for x, y := range s {
		id := r.Intn(99_000_000 - 10_000_000)+ 10_000_000 // Only 8 digit ids
		studentMap[strconv.FormatInt(int64(id), 10)] = y
		studentNameMap[id] = s[x].FName
		studentLastNameMap[id] = s[x].LName
		studentAgeMap[id] = s[x].Age
		studentGradeMap[id] = s[x].Grade
		studentCitizenMap[id] = s[x].Citizen
	}
}

/*
Recives value to search for and a sorted array of ints to search in.
Returns the index value of the desired value.
*/
func binarySearch(v int, s []int) int {
	indMin, indMax := 0, len(s)-1

	for indMin < indMax {
		indMid := int(indMin + (indMax-indMin)/2)

		if !(s[indMid] >= v) {
			indMin = indMid + 1
		} else {
			indMax = indMid
		}

	}

	return indMin
}

func main() {
	readedMap = read()
	tempRngStudents()
	write(studentMap)

	for exec {
		fmt.Println(rl.GetMouseX(), rl.GetMouseY())

		update()
		render()

	}
}
