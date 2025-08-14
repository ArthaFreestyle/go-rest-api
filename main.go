package main

import (
	"ArthaFreestyle/go-rest-api/app"
	"ArthaFreestyle/go-rest-api/injector"
	"ArthaFreestyle/go-rest-api/middleware"
	"net/http"
)

func main() {
	
	categoryController := injector.InitializedService()

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