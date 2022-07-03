package handler

import (
	"net/http"

	"github.com/rabadiyaronak/microserive-go/product-api/data"
)

// swagger:route PUT /products/{id} products updateProduct
// Update a products details
//
// responses:
//	201: noContentResponse
//  404: errorResponse
//  422: errorValidation
//  501: errorResponse

// Update handles PUT requests to update products
func (p *Product) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("[DEBUG] Update product")

	id := getProductIdFromPathVariable(r)

	p.l.Println("[DEBUG] updating product with id:", id)

	product := r.Context().Value(KeyProduct{}).(data.Product)

	err := data.UpdateProduct(id, &product)

	if err == data.ErrorProductNotFound {
		p.l.Println("Product not found")
		rw.WriteHeader(http.StatusNotFound)
		data.ToJson(&GenericError{Message: err.Error()}, rw)
		return
	}

	if err != nil {
		p.l.Println("[ERROR] updating product", err)
		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJson(&GenericError{Message: err.Error()}, rw)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}
