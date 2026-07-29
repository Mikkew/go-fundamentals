package main

import (
	"fmt"
	"slices"
)

func slicesExample() {
	// Slice Literal
	// var mySlice = []int{1, 2, 3}
	// fmt.Println(mySlice)

	// mySlice[1] = 23
	// fmt.Println(mySlice)
	// fmt.Println(mySlice[1])

	// Slice vacío: Carece de valor
	// var zeroSlice []int
	// var otherZeroslice []int
	//nil = null
	// fmt.Println(zeroSlice == otherZeroslice)

	x := []int{1, 2, 3, 4, 5}
	y := []int{1, 2, 3, 4, 5}
	z := []int{1, 2, 3, 4, 5, 6}
	//s := []string{"a", "b", "c"}

	fmt.Println(slices.Equal(x, y))
	fmt.Println(slices.Equal(x, z))
	//fmt.Println(slices.Equal(x, s)) //NO Compila
}
