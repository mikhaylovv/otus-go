package inmemorystorage

import (
	"github.com/mikhaylovv/otus-go/hw_8/calendar/storage"
	"time"
)

// InMemoryStorage - in memory storage for Calendar Events
type InMemoryStorage struct {
	events map[uint]storage.Event
	last uint
}

func (s *InMemoryStorage) findEventIndex(event storage.Event) (uint, error) {
	for idx, ev := range s.events {
		if ev == event {
			return idx, nil
		}
	}

	return 0, storage.ErrEventNotFound
}

// AddEvent - add new event in Storage or error ErrEventAlreadyExist
func (s *InMemoryStorage) AddEvent(e storage.Event) (uint, error) {
	idx := s.last
	s.events[idx] = e
	s.last++
	return idx, nil
}

// DeleteEvent - delete existing event from storage
func (s *InMemoryStorage) DeleteEvent(id uint) error {
	delete(s.events, id)
	return nil
}

// ChangeEvent - changes existing event
func (s *InMemoryStorage) ChangeEvent(oldID uint, new storage.Event) error {
	s.events[oldID] = new
	return nil
}

// GetEvents - gets events from date to date, error always nil
func (s *InMemoryStorage) GetEvents(from time.Time, to time.Time) ([]storage.Event, error) {
	var res []storage.Event
	for _, ev := range s.events {
		if ev.Date.After(from) && ev.Date.Before(to) {
			res = append(res, ev)
		}
	}

	return res, nil
}
