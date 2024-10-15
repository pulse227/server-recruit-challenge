package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/pulse227/server-recruit-challenge-sample/api"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(
		os.Stdout,
		&slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	))
	slog.SetDefault(logger)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	r, err := api.NewRouter("root", "root", "localhost:13306", "myapp")
	if err != nil {
		log.Fatalf("new app error: %v", err)
	}

	server := &http.Server{
		Addr:    ":8888",
		Handler: r,
	}
	go func() {
		<-ctx.Done()
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		server.Shutdown(ctx)
	}()
	slog.Info("server start running at :8888")
	log.Fatal(server.ListenAndServe())
}
