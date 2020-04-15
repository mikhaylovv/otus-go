package dbstorage

import (
	"github.com/jmoiron/sqlx"
	"github.com/mikhaylovv/otus-go/hw_8/calendar/storage"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"time"
)

// SQLXStorage - sqlx db storage for Calendar Events
type SQLXStorage struct {
	db     *sqlx.DB
	logger *zap.Logger
}

// NewSQLXStorage - constructor for SQLXStorage object
func NewSQLXStorage(db *sqlx.DB, logger *zap.Logger) *SQLXStorage {
	return &SQLXStorage{
		db:     db,
		logger: logger,
	}
}

// AddEvent - add new event in Storage
func (s SQLXStorage) AddEvent(e storage.Event) (uint, error) {
	lastInsertID := 0
	err := s.db.QueryRow("insert into events (date, title, description) "+
		"values ($1, $2, $3) returning id", e.Date, e.Title, e.Description).Scan(&lastInsertID)

	if err != nil {
		return 0, errors.Wrap(err, "AddEvent fail on DB insert query")
	}

	return uint(lastInsertID), nil
}

// DeleteEvent - delete existing event from storage
func (s SQLXStorage) DeleteEvent(id uint) error {
	res, err := s.db.Exec("delete from events where id = $1", id)

	if err != nil {
		return errors.Wrap(err, "DeleteEvent fail on DB delete query")
	}

	if ra, _ := res.RowsAffected(); ra == 0 {
		return storage.ErrZeroRowsAffected
	}

	return nil
}

// ChangeEvent - changes existing event
func (s SQLXStorage) ChangeEvent(new storage.Event) error {
	res, err := s.db.NamedExec(
		"update events "+
			"set date=:date, title=:title, description=:description "+
			"where id=:id",
		map[string]interface{}{
			"id":          new.ID,
			"date":        new.Date,
			"title":       new.Title,
			"description": new.Description,
		},
	)

	if err != nil {
		return errors.Wrap(err, "ChangeEvent fail on DB update query")
	}

	if ra, _ := res.RowsAffected(); ra == 0 {
		return storage.ErrZeroRowsAffected
	}

	return nil
}

// GetEvents - gets events from date to date
func (s SQLXStorage) GetEvents(from time.Time, to time.Time) ([]storage.Event, error) {
	type eventDTO struct {
		ID int `db:"id"`
		Date time.Time `db:"date"`
		Title string `db:"title"`
		Description string `db:"description"`
	}

	dst := make([]eventDTO, 0)
	err := s.db.Select(dst, "select id, date, title, description from events where date between $1, $2", from, to)

	if err != nil {
		return nil, errors.Wrap(err, "ChangeEvent fail on DB update query")
	}

	var ret []storage.Event
	for _, e := range dst {
		ret = append(ret, storage.Event{
			ID:          uint(e.ID),
			Date:        e.Date,
			Title:       e.Title,
			Description: e.Description,
		})
	}

	return ret, nil
}
