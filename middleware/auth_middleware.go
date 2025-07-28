package middleware

import (
	"ArthaFreestyle/go-rest-api/model/web"
	"encoding/json"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Header.Get("X-API-Key") == "RAHASIA" {
		middleware.Handler.ServeHTTP(w, r)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
		}

		encoder := json.NewEncoder(w)
		err := encoder.Encode(webResponse)
		if err != nil {
			panic(err)
		}
	}
}
