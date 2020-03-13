package calendar

import (
	"github.com/mikhaylovv/otus-go/hw_8/calendar/storage"
	"go.uber.org/zap"
)

// Calendar - base structure for Calendar micro service
type Calendar struct {
	Storage storage.Storage
	Logger *zap.Logger
}
