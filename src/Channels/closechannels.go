package main

import "fmt"

func main(){
	jobs := make(chan int,5)
	done := make(chan bool)

	go func(){
		for{
			j, more := <- jobs
			if more {
				fmt.Println("recived job",j)
			}else{
				fmt.Println("recived all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <=3 ; j++ {
		jobs <- j
		fmt.Println("sent job",j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done

	_, ok := <-jobs
	fmt.Println("recived more jobs", ok)
}