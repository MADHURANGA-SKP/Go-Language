package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
    TaskName string 
    TaskTime  string 
    TaskDate  string 
}

var tasks  []Task

func readTask(tasks []Task){
 if tasks == nil {
	fmt.Printf("Tasks are empty...!")
 }else{
	for _, tasks := range tasks {
		fmt.Printf("Task Name:%s | Task Time:%s | Task Date:%s\n", tasks.TaskName, tasks.TaskTime,tasks.TaskDate)	
	}
 }
}

func showTaskById(id int, tasks []Task){
	if tasks != nil {
		for _, tasklist := range tasks {
			if tasks[id] == tasklist {
			 fmt.Printf("Task Name:%s | Task Time:%s| Task Date:%s\n\n", tasklist.TaskName,tasklist.TaskTime,tasklist.TaskDate)
			}
		 }
	}else{
		fmt.Printf("Tasks are empty...!\n")
	 }
   }

func createTask(tasks *[]Task){
	var taskN string
	var taskD string
	var taskT string
	fmt.Println()

	fmt.Print("Enter Task Name : ")
	fmt.Scan(&taskN)

	fmt.Print("Enter Task Time : ")
	fmt.Scan(&taskT)

	fmt.Print("Enter Task Date : ")
	fmt.Scan(&taskD)

	Task_ := Task{TaskName:taskN, TaskDate:taskD, TaskTime:taskT}

	*tasks = append(*tasks, Task_) 
	fmt.Println("Task successfully Created..!")
	fmt.Println("\n")

}

func deleteTaskById(id int, tasks *[]Task){
	//id is index to be deleted
	i := id -1
   fmt.Println("The Task to be deleted is:", i)
   if i >=0 && i< len(*tasks) {
	//remove an element from a slice at index i
	*tasks = append((*tasks)[:i],(*tasks)[i+1:]...)
 
	//get portion of the struct till i: and i+1: 
	//then concat both parts of the struct and equal it to the *tasks
	for _, tasks := range *tasks {
		fmt.Printf("Task Name:%s | Task Time:%s | Task Date:%s\n", tasks.TaskName,tasks.TaskTime,tasks.TaskDate)	
		}
   } else {
	  fmt.Println("The given index is not in the list.")
   }

}

func updateTask(id int, tasks []Task){
		//id is index to be update
	i := id -1
	var taskN string
	var taskD string
	var taskT string
	if tasks != nil {
		fmt.Println()

		fmt.Print("Enter Task Name : ")
		fmt.Scan(&taskN)
		//& used to pass the memory address of taskN to the fmt using scan.

		fmt.Print("Enter Task Time : ")
		fmt.Scan(&taskT)

		fmt.Print("Enter Task Date : ")
		fmt.Scan(&taskD)
		
		//go through the task collection ignoreing indexes 
		//by getting values to the tasklist without specifing the type of the variable
		//check wheather user's desire thats matches with the tasks indexes
		for _, tasklist := range tasks {
			if tasks[i] == tasklist {
		//when goes throug the each index of task struct check if user index matches the tasks indexes
		//if it does update the values of that belongs to that exact index
				tasks[i] = Task{TaskName:taskN, TaskTime:taskT, TaskDate:taskD}
				fmt.Println("Task successfully Updated..!")
			}
		}
	}else{
		fmt.Printf("Tasks are empty...!\n")
	 }
}

func saveTasksToFile(tasks []Task, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(tasks)
	if err != nil {
		return err
	}

	fmt.Println("Tasks successfully saved to", filename)
	return nil
}

func loadTasksFromFile(filename string) ([]Task, error) {
	var tasks []Task

	file, err := os.Open(filename)
	if err != nil {
		return tasks, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)
	if err != nil {
		return tasks, err
	}

	fmt.Println("Tasks successfully loaded from", filename)
	return tasks, nil
}



func listOfOptions() {
	fmt.Println("HI This is an TODO APP")
	fmt.Println("1. Show List of Tasks")
	fmt.Println("2. Show Task by Task Id")
	fmt.Println("3. Add a new Task")
	fmt.Println("4. Update a Task by ID")
	fmt.Println("5. Delete a task")
	fmt.Println("0. Exit")
	fmt.Print("Enter your option number: ")
}


func reload(){
	var input string
	fmt.Print("Type 'Enter' go back to menu: ")
	fmt.Scan(&input)
	if input == "Enter"{
		menu()
	}else{
		fmt.Println("Invalid Input...!\n")
		reload()
	}
}

func menu() {
	var option int
	listOfOptions()
	for {
		fmt.Scan(&option)
	if option == 0 {
		fmt.Println("BYE BYE.!")
		os.Exit(0)
	}else{
		switch option{
		case 1:
			readTask(tasks)
			fmt.Println("\n")
			listOfOptions()
		case 2:
			var id int
			fmt.Println()
			fmt.Print("Enter Task ID that you want to get :")
			fmt.Scan(&id)
			showTaskById(id, tasks)
			reload()
		case 3:
			createTask(&tasks)
			reload()
		case 4:
			var id int
			fmt.Println()
			fmt.Print("Enter Task ID that you want to Update: ")
			fmt.Scan(&id)
			updateTask(id, tasks)
			fmt.Println("\n")
			reload()
		case 5:
			var id int
			fmt.Println()
			fmt.Print("Enter Task ID that you want to Delete: ")
			fmt.Scan(&id)
			deleteTaskById(id, &tasks)
			fmt.Println("\n")
			reload()
		default:
			fmt.Println("Re-enter your choice!")
			menu()
		}
	}

	}
}

func main() {
	// Load tasks from file at the start
	tasks, err := loadTasksFromFile("tasks.json")
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	// Run the menu
	menu()

	// Save tasks to file before exiting
	err = saveTasksToFile(tasks, "tasks.json")
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}
}