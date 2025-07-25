package main

import (
	"ArthaFreestyle/go-rest-api/controller"
	"ArthaFreestyle/go-rest-api/app"
	"ArthaFreestyle/go-rest-api/repository"
	"ArthaFreestyle/go-rest-api/service"
	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository,db,validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()

	router.GET("/api/categories",categoryController.FindAll)
	router.GET("/api/categories/:categoryId",categoryController.FindById)
	router.POST("/api/categories",categoryController.Create)
	router.PUT("/api/categories/:categoryId",categoryController.Update)
	router.DELETE("/api/categories/:categoryId",categoryController.Delete)

	server := http.Server{
		Addr: "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil{
		panic(err)
	}
}