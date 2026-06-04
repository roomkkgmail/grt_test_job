package workerpool

import (
	"fmt"
	"sync"
	"time"
)

// RunWorkers creates numWorkers to process numJobs using a worker pool pattern.
func RunWorkers(numWorkers int, numJobs int) {
	jobs := make(chan int, numJobs)
	var wg sync.WaitGroup

	// Start workers
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for job := range jobs {
				fmt.Printf("Worker [%d] processing job [%d]\n", id, job)
				time.Sleep(1 * time.Second)
			}
		}(w)
	}

	// Send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Wait for all workers to finish
	wg.Wait()

	// All workers have completed processing
	fmt.Println("All jobs have been processed.")
}
