package hw5

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createTasks(task func()error, num int) []func()error {
	tasks := make([]func()error, 0, num)
	for i := 0; i < num; i++ {
		tasks = append(tasks, task)
	}
	return tasks
}

func TestRunSimplePositive(t *testing.T) {
	as := assert.New(t)

	res := make(chan struct{}, 40)
	as.Nil(Run(createTasks(func()error { res <- struct{}{}; return nil }, 30), 5, 1))
	as.EqualValues(30, len(res))
}

func TestRunSimpleNegative(t *testing.T) {
	as := assert.New(t)

	res := make(chan struct{}, 40)
	as.NotNil(Run(createTasks(func()error { res <- struct{}{}; return errors.New("lol, kek") }, 10), 5, 6))
	as.LessOrEqual(len(res), 16)
}

func TestRunSingleWorker(t *testing.T) {
	as := assert.New(t)

	res := make(chan struct{}, 40)
	as.Nil(Run(createTasks(func()error { res <- struct{}{}; return nil }, 30), 1, 1))
	as.EqualValues(30, len(res))
}

func TestRunZeroWorkers(t *testing.T) {
	as := assert.New(t)

	res := make(chan struct{}, 40)
	as.NotNil(Run(createTasks(func()error { res <- struct{}{}; return nil }, 30), 0, 1))
	as.EqualValues(0, len(res))
}

func TestRunSeveralErrors(t *testing.T) {
	as := assert.New(t)

	res := make(chan struct{}, 40)
	fns := createTasks(func()error { return errors.New("shit") }, 4)
	fns = append(fns, createTasks(func()error { res <- struct{}{}; return nil }, 20)...)
	as.Nil(Run(fns, 10, 5))
	as.EqualValues(20, len(res))
}

func TestRunSeveralErrorsNegative(t *testing.T) {
	as := assert.New(t)

	res := make(chan struct{}, 40)
	fns := createTasks(func()error { res <- struct{}{}; return errors.New("shit") }, 6)
	fns = append(fns, createTasks(func()error { res <- struct{}{}; return nil }, 20)...)
	as.NotNil(Run(fns, 10, 5))
	as.Less(len(res), 15)
}