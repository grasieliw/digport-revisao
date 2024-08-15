package main

import "fmt"

//Repita a lógica de comparação do exercício anterior só que dessa vez para uma lista de números.
//Nesse exercício não iremos retornar uma string, em vez disso, caso o número for maior que zero some 10
//e printe o resultado; caso for igual, some 2 e printe o resultado;
//caso for menor, some 23 e printe o resultado

func main() {

	//declarar lista
	lista := []int{1, 2, 3, 4, 5}

	// criar variavel de retorno e adicionar os parametros
	resultadoDisplay := TesteValor(lista)

	fmt.Println(resultadoDisplay)
}

func TesteValor(n []int) []int {

	//le os valores do slice
	//compara os valores lidos
	//calcula o resultado
	//retorna o resultado

	for i, num := range n {
		if num > 0 {
			n[i] = num + 10
		} else if num < 0 {
			n[i] = num + 23
		} else {
			n[i] = num + 2
		}
	}
	return n
}
