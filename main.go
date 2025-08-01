package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"github.com/abrshDev/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api ", log.LstdFlags)


  ph := handlers.NewProduct(l)
	sm := http.NewServeMux()

	
	sm.Handle("/",ph)
	s := &http.Server{
		Addr:         ":1990",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	go func() {
		log.Fatal(s.ListenAndServe())
	}()
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)
	sig := <-sigChan
	l.Println("received terminate, graceful shutdown", sig)
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)

}
