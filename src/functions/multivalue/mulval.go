package main

import "fmt"

func calculate(a int, b int)(int, int){
	var add int
	var mull int
	add = a+b
	mull = a*b
	return add,mull
}

func main(){
	var add,mull = calculate(3,5)
	fmt.Printf("addition : %d \n mulltiplication : %d", add, mull)
}