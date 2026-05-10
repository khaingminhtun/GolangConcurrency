package worker

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/khaingminhtun/job-system/internal/pool"
	"github.com/khaingminhtun/job-system/internal/queue"
	"github.com/khaingminhtun/job-system/internal/stats"
)

func StartWorker(
	ctx context.Context,
	id int,
	q *queue.JobQueue,
	stats *stats.Stats,
	wg *sync.WaitGroup) {

	defer wg.Done()

	for {

		// check shutdown first
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d shutting down\n", id)
			return
		default:
		}
		// try get job
		j, ok := q.Dequeue(ctx)
		if !ok {
			fmt.Printf("Worker %d shutting down\n", id)
			return
		}

		// get reusable buffer
		buf := pool.BufferPool.Get().([]byte)

		// simulate processing
		copy(buf, j.Data)
		fmt.Printf(
			"Worker %d processing job %d: %s\n",
			id,
			j.ID,
			j.Data,
		)

		time.Sleep(time.Second)
		stats.IncrementProcessed()

		// return buffer to pool
		pool.BufferPool.Put(buf)
	}
}
