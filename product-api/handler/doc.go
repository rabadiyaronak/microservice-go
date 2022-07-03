package handler

import "github.com/rabadiyaronak/microserive-go/product-api/data"

//
// NOTE: Types defined here are purely for documentation purposes
// these types are not used by any of the handers

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// Validation errors defined as an array of strings
// swagger:response errorValidation
type errorValidationWrapper struct {
	// Collection of the errors
	// in: body
	Body ValidationError
}

// A list of products returns in the response
//	swagger:response productsResponse
type productsResponseWrapper struct {

	// All products in the systems
	//	in:body
	Body []data.Product
}

// Wrapper ds to represent a single product
//	swagger:response productResponse
type productResponseWrapper struct {

	//	All products in the systems
	//	in: body
	Body data.Product
}

//	swagger:parameters getProductById updateProduct deleteProduct
type productIdPathParam struct {
	//	in: path
	//	required: true
	// 	min: 1
	//	pattern: [0-9]+
	ID int `json:"id"`
}

// No content is returned by this API endpoint
// swagger:response noContentResponse
type noContentResponseWrapper struct {
}

// swagger:parameters updateProduct createProduct
type productParamsWrapper struct {
	// Product data structure to Update or Create.
	// Note: the id field is ignored by update and create operations
	// in: body
	// required: true
	Body data.Product
}
