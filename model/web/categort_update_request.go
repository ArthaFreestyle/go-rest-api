package web

type CategoryUpdateRequest struct{
	Id int64 `json:"Name" validate:"required"`
	Name string `json:"-" validate:"required,max=200,min=1"`
}