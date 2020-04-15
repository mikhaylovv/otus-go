package main

import (
	"flag"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/mikhaylovv/otus-go/hw_8/calendar"
	"github.com/mikhaylovv/otus-go/hw_8/calendar/storage/dbstorage"
	"github.com/mikhaylovv/otus-go/hw_8/config"
	"github.com/mikhaylovv/otus-go/hw_8/grpcserver"
	"github.com/mikhaylovv/otus-go/hw_8/httpserver"
	"go.uber.org/zap"
	"io/ioutil"
	"log"
)

func newLogger(level, path string) (*zap.Logger, error) {
	atom := zap.NewAtomicLevel()
	err := atom.UnmarshalText([]byte(level))
	if err != nil {
		return nil, err
	}

	cfg := zap.NewProductionConfig()
	cfg.Level = atom
	cfg.OutputPaths = []string{path}

	return cfg.Build()
}

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "", "path to JSON config file")
	flag.Parse()

	rawCfg, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatal("read config error", err)
	}

	cfg, err := config.ParseConfig(rawCfg)
	if err != nil {
		log.Fatal("parse config error", err)
	}

	lg, err := newLogger(cfg.LogLevel, cfg.LogFile)
	if err != nil {
		log.Fatal("create logger error", err)
	}

	hsrv := httpserver.NewHTTPServer(cfg.HTTPListen, lg)

	db, err := sqlx.Connect("pgx", cfg.PostgresDsn)
	if err != nil {
		lg.Fatal("connect db error", zap.Error(err))
	}
	defer func() {
		err := db.Close()
		if err != nil {
			lg.Fatal("close db error", zap.Error(err))
		}
	}()

	ss := dbstorage.NewSQLXStorage(db, lg)
	gsrv := grpcserver.NewServer(ss, cfg.GRPSListen, lg)
	c := calendar.NewCalendar(ss, hsrv, gsrv, lg)

	c.Start()
}
