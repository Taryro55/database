package main

import rl "github.com/gen2brain/raylib-go/raylib"

var (
	loops          int
	secsSinceStart int

	readedMap map[string]Student
	fPath     = "//studentDB.toml"

	studentSlice         []Student
	studentIdSlice       []string
	studentNameSlice     []string
	studentLastNameSlice []string
	studentAgeSlice      []string
	studentGradeSlice    []string
	studentCitizenSlice  []string

	studentMap         map[string]Student
	studentNameMap     map[int]string
	studentLastNameMap map[int]string
	studentAgeMap      map[int]int
	studentGradeMap    map[int]int
	studentCitizenMap  map[int]bool
)

const (
	HEIGHT       = int32(768)
	WINDOW_TITLE = "School DB"
	RAND         = 5
)

var (
	exec          = true
	height        = HEIGHT
	width         = (height / 9) * 16
	offsetX       = width * 2 / 100
	offsetY       = height * 3 / 100
	recStudent    rl.Rectangle
	recBackground = rl.Rectangle{
		X:      float32(offsetX),
		Y:      float32(offsetY),
		Width:  float32(width - (2 * offsetX)),
		Height: float32(height - (2 * offsetY)),
	}
	recForegound = rl.Rectangle{
		X:      float32(55),
		Y:      float32(90),
		Width:  float32(width),
		Height: float32(height),
	}

	cBackground, _ = ParseHexColor("#121212")
	cBoxed, _      = ParseHexColor("#2c2c2c")
	cPrimary, _    = ParseHexColor("#D03D56")
	cError, _      = ParseHexColor("#CF6679")

	o string
	y int32

	mainMenuCooldown Cooldown
	searchCooldown   Cooldown
	addCooldown      Cooldown
	delCooldown      Cooldown
	modCooldown      Cooldown

	alphabeth      = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	alphabethSlice = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	searchInput Input
	delInput    Input
	
	idAddInput      Input
	nameAddInput    Input
	lnameAddInput   Input
	ageAddInput     Input
	gradeAddInput   Input
	citizenAddInput Input

	oldIdModInput	Input
	idModInput      Input
	nameModInput    Input
	lnameModInput   Input
	ageModInput     Input
	gradeModInput   Input
	citizenModInput Input

	errorSlice		[]string
	errorLooped 	bool
	errorText		string
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

type Cooldown struct {
	Pressed bool
	Loops   int
	OnMenu  bool
}

type Input struct {
	InputText     []int
	OnInputBox    bool
	FramesCounter int
	LetterCount   int
	Valid         bool
}
