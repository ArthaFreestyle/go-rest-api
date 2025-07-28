package main

import (
	"ArthaFreestyle/go-rest-api/app"
	"ArthaFreestyle/go-rest-api/controller"
	"ArthaFreestyle/go-rest-api/middleware"

	"ArthaFreestyle/go-rest-api/repository"
	"ArthaFreestyle/go-rest-api/service"
	"net/http"

	"github.com/go-playground/validator"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository,db,validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr: "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	if err != nil{
		panic(err)
	}
}