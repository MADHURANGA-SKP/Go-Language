package main

import "fmt"

func main(){
	//data
	tasks := []Task{
	&EmailTask{Email: "john@example.com", Subject: "Important Update", Message: "Please find the attached document."},
	&ImageProcessingTask{ImageUrl: "/images/photo123.jpg"},

	&EmailTask{Email: "susan@example.com", Subject: "Meeting Reminder", Message: "Don't forget our team meeting at 3 PM."},
	&ImageProcessingTask{ImageUrl: "/images/scenery456.jpg"},

	&EmailTask{Email: "alex@example.com", Subject: "Feedback Request", Message: "We value your opinion. Please share your feedback."},
	&ImageProcessingTask{ImageUrl: "/images/art789.jpg"},

	&EmailTask{Email: "emily@example.com", Subject: "Vacation Photos", Message: "Enjoy these snapshots from our trip!"},
	&ImageProcessingTask{ImageUrl: "/images/beach101.jpg"},

	&EmailTask{Email: "david@example.com", Subject: "Project Update", Message: "Here's the latest progress report."},
	&ImageProcessingTask{ImageUrl: "/images/chart567.jpg"},
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