package main

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Customer struct {
	ID      int `gorm:"primaryKey"`
	Name    string
	Address string
	Uuid    uuid.UUID
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/gocourse"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Customer{})

	//db.Create(&Customer{
	//	Name:    "Daenerys Targaryen",
	//	Address: "Red keep, kings landing",
	//	Uuid:    uuid.New(),
	//})

	//var customer Customer
	//db.First(&customer, 2)
	//db.First(&customer, "name = ?", "Daenerys Targaryen")
	//db.Where("name LIKE ?", "%Daenerys%").Find(&customer)
	//fmt.Println(customer)

	var customers []Customer
	db.Find(&customers)
	fmt.Println(customers)

}
