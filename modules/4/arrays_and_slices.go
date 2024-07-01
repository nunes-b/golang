package main

import (
	"fmt"
)

func arrays_and_slices() {
	var array [2]string
	array[0] = "Olá Mundo"
	array[1] = "Golang Is Awesome"
	// fmt.Println(array[0])
	// fmt.Println(array)

	// numerosPrimos := [6]int{2, 3, 5, 7, 11, 13}
	// fmt.Println(numerosPrimos)

	slice := make([]string, 2)
	slice[0] = "Rômulo"
	slice[1] = "Nunes"

	slice = append(slice, "Gabis")
	fmt.Println(slice)

	// Arrays e Slices são homogenios [ tem o mesmo tipo]
	// [1, 2, 3, 4] - []int
	// [1.1, 2.2, 3.3, 4.4] - []float64
	// ["Rumo", "para", "go"] - []string

	// 1 - Arrays = Tamanho fixo, de zero a N elementos do mesmo tipo.
	// Acessa valores com indices arr[1]
	// por conta do tamanho não é tão usado.
	// Tem a func len

	// 2 - Slices = Tamanho dinamico, de zero a N elementos do mesmo tipo.
	// Acessa valores com indices arr[1]
	// Tem a func len e append.

	//Maps : hetergenios, podem ter tipos diferentes
	// estrutura chave - valor
	// [key] = value
	// Chave tem um tipo e o valor pode ter outro.
	// map[string]int
	// {"Rômulo": 10, "Gabs": 20}
	// map[string]string
	// {"rômulo": "nunes", "gabs": "calixto"}

}
