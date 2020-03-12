package calendar

import "github.com/mikhaylovv/otus-go/hw_8/calendar/storage"

// Calendar - base structure for Calendar micro service
type Calendar struct {
	Storage storage.Storage
}
