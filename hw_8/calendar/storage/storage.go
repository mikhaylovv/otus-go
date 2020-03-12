package storage

import "time"

// Event -  structure of  Calendar simple event
type Event struct {
	Date        time.Time
	Title       string
	Description string
}

// Storage - provides interface for any Event storage
type Storage interface {
	AddEvent(e Event) error
	DeleteEvent(e Event) error
	ChangeEvent(old Event, new Event) error
	GetEvents(from time.Time, to time.Time) ([]Event, error)
}
