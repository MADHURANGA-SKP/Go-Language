package main

import "fmt"

func main(){
	var a [5]int
	fmt.Println("emp:",a)//show the array 1

	a[4] = 100//initilize 4 index as 100
	fmt.Println("set:",a)//show the array a
	fmt.Println("get:",a[4])//get 4th index element
	fmt.Println("len:",len(a))//get the length of the array a


	b := [5]int{1,2,3,4,5}//initialize an array called b
	fmt.Println("dcl",b)//show the array b

	var twoD[2][3]int//initialize 2d array called twoD
	for i := 0; i<2; i++ {//goes through the rows of twoD
		for j := 0; j<3; j++ { //goes through the cols of the twoD
			twoD[i][j] = i + j 
	} //intialize elements to the twoD by addition values of i and j
}
	fmt.Println("2d: ",twoD)
}