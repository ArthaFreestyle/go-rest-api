package exception

import (
	"ArthaFreestyle/go-rest-api/model/web"
	"encoding/json"
	"net/http"
	"github.com/go-playground/validator"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, error interface{}) {
	if notFoundError(w, r, error) {
		return
	}
	if ValidationError(w,r,error){
		return
	}
	InternalServerError(w, r, error)
}

func notFoundError(w http.ResponseWriter, r *http.Request, error interface{}) bool {
	exception, ok := error.(NotFoundError)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   exception.Error,
		}

		encoder := json.NewEncoder(w)
		err := encoder.Encode(webResponse)
		if err != nil {
			panic(err)
		}
		return true
	} else {
		return false
	}
}

func ValidationError(w http.ResponseWriter, r *http.Request, error interface{}) bool{
	exception, ok := error.(validator.ValidationErrors)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   exception.Error(),
		}

		encoder := json.NewEncoder(w)
		err := encoder.Encode(webResponse)
		if err != nil {
			panic(err)
		}
		return true
	} else {
		return false
	}
}

func InternalServerError(w http.ResponseWriter, r *http.Request, error interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   error,
	}

	encoder := json.NewEncoder(w)
	err := encoder.Encode(webResponse)
	if err != nil {
		panic(err)
	}
}
