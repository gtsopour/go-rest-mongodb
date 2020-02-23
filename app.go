package main
import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"go-rest-mongodb/handler"
	. "go-rest-mongodb/config"
)

var config = Config{}

func init() {
  config.Read()
  fmt.Println(config.Server)
  fmt.Println(config.Database)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/places", handler.GetAllPlaces).Methods("GET")
	r.HandleFunc("/places/{id}", handler.GetPlaceById).Methods("GET")
	r.HandleFunc("/places", handler.CreatePlace).Methods("POST")
	r.HandleFunc("/places", handler.UpdatePlace).Methods("PUT")
	r.HandleFunc("/places", handler.DeletePlace).Methods("DELETE")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
