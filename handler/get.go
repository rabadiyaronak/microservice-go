package handler

import (
	"net/http"

	"github.com/rabadiyaronak/product-api/data"
)

func (p *Product) ListAll(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("[DEBUG] Get All products")

	products := data.GetProducts()

	err := data.ToJson(products, rw)

	if err != nil {
		p.l.Println("[ERROR] serializing product", err)
	}

}

func (p *Product) GetProductById(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("[DEBUG] Get Product by id")

	id := getProductIdFromPathVariable(r)

	p.l.Println("[DEBUG] Get Product id:", id)

	prod, err := data.GetProductById(id)

	switch err {
	case nil:
	case data.ErrorProductNotFound:
		p.l.Println("[ERROR] fetching product", err)

		rw.WriteHeader(http.StatusNotFound)
		data.ToJson(&GenericError{Message: err.Error()}, rw)
		return
	default:
		p.l.Println("[ERROR] fetching product", err)

		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJson(&GenericError{Message: err.Error()}, rw)
		return
	}
	err = data.ToJson(prod, rw)
	if err != nil {
		// we should never be here but log the error just incase
		p.l.Println("[ERROR] serializing product", err)
	}

}
