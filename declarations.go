package main


var (
	fPath = "//studentDB.toml"
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
	FName   string
	LName   string
	Age     int
	Grade   int
	Citizen bool
}