package main

// inferencia de tipos
// escopo global
var (
	b bool = true
	c int
	d string
	e float64
)

func main() {

	a := "X"  // shortcut - usado apenas na primeira vez que a variavel é criada
	a := "fd" // não funciona porque ela já existe

	var
	//b = 20 não posso mudar o tipo de uma variavel (fortemente tipada)

	// var a string - escopo local
	println (b)
}
