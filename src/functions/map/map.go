package main

import (
	"fmt"
	"maps"
)

func main(){
	 m := make(map[string]int)
	 //create an empty map, use the builtin make

	 //Set key/value pairs using typical name[key] = val syntax
	 m["k1"] = 7
	 m["k2"] = 13

	 fmt.Println("map:", m)

	 v1 := m["k1"]//Get a value for a key with name[key]
	 fmt.Println("v3:",v1)

	 v3 := m["k3"]//Get a value for a key with name[key]
	 fmt.Println("v3:",v3)

	 fmt.Println("len:", len(m))//print the length of arrat m

	 delete(m, "k2")//delete selected key value from the  array
	 //delete removes key/value pairs from a map.
	 fmt.Println("map:",m)

	 clear(m)
	 //remove all key/value pairs from a map
	 fmt.Println("map:",m)

	 //blank identifier in Go. We can use 
	 //the blank identifier to declare and use the unused variables.
	 _, prs := m["k2"]
	 fmt.Println("prs:",prs)
	//when map inclues '' like zero values blank identfier can use
	//to ignrore them fro mthe array

	//declare and initialize a new map
	 n := map[string]int{"foo": 1, "bar": 2}
	 fmt.Println("map:", n)

	 n2 := map[string]int{"foo": 1, "bar":2}
	 if maps.Equal(n,n2){
		fmt.Println("n == n2")
	 }
}