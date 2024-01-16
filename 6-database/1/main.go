package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"log"
)

type Product struct {
	ID    int
	Name  string
	Price float64
	Uuid  string
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		Name:  name,
		Price: price,
		Uuid:  uuid.New().String(),
	}
}

func create(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("INSERT INTO products(name, price, uuid) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(product.Name, product.Price, product.Uuid)
	if err != nil {
		return err
	}

	fmt.Println("Created")
	return nil
}

func getAll(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []Product

	for rows.Next() {
		var p Product

		err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Uuid)
		if err != nil {
			log.Fatal(err)
		}

		products = append(products, p)
	}

	return products, nil
}

func getOne(db *sql.DB, uuid string) (*Product, error) {
	stmt, err := db.Prepare("SELECT * FROM products WHERE uuid = ?")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var p Product

	err = stmt.QueryRow(uuid).Scan(&p.ID, &p.Name, &p.Price, &p.Uuid)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func update(db *sql.DB, product Product) error {
	stmt, err := db.Prepare("UPDATE products SET name = ?, price = ? WHERE uuid = ?")
	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(product.Name, product.Price, product.Uuid)
	if err != nil {
		return err
	}

	fmt.Println("Updated")
	return nil
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/gocourse")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	product, err := getOne(db, "642d6d7a-bd9a-4e5e-a659-a3d3554c2dd7")
	if err != nil {
		panic(err)
	}

	product.Name = "Caneca azul"

	fmt.Println(product)

	products, err := getAll(db)
	if err != nil {
		panic(err)
	}

	fmt.Println(products)

	err = update(db, *product)
	if err != nil {
		panic(err)
	}

	/*product := NewProduct("Bala de goma", 9.99)
	err = create(db, product)
	if err != nil {
		panic(err)
	}*/
}
