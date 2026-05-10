package main

import (
	"fmt"
	"sync"

	"github.com/khaingminhtun/job-system/internal/job"
	"github.com/khaingminhtun/job-system/internal/queue"
	"github.com/khaingminhtun/job-system/internal/worker"
)

func main() {
	q := &queue.JobQueue{}

	// crate jobs
	for i := 1; i <= 30; i++ {
		q.Enqueue(job.Job{
			ID:   i,
			Data: fmt.Sprintf("task-%d", i),
		})
	}

	var wg sync.WaitGroup

	workerCount := 3

	wg.Add(workerCount)

	for i := 1; i <= workerCount; i++ {
		go worker.StartWorker(i, q, &wg)
	}

	wg.Wait()

	fmt.Println("All workers completed")
}
