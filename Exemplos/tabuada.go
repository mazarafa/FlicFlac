package main

import (
    "fmt"
)

func main() {
    var x int 
    var total int
    
    fmt.Print("Informe o numero:")
    fmt.Scanf("%d\n", &x)
    fmt.Print("\n")

    fmt.Println("=== tabuada ===")
    for i:=0; i < 10; i++{
        total = x*i
        
        fmt.Println(x, "x", i,"=", total)
    }
    
}