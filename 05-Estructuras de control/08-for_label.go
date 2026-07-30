package main

import "fmt"

func forLabel() {
	examples := []string{"Hola", "Carambola", "Bola", "Polola"}

outer:
	for _, example := range examples {
		for index, value := range example {
			fmt.Println(index, value, string(value))

			if value == 'l' {
				continue outer
			}
		}
	}
}
