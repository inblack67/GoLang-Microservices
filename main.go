package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Server starting...")

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		data, err := io.ReadAll(r.Body)

		if err != nil {

			http.Error(rw, "idk", http.StatusInternalServerError)

			// fmt.Println("error in root path")
			// rw.WriteHeader(http.StatusInternalServerError)
			// return
		}

		fmt.Println("data => ", data)
		rw.Write([]byte(data))
	})

	http.ListenAndServe(":5000", nil) // bind this port to every ip => or can explicitly say the loopback address
}
