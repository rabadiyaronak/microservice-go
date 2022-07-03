package data

import "testing"

func TestValidateProduct(t *testing.T) {
	p := &Product{
		Name:  "Test Product",
		Price: 10.00,
		SKU:   "abcd-abcd-abcd",
	}
	v := NewValidation()
	err := v.Validate(p)

	if err != nil {
		t.Fatal(err)
	}
}
