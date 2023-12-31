package main

import (
	"Project1/handlers"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	l := log.New(os.Stdout, "Product-api", log.LstdFlags)
	// ph := handlers.NewProducts(l)

	// sm := http.NewServeMux()
	// sm.Handle("/", ph)

	userRepo, _ := handlers.NewRepository()

	sm := http.NewServeMux()
	sm.Handle("/", userRepo)

	//l.Println("Connection done", userRepo)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	C := make(chan os.Signal)
	signal.Notify(C, os.Interrupt)
	signal.Notify(C, os.Kill)

	sig := <-C
	l.Println("Received terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
	//http.ListenAndServe(":9090", sm)
}
