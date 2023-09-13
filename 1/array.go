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

	var meuArray [3]int
	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 30

	fmt.Println(meuArray[0])

	//percorrendo array
	for i, v := range meuArray {
		fmt.Printf("O valor do indice é %d e o valor é %d\n", i, v)
	}
}
