package data

import (
	"fmt"
	"time"
)

//	Product defines model of product API
//	swagger:model
type Product struct {
	// id for the product
	//	required: false
	// 	min: 1
	//	pattern: [0-9]+
	Id int `json:"id"` //Unique identifier for the product

	//	name for the product
	//
	//	required: true
	// max length : 255
	Name string `json:"name" validate:"required"`

	// the description for this poduct
	//
	// required: false
	// max length: 10000
	Description string `json:"description"`

	// the price for the product
	//
	// required: true
	// min: 0.01
	Price float32 `json:"price" validate:"gt=0"`

	// the SKU for the product
	//
	// required: true
	// pattern: [a-z]+-[a-z]+-[a-z]+
	SKU       string `json:"sku" validate:"required,sku"`
	CreatedOn string `json:"-"`
	UpdatedOn string `json:"-"`
	DeletedOn string `json:"-"`
}

type Products []*Product

func GetProducts() Products {
	return productList
}

func GetProductById(id int) (*Product, error) {
	i := findIndexByProductId(id)

	if i == -1 {
		return nil, ErrorProductNotFound
	}

	return productList[i], nil
}

func AddProduct(p *Product) {
	//get next sequence id
	p.Id = getNextId()
	productList = append(productList, p)
}

func UpdateProduct(id int, p *Product) error {
	i := findIndexByProductId(id)

	if i == -1 {
		return ErrorProductNotFound
	}

	p.Id = id
	productList[i] = p

	return nil
}

func DeleteProduct(id int) error {
	i := findIndexByProductId(id)

	if i == -1 {
		return ErrorProductNotFound
	}

	productList = append(productList[:i], productList[i+1])

	return nil

}

var ErrorProductNotFound = fmt.Errorf("Product Not Found")

func findIndexByProductId(id int) int {
	for i, p := range productList {
		if p.Id == id {
			return i
		}
	}

	return -1
}

func getNextId() int {
	lp := productList[len(productList)-1]
	return lp.Id + 1
}

var productList = []*Product{
	{
		Id:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "ABC123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
		DeletedOn:   "",
	},
	{
		Id:          2,
		Name:        "Espresso",
		Description: "short and strong coffee without milk",
		Price:       1.99,
		SKU:         "ZXC23",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
		DeletedOn:   "",
	},
}
