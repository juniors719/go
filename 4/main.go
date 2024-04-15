package main

const a = "Valor Constante"

// escopo global
var (
	// declarando uma variável b do tipo booleana
	// por padrão, o valor inferido é false
	b bool

	// valor padrão = 0
	c int

	// valor padrão = ""
	d string

	// valor padrão = +0.000000e+000
	e float64

	f bool = true // declarando e atribuindo
)

func main() {
	a := "x" // short declaration
	println(a)
}
