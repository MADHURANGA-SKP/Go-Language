package main

import "fmt"

func main() {
	var a int = 10
LOOP:
	for a < 20 {
		if a == 15 {
			/*skipp the iteration*/
			a = a + 1
			goto LOOP
		}
		fmt.Printf("value of a: %d\n", a)
		a++
	}
}