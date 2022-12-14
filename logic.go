package main

import (
	"fmt"
	"image/color"
	"math"
	"strconv"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// START OF SORTING
/*
Recives value to search for and a sorted array of ints to search in.
Returns the index value of the desired value.
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
	// fmt.Println(v, s[indMin], s)

	if v == s[indMin] {
		return indMin
	} else {
		return -1
	}

}

func bubbleSort(s []int) []int {
	for y := 0; y < len(s)-1; y++ {
		for v := 0; v < len(s)-y-1; v++ {
			if s[v] > s[v+1] {
				s[v], s[v+1] = s[v+1], s[v]
			}
		}
	}
	return s
}

func resortString(oldinxex, newindex, sort []string) []string {
	r := make([]string, 0)
	for _, v := range newindex {
		for i2, v2 := range oldinxex {
			if v == v2 {
				r = append(r, sort[i2])
			}
		}

	}

	return r
}

// Probably the most stupid an inefficient func ever created
// buy it works so im happy... kind of
func bSortString(m map[int]string) MapMod[string] {
	// fmt.Println(m, len(m))
	text := make([]string, 0)
	ids := make([]int, 0)
	stringsInt, idsStrSlice := make([][]string, 0), make([]string, 0)

	for k, val := range m {
		ids = append(ids, k)
		text = append(text, val)
		eachStr := make([]string, 0)
		for _, char := range val {
			str := ""
			for i, r := range alphabeth {
				charIntVal := ""
				if string(char) == strings.ToLower(string(r)) {
					charIntVal = strconv.FormatInt(int64(i), 10)
					str = str + charIntVal
				}
			}
			eachStr = append(eachStr, str)
		}
		stringsInt = append(stringsInt, eachStr)
	}

	firstChars := make([]string, 0)
	for _, v := range stringsInt {
		firstChar := v[0]
		firstChars = append(firstChars, firstChar)
	}

	sorted := intToStrSlice(bubbleSort(strToIntSlice(firstChars)))

	for i := range sorted {
		for index, v := range firstChars {
			id := ids[index]
			if sorted[i] == v {
				idsStrSlice = append(idsStrSlice, strconv.FormatInt(int64(id), 10))
			}
		}
	}

	r := make([]string, 0)
	for _, v := range idsStrSlice {
		s, _ := strconv.Atoi(v)
		r = append(r, m[s])
	}


	// map1 := make(map[int]string, 0)
	// for i1, id := range idsStrSlice {
	// 	intId, _ := strconv.Atoi(id)
	// 	for i2, s := range r {
	// 		if i1 == i2 {
	// 			map1[intId] = s
	// 		}
	// 	}
	// }
	// for a, b := range map1 {
	// 	for _, v := range idsStrSlice {
	// 		v1, _ := strconv.Atoi(v)
	// 		if a == v1 {
	// 			fmt.Println("key: ", a, " val: ", b)
	// 		}
	// 	}
	// }

	return MapMod[string]{r, idsStrSlice}
}

// Transforms a list of strings that are numbers into a string
// based on the abc index.
func intToStrAsPerAbcID(s int) string {
	r := ""
	for i, run := range alphabeth {
		if s == i {
			r = r + strings.ToLower(string(run))
		}
	}
	return r
}

func bSortInt(m map[int]int) MapMod[int] {
	inv := make(map[int]int, len(m))
	s1 := make([]int, 0)
	sVals, sKeys, i := make([]int, len(m)), make([]int, len(m)), 0

	for k, v := range m {
		inv[k] = v
	}

	for k, v := range inv {
		sKeys[i] = k
		sVals[i] = v
		i++
	}

	bSorted := bubbleSort(sVals)

	for _, v := range bSorted {
		for i, v2 := range m {
			if v == v2 {
				s1 = append(s1, i)
				break
			}
		}
	}

	return MapMod[int]{bSorted, s1}
}

/*
Returns a MapMod with the key = all the ids && value = index number of the list
where the switch from false to true happens.
*/
func bSortBool(m map[int]bool) MapMod[int] {
	l, l1, l2 := make([]int, 0), make([]int, 0), make([]int, 0)
	mm := MapMod[int]{}

	for k, v := range m {
		if !v {
			l1 = append(l1, k)

		} else if v {
			l2 = append(l2, k)
		}
	}

	l1 = bubbleSort(l1)
	l2 = bubbleSort(l2)
	l = append(l, l1...)
	l = append(l, l2...)

	divIndx := []int{len(l1)}
	values, val := make([]int, 0), 0
	for i, _ := range l {
		if i == divIndx[0] {
			val = 1
		}
		values = append(values, val)
	}

	mm.key = l
	mm.val = values

	return mm
}

func binaryToBool(s []int) []bool {
	r := make([]bool, 0)
	for _, v := range s {
		if v == 0 {
			r = append(r, false)
		} else if v == 1 {
			r = append(r, true)
		}
	}
	return r
}

// END OF SORTING

// START OF RENDERS
func ParseHexColor(s string) (c color.RGBA, err error) {
	c.A = 0xff
	switch len(s) {
	case 7:
		_, err = fmt.Sscanf(s, "#%02x%02x%02x", &c.R, &c.G, &c.B)
	case 4:
		_, err = fmt.Sscanf(s, "#%1x%1x%1x", &c.R, &c.G, &c.B)
		// Double the hex digits:
		c.R *= 17
		c.G *= 17
		c.B *= 17
	default:
		err = fmt.Errorf("invalid length, must be 7 or 4")

	}
	return
}

func getCollsX() []int32 {
	tSlice := []int32{}
	for x := 0; x < 6; x++ {
		y := int32(quadFunc(float64(x), 2.083, -30, 129.67, 43.04, 71.79))
		tSlice = append(tSlice, y)
	}
	return tSlice
}

/*
Returns the quadratic of the input
*/
func quadFunc(x, a, b, c, d, e float64) float64 {
	z := (a*math.Pow(x, 4) + b*math.Pow(x, 3) + c*math.Pow(x, 2) + d*x + e)
	return z
}

func drawColl(s []string, x int32) []string {
	y = 200
	s = scroll(s)
	for _, v := range s {
		rl.DrawText(v, x, y, 30, rl.White)
		y = y + 58
	}
	return s
}

func scroll(s []string) []string {
	if rl.GetMouseWheelMove() > 0 {
		o, s = s[0], s[1:]
		s = append(s, o)
	} else if rl.GetMouseWheelMove() < 0 {
		o, s = s[len(s)-1], s[:len(s)-1]
		s = append(s, "")
		copy(s[0+1:], s[0:])
		s[0] = o
	}
	return s

}

func button(posx, posy, width, height int32) bool {
	if (posx < rl.GetMouseX()) &&
		(rl.GetMouseX() < width) &&
		(posy < rl.GetMouseY()) &&
		(rl.GetMouseY() < height) {
		return true
	}
	return false
}

func search(searchFor int, b bool) int {
	// How the absolute fuck can I mess up something that was working flawlessly
	// fmt.Println("")
	old := make([]string, 0)
	for _, v := range studentIdSlice {
		old = append(old, v)
	}
	searchInSorted := bubbleSort(strToIntSlice(studentIdSlice))
	// fmt.Println(studentAgeSlice)

	if b {
		studentIdSlice = intToStrSlice(searchInSorted)
		studentNameSlice = resortString(old, studentIdSlice, studentNameSlice)
		studentLastNameSlice = resortString(old, studentIdSlice, studentLastNameSlice)
		studentAgeSlice = resortString(old, studentIdSlice, studentAgeSlice)
		studentGradeSlice = resortString(old, studentIdSlice, studentGradeSlice)
		studentCitizenSlice = resortString(old, studentIdSlice, studentCitizenSlice)
	}
	// fmt.Println(studentAgeSlice)

	save()

	index := binarySearch(searchFor, searchInSorted)
	return index
}

func moveSliceToTop(index int, s []string) {
	indexVal := s[index]
	if s[0] != indexVal {
		s[index] = s[0]
		s[0] = indexVal
	}
}

func sliceStrDelete(index int, slice []string) []string {
	copy(slice[index:], slice[index+1:])
	slice[len(slice)-1] = ""
	slice = slice[:len(slice)-1]
	return slice
}

func resetInput(i Input) Input {
	i.LetterCount = 0
	i.InputText = []int{}
	return i
}

// END OF RENDERS

// START OF SLICE MANIPULATION
func intSliceToBool(s []int) []bool {
	r := make([]bool, 0)
	for _, v := range s {
		if v == 0 {
			r = append(r, false)
		} else if v == 1 {
			r = append(r, true)
		}
	}
	return r
}

func boolSliceToString(s []bool) []string {
	r := make([]string, 0)
	for _, v := range s {
		r = append(r, strconv.FormatBool(v))
	}
	return r
}

func sliceContains(k string, s []string) bool {
	for _, v := range s {
		if k == v {
			return true
		}
	}
	return false
}

func strToIntSlice(s []string) []int {
	r := make([]int, 0)
	for _, v := range s {
		d, _ := strconv.Atoi(v)
		r = append(r, d)
	}
	return r
}

func intToStrSlice(s []int) []string {
	r := make([]string, 0)
	for _, v := range s {
		d := strconv.FormatInt(int64(v), 10)
		r = append(r, d)
	}
	return r
}

func intSliceToStr(s []int) string {
	r := ""
	for _, v := range s {
		d := strconv.FormatInt(int64(v), 10)
		r = r + d
	}
	return r
}

func unixSliceToStr(s []int) string {
	str := ""
	for _, v := range s {
		if inBetween(v, 48, 58) {
			v = unixToInt(v)
			w := strconv.FormatInt(int64(v), 10)
			str = str + w
		} else if inBetween(v, 65, 90) {
			o = unixToStr(v)
			str = str + o
		}
	}
	return str

}

func unixToInt(s int) int {
	for inBetween(s, 48, 58) {
		return s - 48
	}
	return -1
}

func unixToStr(s int) string {
	for inBetween(s, 65, 90) {
		for i := range alphabeth {
			if s-65 == i {
				return string(alphabeth[s-65])
			}
		}
	}
	return ""
}

func inBetween(x int, a, b int) bool {
	if a <= x && x <= b {
		return true
	} else {
		return false
	}
}

func isIntSlice(s []int) bool {
	for _, v := range s {
		if !inBetween(v, 48, 58) {
			return false
		}
	}
	return true
}

// END OF SLICE MANIPULATION

func mapContains[V int | string | bool](a string, s map[int]V) bool {
	i, _ := strconv.Atoi(a)
	for k, _ := range s {
		if i == k {
			return true
		}
	}
	return false
}

func updateMaps() {
	for _, v := range studentIdSlice {
		if !mapContains(v, studentNameMap) {
			s, _ := strconv.Atoi(v)
			studentNameMap[s] = studentNameSlice[len(studentIdSlice)-1]
		}
		if !mapContains(v, studentLastNameMap) {
			s, _ := strconv.Atoi(v)
			studentLastNameMap[s] = studentLastNameSlice[len(studentIdSlice)-1]
		}
		if !mapContains(v, studentAgeMap) {
			s, _ := strconv.Atoi(v)
			g, _ := strconv.Atoi(studentAgeSlice[len(studentIdSlice)-1])
			studentAgeMap[s] = g
		}
		if !mapContains(v, studentGradeMap) {
			s, _ := strconv.Atoi(v)
			g, _ := strconv.Atoi(studentGradeSlice[len(studentIdSlice)-1])
			studentGradeMap[s] = g
		}
		if !mapContains(v, studentCitizenMap) {
			s, _ := strconv.Atoi(v)
			g, _ := strconv.ParseBool(studentCitizenSlice[len(studentIdSlice)-1])
			studentCitizenMap[s] = g
		}
	}
}

func updateMainMap() {
	ss := make(map[int]Student, 0)
	for _, v := range studentIdSlice {
		id, _ := strconv.Atoi(v)
		y := Student{
			studentNameMap[id],
			studentLastNameMap[id],
			studentAgeMap[id],
			studentGradeMap[id],
			studentCitizenMap[id],
		}
		ss[id] = y
	}

	for k := range studentMap {
		delete(studentMap, k)
	}

	for k, v := range ss {
		s := strconv.FormatInt(int64(k), 10)
		studentMap[s] = v
	}
}
