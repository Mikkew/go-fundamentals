package main

import "fmt"

func increment(number *int) {
	*number++
}

type User struct {
	Name string
	Age  int
}

func birthday(user *User) {
	user.Age++
}

func swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	x := 10
	increment(&x)
	fmt.Println("x: ", x)

	user := User{
		Name: "Ana",
		Age:  20,
	}
	birthday(&user)
	fmt.Println("User: ", user)

	a, b := 1, 2
	swap(&a, &b)
	fmt.Println("a: ", a, "b: ", b)
}
