package handler

import (
	"essential-practices/project/product-api/data"
	"log"
	"net/http"
	"regexp"
	"strconv"
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

	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}
	if r.Method == http.MethodPut {
		p.l.Println("PUT", r.URL.Path)

		regexpIdFinder := regexp.MustCompile(`/([0-9]+)`)
		g := regexpIdFinder.FindAllStringSubmatch(r.URL.Path, -1)

		if len(g) != 1 {
			p.l.Println("Invalid URI more than one id")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		if len(g[0]) != 2 {
			p.l.Println("Invalid URI capture group should be two")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			p.l.Println("Invalid URI unable to convert id to number")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		p.l.Println("extracted id ", id)

		p.updateProduct(id, rw, r)

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

func (p *Product) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	product := &data.Product{}

	err := product.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "Unable to unmarshal request", http.StatusBadRequest)
	}

	data.AddProduct(product)
}

func (p *Product) updateProduct(id int, rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT Product")

	product := &data.Product{}

	err := product.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "Unable to unmarshal request", http.StatusBadRequest)
	}

	err = data.UpdateProduct(id, product)

	if err == data.ErrorProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

}
