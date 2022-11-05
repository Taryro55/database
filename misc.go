package main

import "os"

var (
	file *os.File
	studentSlice []Student
		studentIdSlice []int
		studentNameSlice []string
		studentLastNameSlice []string
		studentAgeSlice []int
		studentGradeSlice []int
		studentCitizenSlice []bool

)
type Student struct {
	id 		int
	fName   string
	lName   string
	age     int
	grade   int
	citizen bool
}