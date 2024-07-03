package main

import (
	"fmt"
)

//Loops
// Laçõs de repetições

func for_loops() {
	// sum := 0

	// * for *
	// for i := 0; i < 10; i++ {
	// 	fmt.Println("index: ", i)
	// 	sum += i
	// 	fmt.Println("soma: ", sum)
	// }

	// * for infinito *
	// for {
	// 	fmt.Println("loop infinito")
	// 	time.Sleep(2 * time.Second)
	// }

	// * for range *
	frutas := []string{"maca", "banana", "laranja"}
	for i, fruta := range frutas {
		fmt.Println(i, fruta)
	}

}
