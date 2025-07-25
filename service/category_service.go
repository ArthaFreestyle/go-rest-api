package service

import (
	"context"
	"ArthaFreestyle/go-rest-api/model/web"
)


type CategoryService interface {
	Create(ctx context.Context,request web.CategoryCreateRequest) web.CategoryResponse
	Update(ctx context.Context,request web.CategoryUpdateRequest) web.CategoryResponse
	Delete(ctx context.Context,categoryId int64)
	FindById(ctx context.Context,categoryId int64) web.CategoryResponse
	FindAll(ctx context.Context) []web.CategoryResponse
}
