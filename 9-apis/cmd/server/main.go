package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jfromjefferson/gi-course-9/configs"
	"github.com/jfromjefferson/gi-course-9/internal/entity"
	"github.com/jfromjefferson/gi-course-9/internal/infra/database"
	"github.com/jfromjefferson/gi-course-9/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(entity.Product{}, entity.User{})

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Post("/products", productHandler.CreateProduct)
	router.Get("/product/{id}", productHandler.GetProduct)
	router.Get("/products", productHandler.GetProducts)
	router.Put("/product/{id}", productHandler.UpdateProduct)
	router.Delete("/product/{id}", productHandler.DeleteProduct)

	fmt.Println("Server running on port 8000")
	err = http.ListenAndServe(":8000", router)
	if err != nil {
		panic(err)
	}

}
