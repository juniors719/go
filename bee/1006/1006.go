package main

import (
	"fmt"
)

func main() {
	var a, b, res float64
	fmt.Scan(&a, &b)
	res = ((a * 3.5) + (b * 7.5)) / 11
	fmt.Printf("MEDIA = %.5f\n", res)
}
