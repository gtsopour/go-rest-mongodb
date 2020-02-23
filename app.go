package main
import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"go-rest-mongodb/controllers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/places", controllers.GetAllPlaces).Methods("GET")
	r.HandleFunc("/places/{id}", controllers.GetPlaceById).Methods("GET")
	r.HandleFunc("/places", controllers.CreatePlace).Methods("POST")
	r.HandleFunc("/places", controllers.UpdatePlace).Methods("PUT")
	r.HandleFunc("/places", controllers.DeletePlace).Methods("DELETE")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
