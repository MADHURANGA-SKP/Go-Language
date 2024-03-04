package main

import (
	"fmt"
	"os"
	"strings"
)

type User struct {
    ItemName string 
	ItemQty int
	Price int
}

var Product  []User 

func ReadInventory(items []Admin){
 if items == nil {
	fmt.Printf("Inventory is empty...!\n")
	fmt.Printf("Unable to proceed. Redirecting to Main Menu..!\n\n")
	main()
	
 }else{
	for _, items := range items { 
		fmt.Printf("Item Name: %s\nItem Quantity: %d\nItem Price: %d\n", items.ItemName,items.ItemQty,items.Price)	
	}
	UserSession()
 }
}

func ReadProduct(Product []User){
	if Product == nil {
	   fmt.Printf("Your Cart is empty...!\n")
	}else{
	fmt.Printf("------------ My Cart ------------\n")
	   for _, Product := range Product { 
		   fmt.Printf("Item Name: %s\n Item Quantity: %d\n Item Price: %d\n", Product.ItemName,Product.ItemQty,Product.Price)	
	   }
	}
}


func CreateProduct(Product *[]User, items *[]Admin) {
	var itemN string
	var itemQ int
	

	fmt.Println()

	fmt.Print("Enter Item Name : ")
	fmt.Scan(&itemN)
	fmt.Print("Enter Item Qty : ")
	fmt.Scan(&itemQ)
	

	var found bool
	for i, item := range *items {
		if item.ItemName == itemN {
			found = true
			
			if item.ItemQty >= itemQ {
				*Product = append(*Product, User{ItemName: itemN, ItemQty: itemQ, Price: item.Price})
				
				(*items)[i].ItemQty -= itemQ 
				fmt.Printf("Item successfully added to cart..!\n")
			} else {
				fmt.Printf("Not enough items in the inventory.\n")
			}
			break
		}
	}

	if !found {
		fmt.Printf("Item not found in the inventory.\n")
	}

	fmt.Println()
}

func DeleteProductById(id int, Product *[]User){
	i := id -1
   if Product != nil {
		fmt.Println("Item number that need to be delted is:", id)
	if i >=0 && i< len(*Product) {
		*Product = append((*Product)[:i],(*Product)[i+1:]...)
		for _, Product := range *Product {
			fmt.Printf("Item Name: %s\nItem Quantity: %d\nItem Price: %d\n", Product.ItemName,Product.Price,Product.ItemQty)
			fmt.Printf("Item successfully Deleted..!\n")	
			}
	} else {
		fmt.Printf("The given Item number is not in the list.\n")
	}
   }else{
	fmt.Printf("Inventory is empty..!")
   }

}

func UpdateProduct(Product *[]User, items []Admin){
	var itemN string
	var itemQ int
	if Product != nil {
		fmt.Println()

		fmt.Print("Enter Item Name : ")
		fmt.Scan(&itemN)
		fmt.Print("Enter Item Qty : ")
		fmt.Scan(&itemQ)
		
		var found bool
		for i, userItem := range *Product {
			if userItem.ItemName == itemN {
				found = true
				
				for x, Inventory := range items {
					if Inventory.ItemName == itemN && Inventory.ItemQty >= itemQ {
						
						(*Product)[i].ItemQty = itemQ
						
						items[x].ItemQty -= itemQ
						fmt.Printf("Item successfully updated in the cart..!\n")
						break
					} else {
						fmt.Printf("Not enough items in the inventory or invalid quantity.\n")
					}
				}
				break
			}
		}

		if !found {
			fmt.Printf("Item not found in the cart.\n")
		}
	}else{
		fmt.Printf("Cart is empty...!\n")
	 }
}

func Total(id int, Admin []User){
	i := id -1
	var itemN string
	var itemQ int
	if Product != nil {
		fmt.Println()

		fmt.Print("Enter Item Name : ")
		fmt.Scan(&itemN)
		for _, itemlist := range Product {
			if Product[i] == itemlist {
				Product[i] = User{ItemName:itemN,ItemQty:itemQ}
				fmt.Printf("Item successfully Updated..!\n")
			}
		}
	}else{
		fmt.Printf("Cart is empty...!\n")
	 }
}

func CalTotal(Product []User) int {
	total := 0
	for _, item := range Product {
		total += item.Price * item.ItemQty
	}
	return total
}

func Invoice(Product []User){
	if Product == nil {
	   fmt.Printf("Your Cart is empty...!\n")
	}else{
		fmt.Println(strings.Repeat("-",20),"Invoice",strings.Repeat("-",20))
		fmt.Println(strings.Repeat("-",47))
		fmt.Printf("| %-20s | %-10s | %-10s |\n", "Item Name", "Quantity", "Price")
		fmt.Println("|----------------------|------------|------------|")

	   for _, Product := range Product { 
		   fmt.Printf("| %-20s | %-10d | %-10d |\n", Product.ItemName,Product.ItemQty,Product.Price,)	
	   }

	   fmt.Println(strings.Repeat("-",47))
	   fmt.Printf("Total Price: %d/=\n", CalTotal(Product))
	}
   }

func UserSession() {
	fmt.Println("-------------------------------")
	fmt.Printf("\n")
	fmt.Println("1. Before start the process Check Inventory for Products avalible or not..!")
	fmt.Println("2. Show My Cart ")
	fmt.Println("3. Add a New Item to My Cart from Inventory")
	fmt.Println("4. Update a Item on My Cart ")
	fmt.Println("5. Delete a Item My Cart")
	fmt.Println("6. Show Invoice")
	fmt.Println("7. Main Menu")
	fmt.Println("0. Exit")
	fmt.Print("Enter your Option By Number: ")
}


func LoadAgain(){
	var input string
	fmt.Print("Type 'Enter' go back to User Menu: ")
	fmt.Scan(&input)
	if input == "Enter"{
		ProductMenu()
	}else{
		fmt.Printf("Invalid Input...!\n")
		LoadAgain()
	}
}

func ProductMenu() {
	var option int
	var id int
	UserSession()
	for {
		fmt.Scan(&option)
	if option == 0 {
		fmt.Printf("BYE BYE.!\n")
		os.Exit(0)
	}else{
		switch option{
		case 1:
			fmt.Printf("-----------Inventory----------\n")
			ReadInventory(items)
		case 2:
			ReadProduct(Product)
			fmt.Println()
			UserSession()
		case 3:
			CreateProduct(&Product, &items)
			LoadAgain()
		case 4:
			fmt.Println()
			UpdateProduct(&Product, items)
			fmt.Println()
			LoadAgain()
		case 5:
			fmt.Println()
			fmt.Print("Enter Item Number that you want to Delete from your list: ")
			fmt.Scan(&id)
			DeleteProductById(id, &Product)
			fmt.Println()
			LoadAgain()
		case 6:
			fmt.Println()
			Invoice(Product)
			UserSession()
		case 7:
			fmt.Println()
			main()
		default:
			fmt.Printf("Re-enter your choice!\n")
			ProductMenu()
		}
	}

	}
}