package repository

import (
	"ArthaFreestyle/go-rest-api/model/domain"
	"context"
	"database/sql"
)

type CategoryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Delete(ctx context.Context, tx *sql.Tx, categoryId int64)
	FindById(ctx context.Context, tx *sql.Tx, categoryId int64) (domain.Category,error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Category
}