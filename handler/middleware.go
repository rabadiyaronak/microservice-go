package handler

import (
	"context"
	"net/http"

	"github.com/rabadiyaronak/product-api/data"
)

func (p *Product) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		product := data.Product{}

		err := data.FromJSON(product, r.Body)

		if err != nil {
			p.l.Println(" [ERROR] deserialization error", err)
			http.Error(rw, "error reading product", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, product)

		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}
