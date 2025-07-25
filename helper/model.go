package helper

import (
	"ArthaFreestyle/go-rest-api/model/domain"
	"ArthaFreestyle/go-rest-api/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id: category.Id,
		Name: category.Name,
	}
}