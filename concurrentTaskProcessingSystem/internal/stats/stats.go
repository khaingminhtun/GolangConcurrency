package stats

import "sync"

type Stats struct {
	mu sync.RWMutex

	ProcessedJobs int
	FailedJobs    int
}

func (s *Stats) IncrementProcessed() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.ProcessedJobs++
}

func (s *Stats) GetProcessed() int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.ProcessedJobs
}
