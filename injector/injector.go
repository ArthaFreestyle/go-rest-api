//go:build wireinject
// +build wireinject

package injector

import (
	"ArthaFreestyle/go-rest-api/app"
	"ArthaFreestyle/go-rest-api/controller"
	"ArthaFreestyle/go-rest-api/repository"
	"ArthaFreestyle/go-rest-api/service"

	"github.com/go-playground/validator"
	"github.com/google/wire"
)


func InitializedService() controller.CategoryController{
	wire.Build(
		app.NewDB,
		validator.New,
		repository.NewCategoryRepository,
		service.NewCategoryService,
		controller.NewCategoryController,
	)
	return nil
}