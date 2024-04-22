package main

import (
	"errors"
	"fmt"
)

func main() {

	valor, err := sum(50, 20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(valor)
}

func sum(a, b int) (int, error) { // == func sum(a int, b int) int {
	if a+b >= 50 {
		return 0, errors.New("O valor da soma ultrapassa 50")
	}
	return a + b, nil // nil == null
}
