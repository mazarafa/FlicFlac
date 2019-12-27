package main

import (
	"fmt"
)

func main() {
	var x int 
	var y int 

	fmt.Print("Informe o numero:")
	fmt.Scanf("%d\n", &x)
<<<<<<< HEAD

	fmt.Print("Informe o numero:")
	fmt.Scanf("%d\n", &y)
=======
	
	fmt.Print("Informe o numero:")
    	fmt.Scanf("%d\n", &y)
	
>>>>>>> 85be239e22799d25d37dd943d2bc580f7ddb05b7
	
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
