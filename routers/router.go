package routers

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Routers() *mux.Router {
	//StrictSlash defines the trailing slash behavior for new routes. The initial value is false.
	//When true, if the route path is "/path/", accessing "/path" will perform a redirect to the former and vice versa.
	r := mux.NewRouter().StrictSlash(true)

	//PathPrefix /api adds a matcher for the URL path prefix.
	s := r.PathPrefix("/api").Subrouter()
	AddPlacesRouter(s)

	r.Use(loggingMiddleware)
	return r
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
