package main

import "fmt"

func main(){
	// local variable definition
	var i, j int

	for i =2; i<100; i++{
		for j=2; j<100; j++ {
			if(i%j==0){
				break;
			}
		}
		if(j>(i/j)){
			fmt.Printf("%d is a prime number\n", i)
		}
	}

}