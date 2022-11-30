package main

import (
	"fmt"
	"image/color"
	"math"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

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

func bSortInt(m map[int]int) MapMod {
	inv := make(map[int]int, len(m))
	s1, s2 := make([]int, len(m)), make([]int, len(m))
	sKeys, i := make([]int, len(m)), 0

	for k, v := range m {
		inv[v] = k
	}

	for k := range inv {
		sKeys[i] = k
		i++
	}

	bSorted := bubbleSort(sKeys)

	for _, i := range bSorted {
		s1, s2 = append(s1, inv[i]), append(s2, i)
	}

	return MapMod{s1, s2}
}

/*
Returns a MapMod with the key = all the ids && value = index number of the list
where the switch from false to true happens.
*/
func bSortBool(m map[int]bool) MapMod {
	l, l1, l2 := make([]int, 0), make([]int, 0), make([]int, 0)
	mm := MapMod{}

	for k, v := range m {
		if !v {
			l1 = append(l1, k)
		}
	}

	divIndx := []int{len(l1)}

	for k, v := range m {
		if v {
			l2 = append(l2, k)
		}
	}

	l1 = bubbleSort(l1)
	l2 = bubbleSort(l2)
	l = append(l, l1...)
	l = append(l, l2...)

	mm.key = l
	mm.val = divIndx

	return mm
}

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
	if rl.IsMouseButtonPressed(0) &&
		(posx < rl.GetMouseX()) &&
		(rl.GetMouseX() < width) &&
		(posy < rl.GetMouseY()) &&
		(rl.GetMouseY() < height) {
		return true
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

func intSliceToStr(s []int) string {
	r := ""
	for _, v := range s {
		d := strconv.FormatInt(int64(v), 10)
		r = r + d
	}
	return r
}

func intUnixSliceToInt(s []int) int {
	str := ""
	for _, v := range s {
		v = unixToInt(v)
		w := strconv.FormatInt(int64(v), 10)
		str = str + w
	}
	i, _ := strconv.Atoi(str)
	return i
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

func search(searchFor int) {
	index := binarySearch(searchFor, strToIntSlice(studentIdSlice))
	indexVal := studentIdSlice[index]
	if studentIdSlice[0] != indexVal {
		studentIdSlice = studentIdSlice[:len(studentIdSlice)-1]
		studentIdSlice = append(studentIdSlice, "")
		copy(studentIdSlice[0+1:], studentIdSlice[0:])
		studentIdSlice[0] = indexVal
	}
}