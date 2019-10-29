package main
import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)
func GetAllPlaces(w http.ResponseWriter, r * http.Request) {
	fmt.Fprintln(w, "Not implemented")
}
func GetPlaceById(w http.ResponseWriter, r * http.Request) {
	fmt.Fprintln(w, "Not implemented")
}
func CreatePlace(w http.ResponseWriter, r * http.Request) {
	fmt.Fprintln(w, "Not implemented")
}
func UpdatePlace(w http.ResponseWriter, r * http.Request) {
	fmt.Fprintln(w, "Not implemented")
}
func DeletePlace(w http.ResponseWriter, r * http.Request) {
	fmt.Fprintln(w, "Not implemented")
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/places", GetAllPlaces).Methods("GET")
	r.HandleFunc("/places/{id}", GetPlaceById).Methods("GET")
	r.HandleFunc("/places", CreatePlace).Methods("POST")
	r.HandleFunc("/places", UpdatePlace).Methods("PUT")
	r.HandleFunc("/places", DeletePlace).Methods("DELETE")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
