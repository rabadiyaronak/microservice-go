package data

import "testing"

func TestValidateProduct(t *testing.T) {
	p := &Product{
		Name:  "Test Product",
		Price: 10.00,
		SKU:   "abcd-abcd-abcd",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}

}
