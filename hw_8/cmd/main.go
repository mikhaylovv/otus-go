package main

import (
	"flag"
	"github.com/mikhaylovv/otus-go/hw_8/calendar"
	"github.com/mikhaylovv/otus-go/hw_8/calendar/storage/inmemorystorage"
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

	s := inmemorystorage.NewInMemoryStorage()
	hsrv := httpserver.NewHTTPServer(cfg.HTTPListen, lg)
	gsrv := grpcserver.NewServer(s, cfg.GRPSListen, lg)
	c := calendar.NewCalendar(s, hsrv, gsrv, lg)

	c.Start()
}
