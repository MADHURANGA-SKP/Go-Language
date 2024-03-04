package main

import "fmt"

func main(){
	var n int
	fmt.Println("please enter an number that you want to check.")

	fmt.Scanln(&n)
	if n % 2 == 0 {
		fmt.Println(n, "even number")
	}else {
		fmt.Println(n,"odd number")
	}
}