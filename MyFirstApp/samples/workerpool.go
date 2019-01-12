package samples

import "fmt"
import "time"

/*
Here’s the worker, of which we’ll run several concurrent instances.
These workers will receive work on the jobs channel and send the corresponding results on results.
We’ll sleep a second per job to simulate an expensive task.
results is the output channel
job is the incoming job channel
*/
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func RunWorkers(workerCount, jobCount int) {

	// Two channels to allow communication from the go threads!
	// in order to send jobs over to each worker!
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// loop through and and create workers - initially stopped as no jobs present
	for w := 1; w <= workerCount; w++ {
		go worker(w, jobs, results)
	}
	//Here we send X jobs and then close that channel to indicate that’s all the work we have.
	for j := 1; j <= jobCount; j++ {
		jobs <- j
	}

	// close that channel after jobs sent!
	close(jobs)
	//Finally we collect all the results of the work.
	for a := 1; a <= jobCount; a++ {
		<-results
	}
}
