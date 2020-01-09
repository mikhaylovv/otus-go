package hw5

import (
	"errors"
	"sync"
)

// Run - function for parallel execution of func() error tasks
// Параллельное исполнение
// Цель: Написать функцию для параллельного выполнения N заданий (т.е. в N параллельных горутинах).
// Функция должна останавливать свою работу если произошло M ошибок Сигнатура функции:
// Run(task []func()error, N int, M int) error Учесть что задания могут выполняться разное время
// Длинна списка задач L = len(tasks) может быть больше или меньше N.
func Run(tasks []func() error, workersNum uint, maxErrorsCount uint) error {

	if workersNum == 0 {
		return errors.New("workers num == 0, can't work without workers")
	}

	wg := &sync.WaitGroup{}
	errChan := make(chan error, workersNum)
	taskChan := make(chan func() error, workersNum)
	resChan := make(chan struct{}, workersNum)

	wg.Add(int(workersNum))
	for i := uint(0); i < workersNum; i++ {
		go work(taskChan, errChan, resChan, wg)
	}

	defer func() {
		close(taskChan)
		wg.Wait()
	}()

	var errsCounter uint = 0

	numTasks := uint(len(tasks))
	for i := uint(0); i < numTasks; i += workersNum {
		step := i + workersNum
		for j := i; j < step && j < numTasks; j++ {
			taskChan <- tasks[j]
		}

		for j := i; j < step && j < numTasks; j++ {

			select {
			case <-errChan:
				errsCounter++
				if errsCounter > maxErrorsCount {
					return errors.New("errors count exceed")
				}

			case <-resChan:
			}
		}

	}

	return nil
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
