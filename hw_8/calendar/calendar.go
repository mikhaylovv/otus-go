package calendar

import (
	"github.com/mikhaylovv/otus-go/hw_8/calendar/storage"
	"github.com/mikhaylovv/otus-go/hw_8/calendar/storage/inmemorystorage"
	"go.uber.org/zap"
)

// Calendar - base structure for Calendar micro service
type Calendar struct {
	storage storage.Storage
	logger *zap.Logger
}

// NewCalendar - creates an empty Calendar
func NewCalendar(lg *zap.Logger) Calendar {
	return Calendar{
		storage: &inmemorystorage.InMemoryStorage{},
		logger: lg,
	}
}

