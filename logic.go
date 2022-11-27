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