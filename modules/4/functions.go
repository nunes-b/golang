package main

import "fmt"

func functions() {
	var a int
	var b int
	imprime("digite dois valores para somar.:")
	fmt.Scan(&a)
	fmt.Scan(&b)
	somaDosValores := soma(a, b)
	fmt.Println("o resultado é:", somaDosValores)
}

// funções que começam com letras minusculas só podem ser
// usadas no proprio pacote, ela é privada.
// Maiusculas é publica, pode ser usada em qualquer parte do projeto
// Precisa tipar o retorno, e os parametros que ela recebe.
func soma(a int, b int) int {
	return a + b
}

func imprime(mensagem string) string {
	fmt.Println(mensagem)
	return mensagem
}
