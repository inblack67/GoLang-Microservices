package models

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	CreatedAt   string  `json:"-"` // omit
}

type TProducts []*Product

var DummyProducts = []*Product{
	{
		ID:          1,
		Name:        "idk",
		Description: "desc",
		Price:       12.0,
		CreatedAt:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "lol",
		Description: "nah",
		Price:       9.0,
		CreatedAt:   time.Now().UTC().String(),
	},
}

func GetProducts() TProducts {
	return DummyProducts
}

// writes directly to the writer (response writer in this case) => not in memory storage => faster
func (p *TProducts) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(p)
}

func (p *Product) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(p)
}

func (p *Product) FromJSON(r io.Reader) error {
	encoder := json.NewDecoder(r)
	return encoder.Decode(p)
}

func (p *Product) GetNextID() int {
	lastID := DummyProducts[len(DummyProducts)-1].ID
	return lastID + 1
}

func (p *Product) AddProduct() {
	DummyProducts = append(DummyProducts, p)
}
