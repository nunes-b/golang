package main

import "fmt"

func variables() {
	var name string

	name = "Rômulo"
	fmt.Println(name)

	var idade int
	idade = 42
	fmt.Println(idade)

	var b, c int = 1, 2
	fmt.Println(b)
	fmt.Println(c)

	fruit := "apple"
	fmt.Println(fruit)

	const idadeRomulo = 10
	fmt.Println(idadeRomulo)
}
