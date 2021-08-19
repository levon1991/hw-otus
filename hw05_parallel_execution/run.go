package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func consumer(taskChan chan Task, errCount *int32, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range taskChan {
		if err := task(); err != nil {
			atomic.AddInt32(errCount, 1)
		}
	}
}

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	taskChan := make(chan Task)

	wg := &sync.WaitGroup{}

	var errCount int32
	for i := 0; i < n; i++ {
		wg.Add(1)
		go consumer(taskChan, &errCount, wg)
	}

	for _, task := range tasks {
		if atomic.LoadInt32(&errCount) >= int32(m) {
			break
		}
		taskChan <- task
	}
	close(taskChan)

	wg.Wait()

	if atomic.LoadInt32(&errCount) >= int32(m) {
		return ErrErrorsLimitExceeded
	}

	return nil
}
