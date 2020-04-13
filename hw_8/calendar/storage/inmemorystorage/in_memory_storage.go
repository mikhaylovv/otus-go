package inmemorystorage

import (
	"github.com/mikhaylovv/otus-go/hw_8/calendar/storage"
	"sync"
	"time"
)

// InMemoryStorage - in memory storage for Calendar Events
type InMemoryStorage struct {
	events map[uint]storage.Event
	last   uint
	mutex  sync.RWMutex
}

// NewInMemoryStorage - construct NewInMemoryStorage object
func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		events: make(map[uint]storage.Event),
		last:   1,
	}
}

// AddEvent - add new event in Storage or error ErrEventAlreadyExist
func (s *InMemoryStorage) AddEvent(e storage.Event) (uint, error) {
	idx := s.last
	s.mutex.Lock()
	e.ID = idx
	s.events[idx] = e
	s.last++
	s.mutex.Unlock()
	return idx, nil
}

// DeleteEvent - delete existing event from storage
func (s *InMemoryStorage) DeleteEvent(id uint) error {
	s.mutex.Lock()
	delete(s.events, id)
	s.mutex.Unlock()
	return nil
}

// ChangeEvent - changes existing event. If event not found send ErrEventNotFound
func (s *InMemoryStorage) ChangeEvent(new storage.Event) error {
	s.mutex.RLock()
	if _, ok := s.events[new.ID]; ok {
		s.mutex.RUnlock()
		s.mutex.Lock()
		s.events[new.ID] = new
		s.mutex.Unlock()
		return nil
	}

	s.mutex.RUnlock()
	return storage.ErrEventNotFound
}

// GetEvents - gets events from date to date, error always nil
func (s *InMemoryStorage) GetEvents(from time.Time, to time.Time) ([]storage.Event, error) {
	var res []storage.Event

	s.mutex.Lock()
	for _, ev := range s.events {
		if ev.Date.After(from) && ev.Date.Before(to) {
			res = append(res, ev)
		}
	}
	s.mutex.Unlock()

	return res, nil
}
