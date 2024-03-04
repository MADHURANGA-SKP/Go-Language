package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to the Shoping Cart--------------------------")
	fmt.Println("Plese Select USER/ADMIN option from the given below :")
	fmt.Println("1. USER")
	fmt.Println("2. ADMIN")
	fmt.Println("0. EXIT")
	fmt.Print("Enter Option Number : ")
	var userInput int
	fmt.Scan(&userInput)
		if userInput == 0 {
			fmt.Println("BYE BYE.!")
			os.Exit(0)
		}else{
			switch userInput{
			case 1:
				ProductMenu()
			case 2:
				AdminMenu()
			default:
				fmt.Printf("Invalid Input...!\n")
				os.Exit(0)
			}
		}
	
	
}