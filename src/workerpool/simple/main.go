package main

import "fmt"

func main(){
	//create new tasks
	tasks := make([]Task, 20)
	for i := 0; i<20; i++{
		tasks[i] =Task{ID: i+1}
	}

	// create a worker pool
	wp := WorkerPool{
		Tasks: tasks,
		concurrency: 5,//number of workers that can run at time
	}

	//run the pool
	wp.Run()
	fmt.Println("all tasks have been processed.!")
}