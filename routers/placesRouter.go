package routers

import (
	"github.com/gorilla/mux"
	"go-rest-mongodb/controllers"
)

func HandlePlacesRouter(r *mux.Router) *mux.Router {
	s := r.PathPrefix("/places").Subrouter()
	s.HandleFunc("", controllers.GetAllPlaces).Methods("GET")
	s.HandleFunc("/{id}", controllers.GetPlaceById).Methods("GET")
	s.HandleFunc("", controllers.CreatePlace).Methods("POST")
	s.HandleFunc("", controllers.UpdatePlace).Methods("PUT")
	s.HandleFunc("", controllers.DeletePlace).Methods("DELETE")
	return s;
}
