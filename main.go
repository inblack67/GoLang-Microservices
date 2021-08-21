package main

import (
	"default/handlers"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server starting...")

	rh := handlers.Root()

	serverMux := http.NewServeMux()

	serverMux.Handle("/", rh)

	http.ListenAndServe(":5000", serverMux) // bind this port to every ip => or can explicitly say the loopback address
}
