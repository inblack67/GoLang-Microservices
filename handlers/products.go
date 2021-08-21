package handlers

import (
	"default/models"
	"net/http"
)

type HProducts struct{}

func Products() *HProducts {
	return &HProducts{}
}

func (hp *HProducts) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		hp.getProducts(rw, r)
		return
	}
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *HProducts) getProducts(rw http.ResponseWriter, r *http.Request) {
	productsList := models.GetProducts()
	// jsonData, err := json.Marshal(productsList) // consumes memory

	err := productsList.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// rw.Write(jsonData)
}
