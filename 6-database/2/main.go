package main

import (
	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	ID      int `gorm:"primaryKey"`
	Name    string
	Address string
	Uuid    uuid.UUID
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/gocourse?charset=utf8&parseTime=True&loc=Local"
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

	//var customers []Customer
	//db.Find(&customers)
	//fmt.Println(customers)

	var customer Customer
	db.First(&customer, 1)
	//customer.Name = "Aegon Targaryen"
	//db.Save(&customer)

	//fmt.Println(customer)

	db.Delete(&customer)

}
