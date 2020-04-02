package grpcserver

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/mikhaylovv/otus-go/hw_8/calendar/storage"
	"github.com/mikhaylovv/otus-go/hw_8/proto/calendarpb"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"net"
	"time"
)

type Server struct {
	storage storage.Storage
	addr    string
	lg      *zap.Logger
}

func NewServer(s storage.Storage, addr string, l *zap.Logger) *Server {
	return &Server{
		storage: s,
		addr:    addr,
		lg:      l,
	}
}

func (s *Server) AddEvent(_ context.Context, ev *calendarpb.CalendarEvent) (*calendarpb.CalendarEventId, error) {
	id, err := s.storage.AddEvent(storage.Event{
		Date:        time.Unix(ev.Date.Seconds, int64(ev.Date.Nanos)),
		Title:       ev.Title,
		Description: ev.Description,
	})

	if err != nil {
		return nil, processError(err)
	}

	ret := &calendarpb.CalendarEventId{
		Id: uint32(id),
	}

	return ret, nil
}

func (s *Server) DeleteEvent(_ context.Context, evId *calendarpb.CalendarEventId) (*empty.Empty, error) {
	err := s.storage.DeleteEvent(uint(evId.Id))
	return nil, processError(err)
}

func (s *Server) ChangeEvent(_ context.Context, ev *calendarpb.CalendarEvent) (*empty.Empty, error) {
	e := storage.Event{
		Id:          uint(ev.Id),
		Date:        time.Unix(ev.Date.Seconds, int64(ev.Date.Nanos)),
		Title:       ev.Title,
		Description: ev.Description,
	}

	err := s.storage.ChangeEvent(e)
	return nil, processError(err)
}

func (s *Server) GetEvents(_ context.Context, di *calendarpb.DateInterval) (*calendarpb.CalendarEvents, error) {
	from, err := ptypes.Timestamp(di.GetFrom())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Bad DateInterval from")
	}

	to, err := ptypes.Timestamp(di.GetTo())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Bad DateInterval from")
	}

	evs, err := s.storage.GetEvents(from, to)
	if err != nil {
		return nil, processError(err)
	}

	ret := &calendarpb.CalendarEvents{
		Events: make([]*calendarpb.CalendarEvent, len(evs)),
	}

	for _, ev := range evs {
		ret.Events = append(ret.Events, &calendarpb.CalendarEvent{
			Id: uint32(ev.Id),
			Date: &timestamp.Timestamp{
				Seconds: ev.Date.Unix(),
				Nanos:   int32(ev.Date.UnixNano()),
			},
			Title:       ev.Title,
			Description: ev.Description,
		})
	}

	return ret, nil
}

func (s *Server) StartListen() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}

	srv := grpc.NewServer()
	reflection.Register(srv)

	calendarpb.RegisterCalendarServer(srv, s)
	err = srv.Serve(lis)
	if err != nil {
		return err
	}

	return nil
}

func processError(err error) error {
	if err != nil {
		if err == storage.ErrEventNotFound {
			return status.Error(codes.InvalidArgument, err.Error())
		}
		return status.Error(codes.Internal, err.Error())
	}
	return nil
}
