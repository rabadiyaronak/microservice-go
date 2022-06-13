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
			rw.WriteHeader(http.StatusBadRequest)
			data.ToJson(&GenericError{Message: err.Error()}, rw)
			return
		}

		//validate the product
		errs := p.v.Validate(product)

		if len(errs) != 0 {
			p.l.Println(" [ERROR] deserialization error", err)

			rw.WriteHeader(http.StatusUnprocessableEntity)
			data.ToJson(&ValidationError{Message: errs.Errors()}, rw)
			return
		}

		//add product to the context
		ctx := context.WithValue(r.Context(), KeyProduct{}, product)

		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}
