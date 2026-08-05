package main

import "fmt"

func incrementValue(number *int) {
	*number++
}

func mutateFirst(mySlice []int) {
	mySlice[0] = 999
}

func mutateSecond(mySlice []int) {
	mySlice[1] = 1995
}

func push(s []int) []int {
	s = append(s, 42)
	return s
}

func main() {
	x := 10
	incrementValue(&x)
	fmt.Println("x: ", x)

	a := []int{1, 2, 3}
	mutateFirst(a)
	fmt.Println("a: ", a)

	b := push(a)
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	mutateSecond(b)
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
}
