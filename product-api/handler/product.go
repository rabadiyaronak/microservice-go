// Package classification  Prodcut Api
//
// Documantation for Product Api
//
// 	Schemes: http
// 	Host: localhost:9090
// 	BasePath: /
// 	Version: 1.0.0
//
// 	Consumes:
// 	- application/json
//
//	Produces:
// 	- application/json
//	swagger:meta
package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rabadiyaronak/microserive-go/product-api/data"
)

type KeyProduct struct{}

type Product struct {
	l *log.Logger
	v *data.Validation
}

func NewProduct(l *log.Logger, v *data.Validation) *Product {
	return &Product{l, v}
}

var ErrInvalidProductPath = fmt.Errorf("invalid Path, path shoudl be /product/[id]")

type GenericError struct {
	Message string `json:"message"`
}

type ValidationError struct {
	Message []string `json:"message"`
}

func getProductIdFromPathVariable(r *http.Request) int {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		panic(err)
	}
	return id

}
