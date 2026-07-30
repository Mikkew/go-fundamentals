package main

import "fmt"

func mapIterator() {
	myMap := map[string]int{
		"a": 1,
		"c": 3,
		"b": 2,
	}

	for i := 0; i < 3; i++ {
		fmt.Println("Loop: ", i)
		for key, value := range myMap {
			fmt.Println(key, value)
		}
	}

}
