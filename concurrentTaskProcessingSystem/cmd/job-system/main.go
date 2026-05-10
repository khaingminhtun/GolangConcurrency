package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/khaingminhtun/job-system/internal/job"
	"github.com/khaingminhtun/job-system/internal/queue"
	"github.com/khaingminhtun/job-system/internal/stats"
	"github.com/khaingminhtun/job-system/internal/worker"
)

func main() {
	q := queue.NewJobQueue()

	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup

	workerCount := 3

	wg.Add(workerCount)

	statistics := &stats.Stats{}

	// Start workers
	for i := 1; i <= workerCount; i++ {
		go worker.StartWorker(ctx, i, q, statistics, &wg)
	}

	// handle os signals
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	// Producer (simulate streaming jobs)
	go func() {
		for i := 1; i <= 10; i++ {
			q.Enqueue(job.Job{
				ID:   i,
				Data: fmt.Sprintf("job-%d", i),
			})

			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for {
			select {
			case <-ctx.Done():
				return

			default:
				fmt.Printf(
					"\n[STATS] Processed Jobs: %d\n",
					statistics.GetProcessed(),
				)

				time.Sleep(2 * time.Second)
			}
		}
	}()

	//  Wait for interrupt
	<-signalChan
	fmt.Println("\nShutdown signal received")

	//  Cancel all workers
	cancel()

	//  Wait for cleanup
	fmt.Println("waiting for workers...")

	q.WakeAll()

	wg.Wait()
	fmt.Println("DONE WAITING")

	fmt.Println("All workers stopped cleanly")
}
