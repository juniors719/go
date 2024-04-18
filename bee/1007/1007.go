package main

import (
	"fmt"
)

func main() {
	var a, b, c, d, res int
	fmt.Scan(&a, &b, &c, &d)
	res = ((a * b) - (c * d))
	fmt.Printf("DIFERENCA = %d\n", res)
}
