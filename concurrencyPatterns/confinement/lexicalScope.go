package main

import (
	"fmt"
	"sync"
)

type Log struct {
	ID      int
	Message string
}

type Result struct {
	WorkerID int
	LogID    int
	Length   int
}

func generateLogs() <-chan Log {
	out := make(chan Log)

	go func() {
		defer close(out)

		for i := 1; i <= 20; i++ {
			out <- Log{ID: i, Message: fmt.Sprintf("Log message %d", i)}
		}
	}()
	return out
}

func worker(
	id int,
	logs <-chan Log,
	results chan<- Result,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	processed := 0

	for log := range logs {

		localLog := log

		processed++

		results <- Result{
			WorkerID: id,
			LogID:    localLog.ID,
			Length:   len(localLog.Message),
		}
	}

	fmt.Printf("worker %d processed %d logs\n", id, processed)
}

func aggregate(results <-chan Result) {
	total := 0

	for result := range results {
		total++

		fmt.Printf(
			"[RESULT] worker=%d log=%d\n",
			result.WorkerID,
			result.LogID,
		)
	}

	fmt.Println("total processed:", total)
}

func main() {
	logs := generateLogs()

	results := make(chan Result)

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)

		go worker(i, logs, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	aggregate(results)
}
