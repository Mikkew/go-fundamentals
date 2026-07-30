package main

import "fmt"

func forComplete() {
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// i := 0
	// for ; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	for i := 0; i < 10; {
		fmt.Println(i)
		if i%2 == 0 {
			i++
		} else {
			i += 2
		}
	}

	// for i := range 10 {
	// 	fmt.Println(i)
	// }
}
