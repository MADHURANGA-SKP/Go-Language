package main

import (
	"fmt"
	"os"
)

type Admin struct {
    ItemName string 
	ItemQty int
	Price int
}

var items  []Admin 

func ReadItem(items []Admin){
 if items == nil {
	fmt.Printf("Inventory is empty...!")
 }else{
	for _, items := range items { 
		fmt.Printf("Item Name: %s\nItem Quantity: %d\nItem Price: %d\n", items.ItemName,items.ItemQty,items.Price)
		fmt.Println("-------------------------------")	
	}
 }
}

func ShowItemById(id int, items []Admin){
	i := id-1
	if items != nil {
		for _, itemlist := range items { 
			if items[i] == itemlist {
			 fmt.Printf("Item Name: %s\n Item Quantity: %d\n Item Price: %d\n", itemlist.ItemName,itemlist.Price,itemlist.ItemQty)
			}
		 }
	}else{
		fmt.Printf("Inventory is empty...!\n\n")
	}
}

func CreateItem(items *[]Admin){
	var itemN string
	var itemQ int
	var itemP int
	fmt.Println()

	fmt.Print("Enter Item Name : ")
	fmt.Scan(&itemN)
	fmt.Print("Enter Item Quantitiy : ")
	fmt.Scan(&itemQ)
	fmt.Print("Enter Item Price : ")
	fmt.Scan(&itemP)

	Item_ := Admin{ItemName:itemN, ItemQty:itemQ, Price:itemP}

	*items = append(*items, Item_) 
	fmt.Println("Item successfully Added to the Inventory..!")
	fmt.Println()
}

func DeleteItemById(id int, items *[]Admin){
	i := id -1
	var deletedItems []Admin
   if i >=0 && i< len(*items) {
	deletedItems = append(deletedItems, (*items)[i])
	
	*items = append((*items)[:i],(*items)[i+1:]...)
		fmt.Println("------------ Deleted Item ------------")
		for _, item := range deletedItems {
			fmt.Printf("Item Name: %s\n Item Quantity: %d\n Item Price: %d\n", item.ItemName, item.ItemQty, item.Price)
		}
   } else {
	  fmt.Printf("The given Item Number is not in the Inventory list.\n")
   }

}

func UpdateItem(id int, items []Admin){
	i := id -1
	var itemN string
	var itemQ int
	var itemP int
	if items != nil {
		fmt.Println()

		fmt.Print("Enter Item Name : ")
		fmt.Scan(&itemN)
		fmt.Print("Enter Item Quantitiy : ")
		fmt.Scan(&itemQ)
		fmt.Print("Enter Item Price : ")
		fmt.Scan(&itemP)

		for _, itemlist := range items {
			if items[i] == itemlist {
				items[i] = Admin{ItemName:itemN,ItemQty:itemQ,Price:itemP}
				fmt.Println("Item successfully Updated..!")
			}
		}
	}else{
		fmt.Printf("Inventory is empty...!\n")
	 }
}

func AdminSession() {
	fmt.Println("--------Admin Inventory---------")
	fmt.Println("1. Show List of Items in the Inventory")
	fmt.Println("2. Show Items by Item Number")
	fmt.Println("3. Add a New Item")
	fmt.Println("4. Update a Item ")
	fmt.Println("5. Delete a Item")
	fmt.Println("6. Main Menu")
	fmt.Println("0. Exit")
	fmt.Print("Enter your option number: ")
}

func Reload(){
	var input string
	fmt.Print("Type 'Enter' go back to Admin menu: ")
	fmt.Scan(&input)
	if input == "Enter"{
		AdminMenu()
	}else{
		fmt.Printf("Invalid Input...!\n")
		Reload()
	}
}

func AdminMenu() {
	var option int
	var id int
	AdminSession()
	for {
		fmt.Scan(&option)
	if option == 0 {
		fmt.Println("BYE BYE.!")
		os.Exit(0)
	}else{
		switch option{
		case 1:
			ReadItem(items)
			fmt.Println()
			AdminSession()
		case 2:
			fmt.Println()
			fmt.Print("Enter Item Number that you want to get: ")
			fmt.Scan(&id)
			ShowItemById(id, items)
			Reload()
		case 3:
			CreateItem(&items)
			Reload()
		case 4:
			fmt.Println()
			fmt.Print("Enter Item Number that you want to Update: ")
			fmt.Scan(&id)
			UpdateItem(id, items)
			fmt.Println()
			Reload()
		case 5:
			fmt.Println()
			fmt.Print("Enter Item Number that you want to Delete: ")
			fmt.Scan(&id)
			DeleteItemById(id, &items)
			fmt.Println()
			Reload()
		case 6:
			fmt.Println()
			main()
		default:
			fmt.Printf("Re-enter your choice!\n")
			AdminMenu()
		}
	}
}
}