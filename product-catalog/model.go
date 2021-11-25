package main

var schema = `
CREATE TABLE products (
    id text,
    product_name text
);
`

type Product struct {
	ID          string `json:"id" db:"id"`
	ProductName string `json:"productName" db:"product_name"`
}

type Products []Product
