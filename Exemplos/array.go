package main

import (
	"fmt"
)

func variadic_soma(numeros ...int) (total int) {
	// Os números recebidos se transformar em um array de inteiros
	fmt.Println("Numeros: ", numeros)

	for _, numero := range numeros {
		total += numero
	}
	
	// O total da soma
	fmt.Println("Total: ", total)
	
	return
}

func main() {
	var numeros = []int{1, 1, 1, 1}
	
	variadic_soma(numeros...)
}

