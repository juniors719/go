package main

import (
	"fmt"
)

const a = "Hello, World!"

type ID int

var (
	b bool    = true
	c int     = 10
	d string  = "Junior"
	e float64 = 1.5
	f ID      = 1
)

func main() {

	// Declarando um array de 3 posições, do tipo int
	var meuArray [3]int
	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 30

	// %T é o tipo da variável
	// %v é o valor da variável
	fmt.Println(meuArray[len(meuArray)-1])

	for i, v := range meuArray {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}
}
