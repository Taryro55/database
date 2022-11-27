package main


var (
	fPath 						= "//studentDB.toml"
	studentMap 					map[string]Student
	studentSlice 				[]Student
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