package storage

import "time"

// Event -  structure of  Calendar simple event
type Event struct {
	ID          uint
	Date        time.Time
	Title       string
	Description string
}

// Storage - provides interface for any Event storage
type Storage interface {
	AddEvent(e Event) (uint, error)
	DeleteEvent(id uint) error
	ChangeEvent(new Event) error
	GetEvents(from time.Time, to time.Time) ([]Event, error)
}
