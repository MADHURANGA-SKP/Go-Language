package main

import (
	"fmt"
	"slices"
)

func main(){
	var s []string
	fmt.Println("uninit", s, s ==nil, len(s)==0)

	s= make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))
	//cap()takes an input and returns its capacity depending on the input type.


	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c,s)//add s element in to c slice
	//copies elements into a destination slice dst from a source slice src
	fmt.Println("cpy:",c)

	l := s[2:5]// elements between index 2 and 5
	// this gets a slice of the elements
	fmt.Println("sl1:",l)

	l = s[:5]// elemnts till 5th index
	fmt.Println("sl2:",l)
	//This slices up to (but excluding)

	l = s[2:]//elements after index 2
	fmt.Println("sl3:",l)
	//This slices up from (but excluding)

	t := []string{"g","h","i"}
	// initialize a variable for slice in a single line
	fmt.Println("dcl:",t)

	t2 := []string{"g","h","i"}
	if slices.Equal(t, t2){
		fmt.Println("t==t2")
}

	twoD := make([][]int,3)
	for i := 0; i< 3; i++{
		innerLen := i +1
		twoD[i] = make([]int, innerLen)
		for j := 0; j< innerLen; j++{
			twoD[i][j] = i +j
		}
	}
	fmt.Println("2d: ",twoD)

}