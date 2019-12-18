package main

import (
    "fmt"
)

func main() {
	//var valor int = -10
	//var primo bool = "true"
	var x int
	var y int  
	//var total int
	fmt.Print("Informe o numero:")
    fmt.Scanf("%d %d\n", &x, &y)
	
	for i:=x; i < y+1; i++{
		//primo = true
		for j:=2; j < i; j++{
			if (i%j == 0){
				//primo = false
				break
			};

		};
		// if(i-valor == 2){
		// 	fmt.Print(valor, i)
		// 	valor = i


		// }
	}
}


