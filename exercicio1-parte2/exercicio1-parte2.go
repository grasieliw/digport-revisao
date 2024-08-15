package main

import "fmt"

//variaveis

var NumeroTaboada int
var Incrementa int

func main() {

	NumeroTaboada := 4
	Incrementa := 1

	for i := 0; i < 10; i++ {
		ResultadoDisplay := taboada(NumeroTaboada, Incrementa)

		fmt.Printf("%d X %d = %d\n", Incrementa, NumeroTaboada, ResultadoDisplay)

		Incrementa++
	}
}

func taboada(NumeroInteiro, ValorIncrementa int) int {

	Resultado := ValorIncrementa * NumeroInteiro //1*4

	return Resultado
}
