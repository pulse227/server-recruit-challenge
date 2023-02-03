package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Lupusdog/server-recruit-challenge-2024/api"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	r := api.NewRouter()

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
	log.Println("server start running at :8888")
	log.Fatal(server.ListenAndServe())
}
