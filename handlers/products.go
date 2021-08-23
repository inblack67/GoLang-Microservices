package handlers

import (
	"default/models"
	"net/http"
	"time"
)

type HProducts struct{}

func Products() *HProducts {
	return &HProducts{}
}

func (p *HProducts) GetProducts(rw http.ResponseWriter, r *http.Request) {
	productsList := models.GetProducts()
	// jsonData, err := json.Marshal(productsList) // consumes memory

	err := productsList.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// rw.Write(jsonData)
}

func (p *HProducts) AddProduct(rw http.ResponseWriter, r *http.Request) {
	newProduct := &models.Product{}
	err := newProduct.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Internal Server Error", http.StatusBadRequest)
		return
	}
	newProduct.ID = newProduct.GetNextID()
	newProduct.CreatedAt = time.Now().String()
	newProduct.AddProduct()
	newProduct.ToJSON(rw)
}
