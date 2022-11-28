package main

import (
	"math/rand"
	"time"
)

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

	for x := 0; x < RAND; x++ {
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
