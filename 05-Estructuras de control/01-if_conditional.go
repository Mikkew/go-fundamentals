package main

import (
	"fmt"
	"math/rand"
)

func ifConditional() {
	// shouldBuy := true
	// debts := true

	// if shouldBuy {
	// 	if debts {
	// 		fmt.Println("Será para la próxima")
	// 	} else {
	// 		fmt.Println("Dale sin miedo")
	// 	}
	// } else {
	// 	fmt.Println("Será para la próxima")
	// }

	if n := rand.Intn(10); n == 0 {
		fmt.Println("Numero cero")
	} else if n > 5 {
		fmt.Println("Número mayor a 5: ", n)
	} else {
		fmt.Println("Número ideal: ", n)
	}
}
