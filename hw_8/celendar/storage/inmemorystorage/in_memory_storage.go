package inmemorystorage

import (
	"github.com/mikhaylovv/otus-go/hw_8/celendar/storage"
	"time"
)

type InMemoryStorage struct {
	events []storage.Event
}

func (s *InMemoryStorage) findEventIndex(event storage.Event) (int, error) {
	for idx, ev := range s.events {
		if ev == event {
			return idx, nil
		}
	}

	return -1, storage.EventNotFoundError
}

func (s *InMemoryStorage) removeEvent(i int) {
	lastIdx := len(s.events)
	s.events[lastIdx-1], s.events[i] = s.events[i], s.events[lastIdx-1]
	s.events = s.events[:lastIdx-1]
}

func (s *InMemoryStorage) AddEvent(e storage.Event) error {
	if _,  err := s.findEventIndex(e); err != storage.EventNotFoundError {
		return storage.EventAlreadyExistError
	}

	s.events = append(s.events, e)
	return nil
}

func (s *InMemoryStorage) DeleteEvent(e storage.Event) error {
	idx, err := s.findEventIndex(e)
	if err != nil {
		return err
	}

	s.removeEvent(idx)
	return nil
}

func (s *InMemoryStorage) ChangeEvent(old storage.Event, new storage.Event) error {
	idx, err := s.findEventIndex(old)
	if err != nil {
		return err
	}

	s.events[idx] = new
	return nil
}

func (s *InMemoryStorage) GetEvents(from time.Time, to time.Time) ([]storage.Event, error) {
	var res []storage.Event
	for _, ev := range s.events {
		if ev.Date.After(from) && ev.Date.Before(to) {
			res = append(res, ev)
		}
	}

	return res, nil
}
