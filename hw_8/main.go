package main

import (
	"github.com/mikhaylovv/otus-go/hw_8/calendar"
	"github.com/mikhaylovv/otus-go/hw_8/calendar/storage/inmemorystorage"
	"time"
)

func main() {
	c := calendar.Calendar{
		Storage: &inmemorystorage.InMemoryStorage{},
	}
	_, _ = c.Storage.GetEvents(time.Now(), time.Now())
}
