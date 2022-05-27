package handler

import (
	"essential-practices/project/product-api/data"
	"log"
	"net/http"
)

type Product struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Product {
	return &Product{l}
}

func (p *Product) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	//default catch for not supported method
	rw.WriteHeader(http.StatusMethodNotAllowed)

}

func (p *Product) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")

	lp := data.GetProducts()

	err := lp.ToJson(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
