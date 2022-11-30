package main

import rl "github.com/gen2brain/raylib-go/raylib"

var (
	loops 						int
	secsSinceStart				int

	readedMap					map[string]Student
	fPath 					  = "//studentDB.toml"
	
	studentSlice 				[]Student
		studentIdSlice 			[]string
		studentNameSlice 		[]string
		studentLastNameSlice 	[]string	
		studentAgeSlice 		[]string
		studentGradeSlice 		[]string
		studentCitizenSlice 	[]string
	
	studentMap 					map[string]Student
		studentNameMap 			map[int]string
		studentLastNameMap 		map[int]string
		studentAgeMap 			map[int]int
		studentGradeMap 		map[int]int
		studentCitizenMap 		map[int]bool

)

const (
	HEIGHT      			  = int32(768)
	WINDOW_TITLE 			  = "School DB"
	RAND					  = 5
)

var (
	exec 					  =	true
	height 					  = HEIGHT
	width 					  = (height / 9) * 16
	offsetX 				  = width * 2 / 100
	offsetY 				  = height * 3 / 100
	recBackground 				rl.Rectangle
	recForegound 				rl.Rectangle

	cBackground, _ 		= ParseHexColor("#121212")
	cBoxed, _ = ParseHexColor("#2c2c2c")
	cPrimary, _ = ParseHexColor("#BB86FC")

	o string
	y int32

	searchCooldown 				Cooldown
	addCooldown 				Cooldown
	delCooldown 				Cooldown
	modCooldown 				Cooldown

	onInputBox 				  = false
	textBox					  = rl.Rectangle{300, 300, 300, 50}
	letterCount				  = 0
	framesCounter			    int
	showInputBox				bool
	inputText					[]int
	alphabeth				  = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	alphabethSlice			  = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
)

type MapMod struct {
	key 						[]int
	val 						[]int
}
type Student struct {
	FName   					string
	LName   					string
	Age     					int
	Grade  						int
	Citizen 					bool
}

type Cooldown struct {
	Pressed	 					bool 
	SecSinceCooldown			int
	Loops 						int
}