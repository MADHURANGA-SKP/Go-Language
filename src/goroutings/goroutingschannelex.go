package main

import (
	"fmt"
	"time"
)

func wokers(id int, jobs <-chan int, results chan<- int){
	for j := range jobs {
		fmt.Println("woker", id, "Started job", j)
		time.Sleep(time.Second)
		fmt.Println("woker", id, "Finished job" ,j)
		results <- j *2
	}
}

func main(){
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go wokers(w, jobs , results)
	}

	for j := 1; j<=3; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a<= numJobs; a++ {
		<- results
	}
}

//this occrs an error