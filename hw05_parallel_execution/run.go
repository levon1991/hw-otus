package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// for signal if count of error over limit.
var sig chan struct{}

// for return error or nil when finish all tasks.
var errChan chan bool

func doTask(t Task, ch chan bool) {
	if err := t(); err != nil {
		select {
		case <-sig:
			return
		default:
			ch <- false
		}
	} else {
		ch <- true
	}
}

func handleTasks(tasks []Task, ch chan bool) {
	for _, t := range tasks {
		select {
		case <-sig:
			return
		default:
			doTask(t, ch)
		}
	}
}

func checkErrors(m, count int, ch chan bool) {
	var errCount, workerCount int

	for v := range ch {
		if v {
			workerCount++
			if workerCount == count {
				close(sig)
				errChan <- true
				return
			}
		} else {
			errCount++
			if errCount >= m {
				close(sig)
				errChan <- false
				return
			}
		}
	}
}

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	count := len(tasks)
	chErrCount := make(chan bool, n+m)
	sig = make(chan struct{})
	errChan = make(chan bool)
	defer close(errChan)

	residue := count % n
	step := count / n

	wg := sync.WaitGroup{}

	var start, end int
	for i := 0; i < count; i += step {
		// Это часть спецально для того чтоб можно было ровно делить задачи между воркерами в том случае если
		// количество задач не кратно количество воркеров
		end = i + step
		start = i
		if residue > 0 {
			end++
			residue--
			i++
		}
		taskSubList := tasks[start:end]
		wg.Add(1)
		go func() {
			handleTasks(taskSubList, chErrCount)
			wg.Done()
		}()
	}

	go func() {
		checkErrors(m, count, chErrCount)
	}()
	wg.Wait()
	close(chErrCount)

	// Wait for result (all complete or have a more than M errors)
	if <-errChan {
		return nil
	}
	return ErrErrorsLimitExceeded
}
