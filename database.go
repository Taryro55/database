package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path"
	"time"
	// rl "github.com/gen2brain/raylib-go/raylib"
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

}

func randStr(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func tempRngStudents() {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	b := false
	rand.Seed(time.Now().UnixNano())

	for x := 0; x < 3; x++ {
		randBool := r.Intn(2)
		if randBool == 0 {
			b = false
		} else if randBool == 1 {
			b = true
		}
		y := Student{
			randStr(5),
			randStr(8),
			r.Intn(99),
			r.Intn(12),
			b}
		studentSlice = append(studentSlice, y)
	}
	createMaps(studentSlice, r)
}

func quit() {
	// open()
	// if err := file.Close(); err != nil {
	// 	// failed to close the file
	// 	log.Fatal(err)
	// }
}

func createMaps(s []Student, r *rand.Rand) {
	studentNameMap, studentLastNameMap, studentAgeMap, studentCitizenMap, studentGradeMap = make(map[int]string), make(map[int]string), make(map[int]int), make(map[int]bool), make(map[int]int)
	for x := range s {
		id := r.Intn(5000)
		studentIdSlice = append(studentIdSlice, id)
		studentNameMap[id] = s[x].FName
		studentLastNameMap[id] = s[x].LName
		studentAgeMap[id] = s[x].Age
		studentGradeMap[id] = s[x].Grade
		studentCitizenMap[id] = s[x].Citizen
	}
}

func main() {
	a := load()
	fmt.Println(a)
	tempRngStudents()
	write(studentSlice)

	// fmt.Println(studentNameMap, studentLastNameMap)

	// fmt.Println("\n", studentAgeMap, "\n", studentGradeMap, "\n", studentCitizenMap, "\n")

	// fmt.Println("age: ")
	// printMapMod(bSortInt(studentAgeMap))
	// fmt.Println("grade: ")
	// printMapMod(bSortInt(studentGradeMap))
	// fmt.Println("pr: ")
	// fmt.Println(bSortBool(studentCitizenMap))

}
