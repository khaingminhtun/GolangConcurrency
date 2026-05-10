package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/khaingminhtun/job-system/internal/job"
	"github.com/khaingminhtun/job-system/internal/queue"
	"github.com/khaingminhtun/job-system/internal/worker"
)

func main() {
	q := queue.NewJobQueue()

	var wg sync.WaitGroup

	workerCount := 3

	wg.Add(workerCount)

	// Start workers
	for i := 1; i <= workerCount; i++ {
		go worker.StartWorker(i, q, &wg)
	}

	// Producer (simulate streaming jobs)
	for i := 1; i <= 10; i++ {
		q.Enqueue(job.Job{
			ID:   i,
			Data: fmt.Sprintf("job-%d", i),
		})

		time.Sleep(500 * time.Millisecond)
	}

	// Wait so workers can process
	time.Sleep(5 * time.Second)

	fmt.Println("main finished producing")
}
