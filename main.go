package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/abrshDev/handlers"
	"github.com/gorilla/mux"
)

// var bindAddress = env.String("BIND_ADDRESSS",false,":1991","bind address for the server")
func main() {

	l := log.New(os.Stdout, "product-api ", log.LstdFlags)

	ph := handlers.NewProduct(l)
	sm := mux.NewRouter()
	getrouter := sm.Methods(http.MethodGet).Subrouter()
	getrouter.HandleFunc("/products", ph.GetProducts)
	putrouter := sm.Methods(http.MethodPut).Subrouter()
	putrouter.Use(ph.MiddleWareValidate)
	putrouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProducts)
	addrouter := sm.Methods(http.MethodPost).Subrouter()
	addrouter.HandleFunc("/", ph.AddProduct)

	s := &http.Server{
		Addr:         ":1991",
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
