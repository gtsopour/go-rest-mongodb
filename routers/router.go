package routers

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Routers() *mux.Router {
	r := mux.NewRouter().StrictSlash(false)
	s := r.PathPrefix("/api").Subrouter()
	s = PlaceRouter(s)
	s.Use(loggingMiddleware)
	return s
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
