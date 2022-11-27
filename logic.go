package main

import "fmt"

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

func printMapMod(m MapMod) {
	for a, x := range m.key {
		for b, y := range m.val {
			if a == b && (x != 0 || y != 0) {
				fmt.Println("Key: ", x, "\tValue: ", y)
			}
		}
	}
	fmt.Println("")
}

func bSortInt(m map[int]int) MapMod {
	inv := make(map[int]int, len(m))
	s1, s2 := make([]int, len(m)), make([]int, len(m))
	
	for k, v := range m { 
		inv[v] = k
	}
	fmt.Println(len(m),m, inv)


	sKeys, i := make([]int, len(inv)), 0
	for k := range inv {
		sKeys[i] = k
		i++
	}
	bSorted := bubbleSort(sKeys)

	for _, i := range bSorted {
		s1, s2 = append(s1, inv[i]), append(s2, i)
	}
	mapMod := MapMod{s1, s2}
	
	return mapMod
}

func boolMapToIntMap(m map[int]bool) map[int]int {
	intMap := make(map[int]int, len(m))
	for k, v := range m {
		if v {
			intMap[k] = 1
		} else if !v {
			intMap[k] = 0
		}
	}
	return intMap
}