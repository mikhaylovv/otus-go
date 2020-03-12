package storage

import "time"

type Event struct {
	Date        time.Time
	Title       string
	Description string
}

type Storage interface {
	AddEvent(e Event) error
	DeleteEvent(e Event) error
	ChangeEvent(old Event, new Event) error
	GetEvents(from time.Time, to time.Time) ([]Event, error)
}
