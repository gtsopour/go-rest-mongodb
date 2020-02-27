package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"go-rest-mongodb/models"
	. "go-rest-mongodb/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

var placesRepository PlacesRepository

var GetAllPlaces = func(w http.ResponseWriter, r * http.Request) {
	places, err := placesRepository.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	log.Info(places);
	respondWithJson(w, http.StatusOK, places)
}

var GetPlaceById = func(w http.ResponseWriter, r * http.Request) {
	fmt.Fprintln(w, "Not implemented!")
}

var CreatePlace = func(w http.ResponseWriter, r * http.Request) {
	defer r.Body.Close()
	var place models.Place
	if err := json.NewDecoder(r.Body).Decode(&place); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	log.Info(place);
	insertResult, err := placesRepository.Insert(place);
	if  err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	place.ID = insertResult.(primitive.ObjectID)
	respondWithJson(w, http.StatusCreated, place)
}

var UpdatePlace = func(w http.ResponseWriter, r * http.Request) {
	fmt.Fprintln(w, "Not implemented")
}

var DeletePlace = func(w http.ResponseWriter, r * http.Request) {
	//Vars returns the route variables for the current request, if any.
	vars := mux.Vars(r)
	//Get id from the current request
	id := vars["id"]
	fmt.Println(id)
	if err := placesRepository.Delete(id); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
