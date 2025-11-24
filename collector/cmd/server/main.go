package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi/v5"
    "github.com/Vishalrana7/observalite/collector/handlers"
    "github.com/Vishalrana7/observalite/collector/db"
)


func main() {
	cfg := db.LoadConfigFromEnv()
	conn, err := db.Connect(cfg)
	if err != nil {log.Fatal(err)}

	r := chi.NewRouter()
	r.Post("/v1/events", handlers.PostEventHandler(conn))

	srv := &http.Server{
		Addr: ":8080",
		Handler: r,
	}

	go func ()  {
		log.Println("Collector listening on :8080")
		if err := srv.ListenAndServe(); err!= nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()


	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	srv.Shutdown(ctx)


}