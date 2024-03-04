package main

import "fmt"

func main() {
    a := [6]int{52, 2, 13, 35, 9, 8}
    for i,x:= range a {
        fmt.Printf("a[%d] = %d\n", i,x)
    }  
}