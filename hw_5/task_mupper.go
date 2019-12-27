package hw5

import (
	"errors"
	"sync"
)

func Run(tasks []func() error, workersNum uint, maxErrorsCount uint) error {

	if workersNum == 0 {
		return errors.New("workers num == 0, can't work without workers")
	}

	wg := &sync.WaitGroup{}
	errChan := make(chan error, workersNum)
	taskChan := make(chan func() error, len(tasks))
	resChan := make(chan struct{}, workersNum)

	wg.Add(int(workersNum))
	for i := uint(0); i < workersNum; i++ {
		go work(taskChan, errChan, resChan, wg)
	}

	for _, f := range tasks {
		taskChan <- f
	}

	var errsCounter uint = 0
	var tasksCounter uint = 0

	defer func() {
		close(taskChan)

		for len(resChan) != 0 {
			<- resChan
		}
		for len(errChan) != 0 {
			<- errChan
		}

		wg.Wait()
	}()

	for {
		select {
		case <-errChan:
			errsCounter++
			if errsCounter > maxErrorsCount {
				return errors.New("errors count exceed")
			}

		case <-resChan:
			tasksCounter++
		}

		if (errsCounter + tasksCounter) == uint(len(tasks)) {
			return nil
		}
	}
}

func work(
	tasks <-chan func() error,
	err chan<- error,
	res chan<- struct{},
	group *sync.WaitGroup) {

	for t := range tasks {
		if er := t(); er != nil {
			err <- er
		} else {
			res <- struct{}{}
		}
	}
	group.Done()
}
