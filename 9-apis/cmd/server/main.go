package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/jfromjefferson/gi-course-9/configs"
	_ "github.com/jfromjefferson/gi-course-9/docs"
	"github.com/jfromjefferson/gi-course-9/internal/entity"
	"github.com/jfromjefferson/gi-course-9/internal/infra/database"
	"github.com/jfromjefferson/gi-course-9/internal/infra/webserver/handlers"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
)

// @title Go API
// @version 1.0
// @description Product API with authentication
// @termsOfService https://swagger.io/terms/

// @contact.name Jefferson Silva
// @contact.url https://magenta-fox-5d5416.netlify.app/
// @contact.email jeffsilva1@outlook.com

// @host localhost:8000
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(entity.Product{}, entity.User{})

	productDB := database.NewProduct(db)
	userDB := database.NewUser(db)

	productHandler := handlers.NewProductHandler(productDB)
	userHandler := handlers.NewUserHandler(userDB, config.TokenAuth, config.JWTExpiresIn)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	//router.Use(LogRequest)

	router.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(config.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", productHandler.CreateProduct)
		r.Get("/{id}", productHandler.GetProduct)
		r.Get("/", productHandler.GetProducts)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	router.Post("/auth", userHandler.GetJWT)
	router.Post("/users", userHandler.Create)

	router.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))

	fmt.Println("Server running on port 8000")
	err = http.ListenAndServe(":8000", router)
	if err != nil {
		panic(err)
	}
}

// LogRequest middleware
func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
