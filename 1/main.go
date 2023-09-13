package main

import "fmt"

const a = "Hello, World"

type ID int

// inferencia de tipos
// escopo global
var (
	b bool    = true
	c int     = 10
	d string  = "Pedro"
	e float64 = 1.2
	f ID      = 1
)

func main() {

	fmt.Printf("O tipo de E é %T", f)
}
