package data

import (
	"fmt"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
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
	p.ID = getNextId()
	productList = append(productList, p)
}

func UpdateProduct(id int, p *Product) error {
	i := findIndexByProductId(id)

	if i == -1 {
		return ErrorProductNotFound
	}

	p.ID = id
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
		if p.ID == id {
			return i
		}
	}

	return -1
}

func getNextId() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

var productList = []*Product{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "ABC123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
		DeletedOn:   "",
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "short and strong coffee without milk",
		Price:       1.99,
		SKU:         "ZXC23",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
		DeletedOn:   "",
	},
}
