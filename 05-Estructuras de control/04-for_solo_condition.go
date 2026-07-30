package main

import "fmt"

func forSoloCondition() {

	// For se solo condicion
	i := 0
	for i < 10 {
		fmt.Println(i)
		if i%2 == 0 {
			i++
		} else {
			i += 2
		}
	}
}
