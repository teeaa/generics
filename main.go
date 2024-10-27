package main

import (
	"context"
	"log/slog"
	"os/signal"
	"syscall"

	"github.com/teeaa/generics/internal/api"
	"github.com/teeaa/generics/internal/config"
	"github.com/teeaa/generics/internal/database"
	"github.com/teeaa/generics/internal/service"
)

func main() {
	conf, err := config.New()
	if err != nil {
		panic(err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM)
	defer stop()

	db, err := database.New(conf)
	if err != nil {
		panic(err)
	}

	srv := service.New(conf, db)

	api, err := api.New(conf, srv)
	if err != nil {
		panic(err)
	}

	<-ctx.Done()
	err = api.Stop()
	if err != nil {
		slog.Error("Error stopping API server", slog.Any("err", err))
	}

	err = db.Close()
	if err != nil {
		slog.Error("Error stopping API server", slog.Any("err", err))
	}
}
