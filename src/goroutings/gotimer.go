package main

import (
	"fmt"
	"time"
)

func main(){
	timer1 := time.NewTimer(2*time.Second)


	<-timer1.C //<-timer1.C blocks on the timer’s channel 
	//C until it sends a value indicating that the timer fired.
	fmt.Println("timer 1 fired")

	timer2 :=time.NewTimer(time.Second)
	go func(){
		<-timer2.C//block the execution until the timer expires.
		fmt.Println("timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("timer 2 stoped")
	}
	time.Sleep(2*time.Second)
}