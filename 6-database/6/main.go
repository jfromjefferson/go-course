package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name     string
	Products []Product `gorm:"many2many:products_categories;"`
}

type Product struct {
	gorm.Model
	Name       string
	Price      int
	Categories []Category `gorm:"many2many:products_categories;"`
}

// Many to Many

func main() {
	dsn := "root:root@tcp(localhost:3306)/gocourse?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}, &Category{})

	//categories := []Category{
	//	0: {
	//		Name: "Cozinha",
	//	},
	//	1: {
	//		Name: "N.A",
	//	},
	//}
	//
	//product := Product{
	//	Name:       "Liquidificador",
	//	Price:      1500,
	//	Categories: categories,
	//}
	//
	//db.Create(&product)

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		for _, product := range category.Products {
			fmt.Printf("Product: %s - Category: %s\n", product.Name, category.Name)
		}
	}

}
