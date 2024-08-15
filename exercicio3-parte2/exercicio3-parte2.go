package main

import "fmt"

func main() {

	lista := []int{1, 2, 3, 4, 5}

	resultadoDisplay := TesteValor(lista)

	fmt.Println(resultadoDisplay)
}

func TesteValor(n []int) []int {

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
