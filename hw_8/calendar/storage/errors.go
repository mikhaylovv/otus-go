package storage

import "errors"

var (
	// ErrEventNotFound - storage error returned if error can not be found in storage
	ErrEventNotFound = errors.New("no such event in storage")
)
