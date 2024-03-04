package main

import "fmt"

func add(a int, b int) int {
	var c int
	c = a + b
	return c 
}

func main(){
	var n = add(3,5)
	fmt.Printf("the sum is : %d", n)

}