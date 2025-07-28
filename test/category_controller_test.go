package test

import (
	"ArthaFreestyle/go-rest-api/app"
	"ArthaFreestyle/go-rest-api/controller"
	"ArthaFreestyle/go-rest-api/middleware"
	"ArthaFreestyle/go-rest-api/repository"
	"ArthaFreestyle/go-rest-api/service"
	"context"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
	"ArthaFreestyle/go-rest-api/model/domain"
	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"strconv"
)

func SetupNewDB() *sql.DB{
	db,err := sql.Open("mysql","root@tcp(localhost:3306)/go_restful_api_test")
	if err != nil{
		panic(err)
	}
	
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(1 * time.Hour)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db

}

func SetupRouter(db *sql.DB) http.Handler{
	
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository,db,validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	return middleware.NewAuthMiddleware(router)
}

func truncateCategory(db *sql.DB){
	db.Exec("TRUNCATE category")
}

func TestCreateCategorySuccess(t *testing.T){
	db := SetupNewDB()
	truncateCategory(db)
	router := SetupRouter(db)

	requestBody := strings.NewReader(`{"name":"Gadget"}`)
	request := httptest.NewRequest(http.MethodPost,"http://localhost:3000/api/categories",requestBody)
	request.Header.Add("X-API-Key","RAHASIA")
	request.Header.Add("Content-Type","application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder,request)

	response := recorder.Result()

	assert.Equal(t,200,response.StatusCode)
}

func TestCreateCategoryFailed(t *testing.T){
	db := SetupNewDB()
	truncateCategory(db)
	router := SetupRouter(db)

	requestBody := strings.NewReader(`{"name":""}`)
	request := httptest.NewRequest(http.MethodPost,"http://localhost:3000/api/categories",requestBody)
	request.Header.Add("X-API-Key","RAHASIA")
	request.Header.Add("Content-Type","application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder,request)

	response := recorder.Result()

	assert.Equal(t,400,response.StatusCode)
}

func TestUpdateCategorySuccess(t *testing.T){
	db := SetupNewDB()
	truncateCategory(db)
	router := SetupRouter(db)

	tx,_ := db.Begin()
	repository := repository.NewCategoryRepository()
	category := repository.Save(context.Background(),tx,domain.Category{
		Name:"Gadget",
	})
	tx.Commit()

	requestBody := strings.NewReader(`{"name":"Batu"}`)
	request := httptest.NewRequest(http.MethodPut,"http://localhost:3000/api/categories/"+ strconv.FormatInt(category.Id,10),requestBody)
	request.Header.Add("X-API-Key","RAHASIA")
	request.Header.Add("Content-Type","application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder,request)

	response := recorder.Result()

	assert.Equal(t,200,response.StatusCode)
}

func TestUpdateCategoryFailed(t *testing.T){
	db := SetupNewDB()
	truncateCategory(db)
	router := SetupRouter(db)

	tx,_ := db.Begin()
	repository := repository.NewCategoryRepository()
	category := repository.Save(context.Background(),tx,domain.Category{
		Name:"Gadget",
	})
	tx.Commit()

	requestBody := strings.NewReader(`{"name":""}`)
	request := httptest.NewRequest(http.MethodPut,"http://localhost:3000/api/categories/"+ strconv.FormatInt(category.Id,10),requestBody)
	request.Header.Add("X-API-Key","RAHASIA")
	request.Header.Add("Content-Type","application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder,request)

	response := recorder.Result()

	assert.Equal(t,400,response.StatusCode)
}

func TestGetCategorySuccess(t *testing.T){
	db := SetupNewDB()
	truncateCategory(db)
	router := SetupRouter(db)

	tx,_ := db.Begin()
	repository := repository.NewCategoryRepository()
	category := repository.Save(context.Background(),tx,domain.Category{
		Name:"Gadget",
	})
	tx.Commit()

	request := httptest.NewRequest(http.MethodGet,"http://localhost:3000/api/categories/"+ strconv.FormatInt(category.Id,10),nil)
	request.Header.Add("X-API-Key","RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder,request)

	response := recorder.Result()

	assert.Equal(t,200,response.StatusCode)
}

func TestGetCategoryFailed(t *testing.T){
	db := SetupNewDB()
	truncateCategory(db)
	router := SetupRouter(db)

	request := httptest.NewRequest(http.MethodGet,"http://localhost:3000/api/categories/404",nil)
	request.Header.Add("X-API-Key","RAHASIA")
	request.Header.Add("Content-Type","application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder,request)

	response := recorder.Result()

	assert.Equal(t,404,response.StatusCode)
}

func TestDeleteCategorySuccess(t *testing.T){
	db := SetupNewDB()
	truncateCategory(db)
	router := SetupRouter(db)

	tx,_ := db.Begin()
	repository := repository.NewCategoryRepository()
	category := repository.Save(context.Background(),tx,domain.Category{
		Name:"Gadget",
	})
	tx.Commit()

	request := httptest.NewRequest(http.MethodDelete,"http://localhost:3000/api/categories/"+ strconv.FormatInt(category.Id,10),nil)
	request.Header.Add("X-API-Key","RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder,request)

	response := recorder.Result()

	assert.Equal(t,200,response.StatusCode)
}

func TestDeleteCategoryFailed(t *testing.T){
	db := SetupNewDB()
	truncateCategory(db)
	router := SetupRouter(db)

	request := httptest.NewRequest(http.MethodDelete,"http://localhost:3000/api/categories/404",nil)
	request.Header.Add("X-API-Key","RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder,request)

	response := recorder.Result()

	assert.Equal(t,404,response.StatusCode)
}

func TestListCategorySuccess(t *testing.T){
	db := SetupNewDB()
	truncateCategory(db)
	router := SetupRouter(db)

	tx,_ := db.Begin()
	repository := repository.NewCategoryRepository()
	repository.Save(context.Background(),tx,domain.Category{
		Name:"Gadget",
	})
	tx.Commit()

	request := httptest.NewRequest(http.MethodGet,"http://localhost:3000/api/categories",nil)
	request.Header.Add("X-API-Key","RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder,request)

	response := recorder.Result()

	assert.Equal(t,200,response.StatusCode)
}

func TestUnauthorized(t *testing.T){
	db := SetupNewDB()
	truncateCategory(db)
	router := SetupRouter(db)

	requestBody := strings.NewReader(`{"name":"Gadget"}`)
	request := httptest.NewRequest(http.MethodPost,"http://localhost:3000/api/categories",requestBody)
	request.Header.Add("Content-Type","application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder,request)

	response := recorder.Result()

	assert.Equal(t,401,response.StatusCode)
}