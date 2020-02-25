package routers

import (
	"github.com/gorilla/mux"
	"go-rest-mongodb/controllers"
)

func PlaceRouter(r *mux.Router) *mux.Router {
	r.HandleFunc("/places", controllers.GetAllPlaces).Methods("GET")
	r.HandleFunc("/places/{id}", controllers.GetPlaceById).Methods("GET")
	r.HandleFunc("/places", controllers.CreatePlace).Methods("POST")
	r.HandleFunc("/places", controllers.UpdatePlace).Methods("PUT")
	r.HandleFunc("/places", controllers.DeletePlace).Methods("DELETE")
	return r;
}
