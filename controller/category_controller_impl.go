package controller

import (
	"ArthaFreestyle/go-rest-api/service"
	"net/http"
	"encoding/json"
	"ArthaFreestyle/go-rest-api/model/web"
	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func (controller *CategoryControllerImpl) Create(w http.ResponseWriter, r *http.Request,params httprouter.Params){
	decoder := json.NewDecoder(r.Body)
	categoryCreateRequest := web.CategoryCreateRequest{}
	err := decoder.Decode(&categoryCreateRequest)
	if err != nil {
		panic(err)
	}
	categoryRespone := controller.CategoryService.Create(r.Context(), categoryCreateRequest)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: categoryRespone,
	}

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(webResponse)
	if err != nil {
		panic(err)
	}
}

func (service *CategoryControllerImpl) Update(w http.ResponseWriter, r *http.Request,params httprouter.Params){

}

func (service *CategoryControllerImpl) Delete(w http.ResponseWriter, r *http.Request,params httprouter.Params){

}

func (service *CategoryControllerImpl) FindById(w http.ResponseWriter, r *http.Request,params httprouter.Params){

}

func (service *CategoryControllerImpl) FindAll(w http.ResponseWriter, r *http.Request,params httprouter.Params){

}
