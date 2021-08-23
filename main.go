package main

import (
	"context"
	"default/handlers"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Server starting...")

	ph := handlers.Products()

	serverMux := mux.NewRouter()

	getRouter := serverMux.Methods("GET").Subrouter()
	postRouter := serverMux.Methods("POST").Subrouter()

	getRouter.HandleFunc("/products", ph.GetProducts)
	postRouter.HandleFunc("/products", ph.AddProduct)

	server := &http.Server{
		Handler:      serverMux,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
		IdleTimeout:  120 * time.Second,
		Addr:         ":5000",
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	signalChannel := make(chan os.Signal)

	signal.Notify(signalChannel, os.Interrupt)
	signal.Notify(signalChannel, os.Kill)

	signalChannelRes := <-signalChannel

	fmt.Println("Recieved termination request, shutting down: reason =>", signalChannelRes)

	timerContext, cancel := context.WithTimeout(context.Background(), 120*time.Second)

	defer cancel()

	err := server.Shutdown(timerContext)

	if err != nil {
		fmt.Println("Error shutting down")
		log.Fatal(err)
	}
}
