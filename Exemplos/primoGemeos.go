package main

import (
	"fmt"
)

func main() {
	var x int 
	var y int 

	fmt.Print("Informe o numero:")
	fmt.Scanf("%d\n", &x)
	
	fmt.Print("Informe o numero:")
    fmt.Scanf("%d\n", &y)
	
	
	var primo bool

	for i := x; i < y+1; i++ {
		primo = true

		for j := 2; j < i; j++ {

			if i%j == 0 {
				primo = false
				break
			}
		}

		if primo {
			fmt.Println(i)
		}
	}
}
