package handler

import (
	"net/http"

	"github.com/rabadiyaronak/product-api/data"
)

func (p *Product) Create(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("[DEBUG] Creat product")
	product := r.Context().Value(KeyProduct{}).(data.Product)

	data.AddProduct(&product)
}
