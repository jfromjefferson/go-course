package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name string
}

type Product struct {
	gorm.Model
	Name       string
	Price      int
	CategoryID int
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
	//	Name: "Eletrônicos",
	//}
	//
	//db.Create(&category)

	//product := Product{
	//	Name:     "Smartphone Samsung A54 128g",
	//	Price:    1500000,
	//	Category: category,
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

	var products []Product
	db.Preload("Category").Preload("Serial").Find(&products)

	for _, product := range products {
		fmt.Println(product.Name, "-", product.Category.Name, "-", product.Serial.Number)
	}

}
