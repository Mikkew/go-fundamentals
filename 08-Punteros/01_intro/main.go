package main

import "fmt"

func main() {
	x := 10
	p := &x // Acceder a la dirección de memoria de x

	fmt.Println("x = ", x)
	fmt.Println("p = ", p)
	fmt.Println("*p = ", *p)

	*p = 99
	fmt.Println("*p = ", *p)
	fmt.Println("x = ", x)

	x = 22
	fmt.Println("*p = ", *p)
	fmt.Println("x = ", x)
}
