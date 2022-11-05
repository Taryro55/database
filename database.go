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

func bSortInt(s []int) []int {
	for y := 0; y < len(s)-1; y++ {
		for v := 0; v < len(s)-y-1; v++ {
			if (s[v] > s[v+1]) {
				s[v], s[v+1] = s[v+1], s[v]
			}
		}
	}
	return s
}

func createSlices(s []Student) {
	for x := range s {
		studentIdSlice = append(studentIdSlice, s[x].id)
		studentNameSlice = append(studentNameSlice, s[x].fName)
		studentLastNameSlice = append(studentLastNameSlice, s[x].lName)
		studentAgeSlice = append(studentAgeSlice, s[x].age)
		studentGradeSlice = append(studentGradeSlice, s[x].grade)
		studentCitizenSlice = append(studentCitizenSlice, s[x].citizen)
	}
}

func main() {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	for x := 0; x < 50; x++ {
		y := Student{
			r.Intn(5000),
			"x",
			"y",
			r.Intn(99),
			r.Intn(12),
			true}
		studentSlice = append(studentSlice, y)
	}
	for c := range studentSlice {
		fmt.Println(studentSlice[c])
	}

	createSlices(studentSlice)

	fmt.Println(bSortInt(studentIdSlice))
	fmt.Println(bSortInt(studentAgeSlice))
	fmt.Println(bSortInt(studentGradeSlice))

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
