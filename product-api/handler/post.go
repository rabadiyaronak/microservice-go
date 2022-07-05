package handler

import (
	"net/http"

	"github.com/rabadiyaronak/microservice-go/product-api/data"
)

// swagger:route POST /products products createProduct
// Create a new product
//
// responses:
//	200: productResponse
//	400: errorResponse
//  422: errorValidation
//  501: errorResponse

// Create handles POST requests to add new products
func (p *Product) Create(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("[DEBUG] Creat product")
	product := r.Context().Value(KeyProduct{}).(data.Product)

	data.AddProduct(&product)
}
