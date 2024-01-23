package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name     string
	Products []Product
}

type Product struct {
	gorm.Model
	Name       string
	Price      int
	CategoryID uint
	Category   Category
	Serial     Serial
}

type Serial struct {
	gorm.Model
	Number    int
	ProductID uint
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/gocourse?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}, &Category{}, &Serial{})

	//category := Category{
	//	Name: "Cozinha",
	//}
	//db.Create(&category)
	//
	//product := Product{
	//	Name:       "Liquidificador",
	//	CategoryID: category.ID,
	//	Price:      1500,
	//}
	//
	//db.Create(&product)
	//
	//serial := Serial{
	//	Number:    123456,
	//	ProductID: product.ID,
	//}
	//
	//db.Create(&serial)

	var categories []Category
	err = db.Model(&Category{}).Preload("Products.Serial").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		for _, product := range category.Products {
			fmt.Printf("Product: %s - Category: %s - Serial: %d\n", product.Name, category.Name, product.Serial.Number)
		}
	}

}
