// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/rabadiyaronak/microserive-go/product-api/sdk/models"
)

// GetProductByIDReader is a Reader for the GetProductByID structure.
type GetProductByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProductByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProductByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetProductByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 501:
		result := NewGetProductByIDNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetProductByIDOK creates a GetProductByIDOK with default headers values
func NewGetProductByIDOK() *GetProductByIDOK {
	return &GetProductByIDOK{}
}

/* GetProductByIDOK describes a response with status code 200, with default header values.

Wrapper ds to represent a single product
*/
type GetProductByIDOK struct {
	Payload *models.Product
}

func (o *GetProductByIDOK) Error() string {
	return fmt.Sprintf("[GET /products/{id}][%d] getProductByIdOK  %+v", 200, o.Payload)
}
func (o *GetProductByIDOK) GetPayload() *models.Product {
	return o.Payload
}

func (o *GetProductByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Product)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProductByIDNotFound creates a GetProductByIDNotFound with default headers values
func NewGetProductByIDNotFound() *GetProductByIDNotFound {
	return &GetProductByIDNotFound{}
}

/* GetProductByIDNotFound describes a response with status code 404, with default header values.

Generic error message returned as a string
*/
type GetProductByIDNotFound struct {
	Payload *models.GenericError
}

func (o *GetProductByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /products/{id}][%d] getProductByIdNotFound  %+v", 404, o.Payload)
}
func (o *GetProductByIDNotFound) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *GetProductByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProductByIDNotImplemented creates a GetProductByIDNotImplemented with default headers values
func NewGetProductByIDNotImplemented() *GetProductByIDNotImplemented {
	return &GetProductByIDNotImplemented{}
}

/* GetProductByIDNotImplemented describes a response with status code 501, with default header values.

Generic error message returned as a string
*/
type GetProductByIDNotImplemented struct {
	Payload *models.GenericError
}

func (o *GetProductByIDNotImplemented) Error() string {
	return fmt.Sprintf("[GET /products/{id}][%d] getProductByIdNotImplemented  %+v", 501, o.Payload)
}
func (o *GetProductByIDNotImplemented) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *GetProductByIDNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
