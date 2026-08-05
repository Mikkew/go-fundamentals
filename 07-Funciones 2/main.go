package main

import "fmt"

func main() {
	// ok := PrintOK("Todo está en orden")
	// PrintOK("Todo cool")
	// PrintOK("Funcionó bien")

	// fmt.Println(ok)

	// demoDeferLIFO()

	fmt.Println(factorial(5))
	fmt.Println(sum([]int{1, 2, 3, 4}))
}

func PrintOK(msg string) string {
	fmt.Printf("OK - %s\n", msg)
	return msg
}

func demoDeferLIFO() {
	x := 10
	defer fmt.Println(x)
	x = 99
	defer PrintOK("Defer #1 ")
	defer PrintOK("Defer #2")
	PrintOK("Cuerpo ")
}

func factorial(number int) int {
	if number < 0 {
		panic("El numero debe ser mayor o igual a cero (0)")
	}

	if number == 0 {
		return 1
	}

	return number * factorial(number-1)
}

func sum(nums []int) int {
	// Caso base
	if len(nums) == 0 {
		return 0
	}

	return nums[0] + sum(nums[1:])
}
