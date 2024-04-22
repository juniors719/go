package main

import "fmt"

func main() {

	salarios := map[string]int{"wagner": 1000, "joao": 2000, "maria": 3000}

	fmt.Println(salarios["wagner"])

	delete(salarios, "wagner")

	fmt.Println(salarios)

	salarios["pedro"] = 1200

	sal := make(map[string]int)
	sal["junior"] = 10000
	sal["joao"] = 2000
	sal["maria"] = 3000
	fmt.Println(sal)

	// _ = ignorando o indice -> blank identifier
	for _, salario := range sal {
		fmt.Printf("Salario: %d\n", salario)
	}

}
