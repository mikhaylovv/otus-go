package storage

import "errors"

var (
	// ErrEventNotFound - storage error returned if error can not be found in storage
	ErrEventNotFound    = errors.New("no such event in storage")

	// ErrZeroRowsAffected - storage error returned if no new errors weer added in db
	ErrZeroRowsAffected = errors.New("zero rows affected")
)
