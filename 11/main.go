package main

import "fmt"

// structs - estruturas
// structs sao tipos de dados que permitem armazenar diferentes tipos de dados
// em um mesmo local
// não são classes, não tem herança, não tem polimorfismo

type Cliente struct {
	nome  string
	idade int
	ativo bool
}

func main() {

	wesley := Cliente{
		nome:  "Wesley",
		idade: 30,
		ativo: true,
	}

	fmt.Printf("Nome: %s\n Idade: %d\n Ativo: %t\n", wesley.nome, wesley.idade, wesley.ativo)

	wesley.ativo = false

	fmt.Printf("Nome: %s\n Idade: %d\n Ativo: %t\n", wesley.nome, wesley.idade, wesley.ativo)
}
