package main

import (
	"fmt"
)

func tabuada(numeros ...int) (total int) {
	//soma = 0
	//x = 10
	for _, numero := range numeros {
		//soma = divide * x
		fmt.Println(divide)
		fmt.Println(numeros * total)
	}
	
	// O total da soma
	//fmt.Println(soma)
	
	return
}

func main() {
	tabuada(2,10)
	// variadic_soma(1, 1)
	// variadic_soma(1, 4, 5, 6)
}

