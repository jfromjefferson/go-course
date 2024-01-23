package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

	tx := db.Begin()
	var category Category
	err = tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&category).Error
	if err != nil {
		panic(err)
	}

	category.Name = "Teste"
	tx.Debug().Save(&category)
	tx.Commit()

}
