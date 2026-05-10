package queue

import (
	"sync"

	"github.com/khaingminhtun/job-system/internal/job"
)

type JobQueue struct {
	jobs []job.Job
	mu   sync.Mutex
}

func (q *JobQueue) Enqueue(j job.Job) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.jobs = append(q.jobs, j)
}

func (q *JobQueue) Dequeue() (job.Job, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.jobs) == 0 {
		return job.Job{}, false
	}

	j := q.jobs[0]
	q.jobs = q.jobs[1:]

	return j, true
}
