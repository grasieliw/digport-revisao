package main

import "fmt"

func main() {

	//criar valor
	//criar string que recebe + adicionar o valor criado como parametro
	//imprimir resultado

	Valor := 2

	resultado := TestaValor(Valor)

	fmt.Printf("%s", resultado)

}

func TestaValor(ValorDeTeste int) string {

	if ValorDeTeste > 0 {
		resultadoTeste := "greater than zero"
		return resultadoTeste
	} else if ValorDeTeste < 0 {
		resultadoTeste := "less than zero"
		return resultadoTeste
	} else {
		resultadoTeste := "zero"
		return resultadoTeste
	}

}
