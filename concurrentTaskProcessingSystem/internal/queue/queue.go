package queue

import (
	"context"
	"sync"

	"github.com/khaingminhtun/job-system/internal/job"
)

type JobQueue struct {
	jobs []job.Job
	mu   sync.Mutex
	cond *sync.Cond
}

func NewJobQueue() *JobQueue {
	q := &JobQueue{}
	q.cond = sync.NewCond(&q.mu)
	return q
}

func (q *JobQueue) Enqueue(j job.Job) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.jobs = append(q.jobs, j)

	q.cond.Signal()

}

func (q *JobQueue) Dequeue(ctx context.Context) (job.Job, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()

	for len(q.jobs) == 0 {

		select {
		case <-ctx.Done():
			return job.Job{}, false
		default:
		}
		q.cond.Wait()
	}

	j := q.jobs[0]
	q.jobs = q.jobs[1:]

	return j, true
}

func (q *JobQueue) WakeAll() {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.cond.Broadcast()
}
