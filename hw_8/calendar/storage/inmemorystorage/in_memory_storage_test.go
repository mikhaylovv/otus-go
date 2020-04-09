package inmemorystorage

import (
	"github.com/mikhaylovv/otus-go/hw_8/calendar/storage"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type fixture struct {
	s  *InMemoryStorage
	e  storage.Event
	e2 storage.Event
}

func setUp() fixture {
	s := NewInMemoryStorage()
	e := storage.Event{
		Date:        time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
		Title:       "testTitle",
		Description: "testDescription",
	}
	e2 := storage.Event{
		Date:        time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
		Title:       "testTitle2",
		Description: "testDescription2",
	}
	return fixture{s, e, e2}
}

func TestAddEvent(t *testing.T) {
	f := setUp()
	idx, err := f.s.AddEvent(f.e)
	f.e.ID = idx
	assert.NoError(t, err)
	assert.Equal(t, f.e, f.s.events[idx])
	idx, err = f.s.AddEvent(f.e2)
	f.e2.ID = idx
	assert.NoError(t, err)
	assert.Equal(t, 2, len(f.s.events))
	assert.Equal(t, f.e2, f.s.events[idx])
}

func TestChangeEvent(t *testing.T) {
	f := setUp()
	idx, err := f.s.AddEvent(f.e)
	assert.NoError(t, err)
	f.e2.ID = idx
	assert.NoError(t, f.s.ChangeEvent(f.e2))
	assert.Equal(t, 1, len(f.s.events))
	assert.Equal(t, f.e2, f.s.events[idx])
}

func TestDeleteEvent(t *testing.T) {
	f := setUp()
	idx1, err := f.s.AddEvent(f.e)
	f.e.ID = idx1
	assert.NoError(t, err)
	idx2, err := f.s.AddEvent(f.e2)
	f.e2.ID = idx2
	assert.NoError(t, err)
	assert.NoError(t, f.s.DeleteEvent(idx2))
	assert.Equal(t, 1, len(f.s.events))
	assert.Equal(t, f.e, f.s.events[idx1])
	assert.NoError(t, f.s.DeleteEvent(idx1))
	assert.Equal(t, 0, len(f.s.events))
}

func TestGetEvents(t *testing.T) {
	f := setUp()
	idx, err := f.s.AddEvent(f.e)
	f.e.ID = idx
	assert.NoError(t, err)
	_, err = f.s.AddEvent(f.e2)
	assert.NoError(t, err)

	ev, err := f.s.GetEvents(
		time.Date(2019, time.January, 2, 0, 0, 0, 0, time.UTC),
		time.Date(2020, time.January, 2, 0, 0, 0, 0, time.UTC),
	)

	assert.NoError(t, err)
	assert.Equal(t, f.e, ev[0])

	ev, err = f.s.GetEvents(
		time.Date(2018, time.January, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC),
	)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(ev))
}
