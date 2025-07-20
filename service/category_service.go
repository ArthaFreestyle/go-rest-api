package service

import (
	"context"
	"ArthaFreestyle/go-rest-api/model/web"
)


type CategoryService interface {
	Create(ctx context.Context,request web.CategoryCreateRequest) web.CategoryResponse
	Update(ctx context.Context)
	Delete(ctx context.Context)
	FindbyId(ctx context.Context)
	FindAll(ctx context.Context)
}
