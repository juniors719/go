package main

import "fmt"

// structs - estruturas
// structs sao tipos de dados que permitem armazenar diferentes tipos de dados
// em um mesmo local
// não são classes, não tem herança, não tem polimorfismo

type Cliente struct {
	nome     string
	idade    int
	ativo    bool
	Endereco // composição
}

type Endereco struct {
	logradouro string
	numero     int
	cidade     string
	estado     string
}

func main() {

	wesley := Cliente{
		nome:  "Wesley",
		idade: 30,
		ativo: true,
		Endereco: Endereco{
			logradouro: "Rua dos Bobos",
			numero:     55,
			cidade:     "São Paulo",
			estado:     "SP",
		},
	}

	wesley.Endereco.cidade = "São Paulo"

	fmt.Println(wesley)
}
