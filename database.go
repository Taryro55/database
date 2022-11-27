package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path"
	"strconv"
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
	studentMap, studentNameMap, studentLastNameMap, studentAgeMap, studentCitizenMap, studentGradeMap = make(map[string]Student), make(map[int]string), make(map[int]string), make(map[int]int), make(map[int]bool), make(map[int]int)
	for x, y := range s {
		id := r.Intn(5000)
		studentMap[strconv.FormatInt(int64(id), 10)] = y
		studentNameMap[id] = s[x].FName
		studentLastNameMap[id] = s[x].LName
		studentAgeMap[id] = s[x].Age
		studentGradeMap[id] = s[x].Grade
		studentCitizenMap[id] = s[x].Citizen
	}
}

/*
Recives
	value to search for
	sorted array of ints to search
	indMin

Returns the index value of the desired value
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

func searchPrompt() {
	z := ""
	fmt.Scanln(">> Search: ", &z)
}

func main() {
	a := read()
	fmt.Println("read ", a)
	tempRngStudents()
	write(studentMap)


	// ! Proof of work for binary search
	searchVal := 45
	slicetoSearch := []int{1, 3, 6, 10, 15, 21, 28, 36, 45, 55}
	x := binarySearch(searchVal, slicetoSearch)
		
	if searchVal == slicetoSearch[x] {
		fmt.Println("The value ", searchVal, " is on index ", x)
	}

}
