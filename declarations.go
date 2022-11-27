package main

import "os"

var (
	file *os.File
	studentSlice []Student
		studentIdSlice 			[]int
		studentNameMap 			map[int]string
		studentLastNameMap 		map[int]string
		studentAgeMap 			map[int]int
		studentGradeMap 		map[int]int
		studentCitizenMap 		map[int]bool

)

type MapMod struct {
	key []int
	val []int
}
type Student struct {
	fName   string
	lName   string
	age     int
	grade   int
	citizen bool
}