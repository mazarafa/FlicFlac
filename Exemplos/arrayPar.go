package main

import (
    "fmt"
)
func main() {
    var x int 
	var count int
	var count2 int
    
    fmt.Print("Informe o numero:")
    fmt.Scanf("%d\n", &x)
    fmt.Print("\n")

    //x = x + 1 gambiarra linda haushaushu
    for i:=1; i <= x; i++{
		

		if( i% 2 == 0){
			count += 1
		}else{
			count2 +=1
		}
        
	}
	fmt.Println("Numeros pares:" , count)
	fmt.Println("Numeros impares:" , count2)
    
}