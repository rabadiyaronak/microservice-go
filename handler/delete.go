package handler

import (
	"net/http"

	"github.com/rabadiyaronak/product-api/data"
)

func (p *Product) Delete(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("[DEBUG] Delete Product")

	id := getProductIdFromPathVariable(r)

	p.l.Println("[DEBUG] Delete product with id:", id)

	err := data.DeleteProduct(id)

	if err == data.ErrorProductNotFound {
		p.l.Println("[ERROR] deleting product id doesn't exist")
		rw.WriteHeader(http.StatusNotFound)
		data.ToJson(&GenericError{Message: err.Error()}, rw)
		return
	}

	if err != nil {
		p.l.Println("[ERROR] deleting product", err)
		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJson(&GenericError{Message: err.Error()}, rw)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}
