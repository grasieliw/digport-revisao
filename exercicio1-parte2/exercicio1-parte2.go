package main

import "fmt"

//variaveis

var NumeroTaboada int
var Incrementa int
var Resultado int

func main() {

	NumeroTaboada := 4
	Incrementa := 1

	for i := 0; i < 10; i++ {
		ResultadoDisplay := taboada(NumeroTaboada, Incrementa)

		fmt.Printf("%d X %d = %d", Incrementa, NumeroTaboada, ResultadoDisplay)

		Incrementa++
	}
}

func taboada(NumeroInteiro, ValorIncrementa int) int {

	Resultado = Incrementa * NumeroInteiro //1*4

	return Resultado
}
