package main

import (
	"fmt"
	"sync"
	"time"
)

//task definition
type Task interface{
	Process()  
}

//email task definition
type EmailTask struct{
	Email string
	Subject string
	Message string
}

//way to process the email tasks
func(t *EmailTask)Process(){
	fmt.Printf("Sending Email to %s\n", t.Email)
	//simulate a time consuming process 
	time.Sleep(2*time.Second)
}

//Image processing task
type ImageProcessingTask struct{
	ImageUrl string
}

//way to process the Image processing task
func(t *ImageProcessingTask)Process(){
	fmt.Printf("Proccessing the Image %s\n", t.ImageUrl)
	//simulate a time consuming process 
	time.Sleep(5*time.Second)
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