package main

import "fmt"

func main(){
	var a int = 10

	for a < 20 {
		fmt.Printf("value of a %d\n",a)
		a++
		if a>15 {
break;
		}
	}
}