package main

import (
	"fmt"
	"sync"
	"time"
)

//task definition
type Task struct{
	ID int  
}

//way to process the tasks
func(t *Task)Process(){
	fmt.Printf("processing task %d\n", t.ID)
	//simulate a time consuming process 
	time.Sleep(2*time.Second)
}

//worker pool definition
type WorkerPool struct{
	Tasks []Task
	concurrency int
	tasksChan chan Task
	wg sync.WaitGroup
}

//function to execute the worker pool
func (wp *WorkerPool) Worker(){
	for task := range wp.tasksChan {
		task.Process()
		wp.wg.Done()
	}
}

func (wp *WorkerPool) Run(){
	//Initialize the tasks channel
	wp.tasksChan = make(chan Task, len(wp.Tasks))

	//Start wokers
	for i := 0; i< wp.concurrency; i++ {
		go wp.Worker()
	}

	//send tasks to the tasks channel
	wp.wg.Add(len(wp.Tasks))
	for _, task := range wp.Tasks{
		wp.tasksChan <- task
	}
	close(wp.tasksChan)

	//wait for all tasks to finish
	wp.wg.Wait()
}