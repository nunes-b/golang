package main

import "fmt"

func if_else() {
	valor := 13

	if valor == 1 {
		fmt.Printf("O valor é: %v\n", valor)
	} else {
		fmt.Printf("O valor é diferente de 1, o valor é: %v\n", valor)
	}

	if 7%2 == 0 {
		fmt.Println("7 e par")
	} else {
		resto := 7 % 2
		fmt.Println(resto)
		fmt.Println("7 e impar")
	}
}
