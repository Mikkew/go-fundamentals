package main

import (
	"fmt"
	"maps"
)

func maicompareMaps() {
	myMapA := map[string]int{
		"hello": 5,
		"world": 10,
	}

	myMapB := map[string]int{
		"hello": 5,
		"world": 10,
	}

	fmt.Println(maps.Equal(myMapA, myMapB))
}
