package main

import (
	"fmt"
)

func main() {
	var num, horas int
	var valorHora float64
	fmt.Scan(&num, &horas, &valorHora)
	fmt.Printf("NUMBER = %d\n", num)
	fmt.Printf("SALARY = U$ %.2f\n", float64(horas)*valorHora)
}
