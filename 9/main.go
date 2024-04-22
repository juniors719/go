package main

// closures - funcoes anonimas

import (
	"fmt"
)

func main() {

	t := func() int {
		return sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10) * 2
	}()

	fmt.Println(t)
}

func sum(numeros ...int) int { // ... == variadic function - aceita um numero variavel de argumentos
	total := 0
	for _, numero := range numeros { // ignorando o indice -> blank identifier
		total += numero // incrementa o total com o valor do numero
	}
	return total // retorna o total
}
