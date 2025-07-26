package controller

import (
	"ArthaFreestyle/go-rest-api/model/web"
	"ArthaFreestyle/go-rest-api/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(CategoryService service.CategoryService) CategoryController{
	return &CategoryControllerImpl{
		CategoryService: CategoryService,
	}
}

func (controller *CategoryControllerImpl) Create(w http.ResponseWriter, r *http.Request,params httprouter.Params){
	decoder := json.NewDecoder(r.Body)
	categoryCreateRequest := web.CategoryCreateRequest{}
	err := decoder.Decode(&categoryCreateRequest)
	if err != nil {
		panic(err)
	}
	categoryResponse := controller.CategoryService.Create(r.Context(), categoryCreateRequest)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: categoryResponse,
	}

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(webResponse)
	if err != nil {
		panic(err)
	}
}

func (controller *CategoryControllerImpl) Update(w http.ResponseWriter, r *http.Request,params httprouter.Params){
	decoder := json.NewDecoder(r.Body)
	categoryUpdateRequest := web.CategoryUpdateRequest{}
	err := decoder.Decode(&categoryUpdateRequest)
	if err != nil {
		panic(err)
	}
	categoryId := params.ByName("categoryId")
	id,err := strconv.ParseInt(categoryId,10,64)
	if err != nil{
		panic(err)
	}
	categoryUpdateRequest.Id = id
	categoryResponse := controller.CategoryService.Update(r.Context(), categoryUpdateRequest)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: categoryResponse,
	}
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(webResponse)
	if err != nil {
		panic(err)
	}
}

func (controller *CategoryControllerImpl) Delete(w http.ResponseWriter, r *http.Request,params httprouter.Params){
	categoryId := params.ByName("categoryId")
	id,err := strconv.ParseInt(categoryId,10,64)
	if err != nil {
		panic(err)
	}
	controller.CategoryService.Delete(r.Context(),id)
	webResponse := web.WebResponse{
		Code : 200,
		Status: "OK",
	}
	w.Header().Set("Content-Type","application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(webResponse)
	if err != nil {
		panic(err)
	}

}

func (controller *CategoryControllerImpl) FindById(w http.ResponseWriter, r *http.Request,params httprouter.Params){
	categoryId := params.ByName("categoryId")
	id,err := strconv.ParseInt(categoryId,10,64)
	if err != nil {
		panic(err)
	}
	categoryResponse := controller.CategoryService.FindById(r.Context(),id)
	webResponse := web.WebResponse{
		Code:200,
		Status: "OK",
		Data:categoryResponse,
	}
	w.Header().Set("Content-Type","application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(webResponse)
	if err != nil {
		panic(err)
	}
}

func (controller *CategoryControllerImpl) FindAll(w http.ResponseWriter, r *http.Request,params httprouter.Params){
	categoryResponse := controller.CategoryService.FindAll(r.Context())
	webResponse := web.WebResponse{
		Code: 200,
		Status:"OK",
		Data: categoryResponse,
	}
	encoder := json.NewEncoder(w)
	err := encoder.Encode(webResponse)
	if err!=nil{
		panic(err)
	}
}
