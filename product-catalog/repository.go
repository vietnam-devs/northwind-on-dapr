package main

import (
	"github.com/jmoiron/sqlx"
)

func initDb(db *sqlx.DB) {
	_, table_check := db.Query("select * from products;")

	if table_check != nil {
		db.MustExec(schema)

		tx := db.MustBegin()
		tx.MustExec("INSERT INTO products (id, product_name) VALUES ($1, $2)", "025f55c5-7f97-44f9-ae58-c57239bcbe16", "product 1")
		tx.MustExec("INSERT INTO products (id, product_name) VALUES ($1, $2)", "b44769be-3353-4bf5-b397-bbc1af071bf1", "product 2")
		tx.Commit()
	}
}

func (p *Products) getProducts(db *sqlx.DB) error {
	return db.Select(p, "SELECT * FROM products ORDER BY product_name ASC")
}

func (p *Product) createProduct(db *sqlx.DB) error {
	err := db.QueryRow(
		"INSERT INTO products(id, product_name) VALUES($1, $2) RETURNING id",
		p.ID, p.ProductName).Scan(&p.ID)

	if err != nil {
		return err
	}

	return nil
}

func (p *Product) updateProduct(db *sqlx.DB) error {
	_, err := db.Exec("UPDATE products SET product_name=$1 WHERE id=$2", p.ProductName, p.ID)
	return err
}

func (p *Product) deleteProduct(db *sqlx.DB) error {
	_, err := db.Exec("DELETE FROM products WHERE id=$1", p.ID)
	return err
}