package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

var sig chan struct{}

func doTask(t Task, ch chan struct{}) {
	if err := t(); err != nil {
		select {
		case <-sig:
			return
		default:
			ch <- struct{}{}
		}
	}
}

func handleTasks(tasks []Task, ch chan struct{}) {
	for _, t := range tasks {
		select {
		case <-sig:
			return
		default:
			doTask(t, ch)
		}
	}
}

func checkErrors(m int, ch chan struct{}) error {
	var errCount int
	for range ch {
		errCount++
		if errCount >= m {
			close(sig)
			return ErrErrorsLimitExceeded
		}
	}
	return nil
}

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	count := len(tasks)
	chErrCount := make(chan struct{}, count)
	sig = make(chan struct{})

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
		wg.Wait()
		close(chErrCount)
	}()
	err := checkErrors(m, chErrCount)
	return err
}
