package main

import (
	"fmt"
	"reflect"
)

type T struct {
	X int
	Y string
	Z []int
	M map[string]int
}

func main() {
	t1 := T{
		X: 1,
		Y: "lei",
		Z: []int{1, 2, 3},
		M: map[string]int{
			"a": 1,
			"b": 2,
		},
	}

	t2 := T{
		X: 1,
		Y: "lei",
		Z: []int{1, 2, 3},
		M: map[string]int{
			"a": 1,
			"b": 3, //changed this for demonstration
		},
	}

	fmt.Println(reflect.DeepEqual(t1, t2))

	a1 := []int{1, 2, 3, 4}
	a2 := []int{1, 2, 3, 4}
	fmt.Println(reflect.DeepEqual(a1, a2))

	m1 := map[string]int{
		"a": 1,
		"b": 2,
	}
	m2 := map[string]int{
		"a": 1,
		"b": 2,
	}
	fmt.Println(reflect.DeepEqual(m1, m2))
}
