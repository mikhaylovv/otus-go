package main

import (
	"flag"
	"github.com/mikhaylovv/otus-go/hw_8/calendar"
	"github.com/mikhaylovv/otus-go/hw_8/calendar/storage/inmemorystorage"
	"github.com/mikhaylovv/otus-go/hw_8/config"
	"github.com/mikhaylovv/otus-go/hw_8/httpserver"
	"go.uber.org/zap"
	"io/ioutil"
	"log"
	"time"
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

	c := calendar.Calendar{
		Storage: &inmemorystorage.InMemoryStorage{},
		Logger:  lg,
	}
	_, _ = c.Storage.GetEvents(time.Now(), time.Now())

	srv := httpserver.NewHTTPServer(lg)

	err = srv.StartListen(cfg.HTTPLiten)
	if err != nil {
		log.Fatal("can't start http server", err)
	}
}
