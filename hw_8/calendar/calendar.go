package calendar

import (
	"github.com/mikhaylovv/otus-go/hw_8/calendar/storage"
	"github.com/mikhaylovv/otus-go/hw_8/grpcserver"
	"github.com/mikhaylovv/otus-go/hw_8/httpserver"
	"go.uber.org/zap"
	"log"
)

// Calendar - base structure for Calendar micro service
type Calendar struct {
	storage storage.Storage
	gserver *grpcserver.Server
	hserver *httpserver.HTTPServer
	logger  *zap.Logger
}

// NewCalendar - creates an empty Calendar
func NewCalendar(s storage.Storage, hs *httpserver.HTTPServer, gs *grpcserver.Server, lg *zap.Logger) Calendar {
	return Calendar{
		storage: s,
		gserver: gs,
		hserver: hs,
		logger:  lg,
	}
}

func (c *Calendar) Start() {
	go func() {
		if err := c.hserver.StartListen(); err != nil {
			log.Fatal("can't start http server", err)
		}

	}()

	if err := c.gserver.StartListen(); err != nil {
		log.Fatal("can't start grpcs server", err)
	}
}
