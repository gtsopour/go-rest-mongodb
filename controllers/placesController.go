package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go-rest-mongodb/config"
	"go-rest-mongodb/models"
	"go-rest-mongodb/repository"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

var GetAllPlaces = func(w http.ResponseWriter, r * http.Request) {
	db, err := config.GetDB()
	if err != nil {
		respondWithError(w, http.StatusServiceUnavailable, err.Error())
		return
	}

	placeRepository := repository.PlaceRepository(db)
	places, err := placeRepository.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, places)
}

var GetPlaceById = func(w http.ResponseWriter, r * http.Request) {
	fmt.Fprintln(w, "Not implemented!")
}

var CreatePlace = func(w http.ResponseWriter, r * http.Request) {
	db, err := config.GetDB()
	if err != nil {
		respondWithError(w, http.StatusServiceUnavailable, err.Error())
		return
	}

	placeRepository := repository.PlaceRepository(db)
	defer r.Body.Close()
	var place models.Place
	if err := json.NewDecoder(r.Body).Decode(&place); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	place.ID = bson.NewObjectId()
	if err := placeRepository.Insert(place); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, place)
}

var UpdatePlace = func(w http.ResponseWriter, r * http.Request) {
	fmt.Fprintln(w, "Not implemented")
}

var DeletePlace = func(w http.ResponseWriter, r * http.Request) {
	db, err := config.GetDB()
	if err != nil {
		respondWithError(w, http.StatusServiceUnavailable, err.Error())
		return
	}

	//Vars returns the route variables for the current request, if any.
	vars := mux.Vars(r)
	//Get id from the current request
	id := vars["id"]

	placeRepository := repository.PlaceRepository(db)
	if err := placeRepository.Delete(id); err != nil {
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
