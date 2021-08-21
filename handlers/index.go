package handlers

import (
	"fmt"
	"io"
	"net/http"
)

type HRoot struct{}

func Root() *HRoot {
	return &HRoot{}
}

func (h *HRoot) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)

	if err != nil {

		http.Error(rw, "idk", http.StatusInternalServerError)

		// fmt.Println("error in root path")
		// rw.WriteHeader(http.StatusInternalServerError)
		// return
	}

	fmt.Println("data => ", data)
	rw.Write([]byte(data))
}
