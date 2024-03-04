package main

import "fmt"

func main(){
	x:= []int{10,20,30,40,50}

	for index, element :=range x {
		fmt.Println(index, " - ", element)
	}
}