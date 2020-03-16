package calendar

import (
	"github.com/mikhaylovv/otus-go/hw_8/calendar/storage"
	"github.com/mikhaylovv/otus-go/hw_8/calendar/storage/inmemorystorage"
)

// Calendar - base structure for Calendar micro service
type Calendar struct {
	storage storage.Storage
}

// NewCalendar - creates an empty Calendar
func NewCalendar() Calendar {
	return Calendar{
		storage: &inmemorystorage.InMemoryStorage{},
	}
}

