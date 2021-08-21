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
)

func main() {
	fmt.Println("Server starting...")

	rh := handlers.Root()
	ph := handlers.Products()

	serverMux := http.NewServeMux()

	serverMux.Handle("/", rh)
	serverMux.Handle("/products", ph)

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
