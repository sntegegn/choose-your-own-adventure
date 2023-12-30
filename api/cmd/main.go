package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/go-playground/form/v4"
)

type config struct {
	jsonPath string
	addr     string
}

type application struct {
	cfg         config
	logger      slog.Logger
	formdecoder form.Decoder
	gopherStory story
}

func main() {
	var cfg config
	flag.StringVar(&cfg.jsonPath, "jsonPath", "gopher.json", "the path to the json file")
	flag.StringVar(&cfg.addr, "addr", ":4002", "the server address")

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	app := &application{
		cfg:         cfg,
		logger:      *logger,
		formdecoder: *form.NewDecoder(),
	}

	srv := http.Server{
		Addr:         cfg.addr,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	app.logger.Info("Starting server at address ", "Addr", cfg.addr)

	err := srv.ListenAndServe()
	app.logger.Error(err.Error())
	os.Exit(0)
}
