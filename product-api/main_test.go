package main

import (
	"fmt"
	"testing"

	"github.com/rabadiyaronak/microserive-go/product-api/sdk/client"
	"github.com/rabadiyaronak/microserive-go/product-api/sdk/client/products"
	"github.com/stretchr/testify/assert"
)

func TestProductClient(t *testing.T) {

	c := client.Default
	params := products.NewListProductsParams()
	resp, err := c.Products.ListProducts(params)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%v", len(resp.GetPayload()))

	assert.NotEmpty(t, resp.GetPayload(), "API should return with some values")

}
