package storage

import "errors"

var (
	// ErrEventAlreadyExist - storage error returned if event already exist in storage
	ErrEventAlreadyExist = errors.New("event already exist in storage")

	// ErrEventNotFound - storage error returned if error can not be found in storage
	ErrEventNotFound     = errors.New("no such event in storage")
)
