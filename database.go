package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path"
	"time"
	// toml "github.com/BurntSushi/toml"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func init() {
	var PATH, _ = os.Getwd()
	path := path.Join(PATH, "//studentDB.toml")

	// Creates a file if it doesnt exist
	if _, err := os.Stat(path); err != nil {
		f, err := os.Create(path)
		if err != nil {
			log.Fatal(err)
		}
		file = f
	}


	// ! Temp rng Students generator

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	b := false

	for x := 0; x < 10; x++ {
		randBool := r.Intn(2)
		if randBool == 0 {
			b = false
		} else if randBool == 1 {
			b = true
		}
		y := Student{
			"x",
			"y",
			r.Intn(99),
			r.Intn(12),
			b}
		studentSlice = append(studentSlice, y)
	}

	createMaps(studentSlice, r)

}

func quit() {
	if err := file.Close(); err != nil {
		// failed to close the file
		log.Fatal(err)
	}
}


func createMaps(s []Student, r *rand.Rand) {
	studentAgeMap, studentCitizenMap, studentGradeMap = make(map[int]int), make(map[int]bool), make(map[int]int)
	for x := range s {
		id := r.Intn(5000)
		studentIdSlice = append(studentIdSlice, id)
		studentAgeMap[id] = s[x].age
		studentGradeMap[id] = s[x].grade
		studentCitizenMap[id] = s[x].citizen
	}
}

func main() {
	



	fmt.Println("\n", studentAgeMap, "\n", studentGradeMap, "\n", studentCitizenMap, "\n")
	
	fmt.Println("age: ")
	printMapMod(bSortInt(studentAgeMap))
	fmt.Println("grade: ")
	printMapMod(bSortInt(studentGradeMap))
	fmt.Println("pr: ")
	fmt.Println(bSortBool(studentCitizenMap))
	
}
