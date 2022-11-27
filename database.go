package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path"
	"time"
	// toml "github.com/BurntSushi/toml"
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
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	for x := 0; x < 10; x++ {
		y := Student{
			"x",
			"y",
			r.Intn(99),
			r.Intn(12),
			true}
		studentSlice = append(studentSlice, y)
	}

	createMaps(studentSlice, r)



	// fmt.Println("\n", studentAgeMap, "\n", studentGradeMap, "\n")
	
	// fmt.Println("age: ")
	// printMapMod(bSortInt(studentAgeMap))
	// fmt.Println("grade: ")
	// printMapMod(bSortInt(studentGradeMap))
	fmt.Println("pr: ")
	fmt.Println(boolMapToIntMap(studentCitizenMap))

	printMapMod(bSortInt(boolMapToIntMap(studentCitizenMap)))

	

	// db := load()
	// printDB(db)
	// print the current values (starting values would be nothing)
	// have an add function
	// asks for everything and adds that to a Student struct, then it adds it to an array, and to a toml
	// have a remove func
	// asks for id to remove and edits the toml
	// have a refresh func
	// re-reads the toml file.

	// IFF using raylib, have a click on an area to sort by that
}
