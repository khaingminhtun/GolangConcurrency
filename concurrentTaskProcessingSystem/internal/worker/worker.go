package worker

import (
	"fmt"
	"sync"
	"time"

	"github.com/khaingminhtun/job-system/internal/queue"
)

func StartWorker(
	id int,
	q *queue.JobQueue,
	wg *sync.WaitGroup) {

	defer wg.Done()

	for {
		j, ok := q.Dequeue()
		if !ok {
			fmt.Printf(
				"Worker %d: no more jobs\n",
				id,
			)
			return
		}

		fmt.Printf(
			"Worker %d processing job %d: %s\n",
			id,
			j.ID,
			j.Data,
		)

		time.Sleep(time.Second)
	}
}
