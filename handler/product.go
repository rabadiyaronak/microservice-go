package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rabadiyaronak/product-api/data"
)

type KeyProduct struct{}

type Product struct {
	l *log.Logger
	v *data.Validation
}

func NewProduct(l *log.Logger, v *data.Validation) *Product {
	return &Product{l, v}
}

var ErrInvalidProductPath = fmt.Errorf("Invalid Path, path shoudl be /product/[id]")

type GenericError struct {
	Message string `json: message`
}

type ValidationError struct {
	Message []string `json: message`
}

func getProductIdFromPathVariable(r *http.Request) int {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		panic(err)
	}
	return id

}
