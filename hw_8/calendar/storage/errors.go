package storage

import "errors"

var (
	EventAlreadyExistError = errors.New("event already exist")
	EventNotFoundError     = errors.New("no such event in storage")
)