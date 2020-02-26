package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go-rest-mongodb/models"
	"go-rest-mongodb/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"time"
)

func init() {
	fmt.Println("Init() Controllers")
}

var GetAllPlaces = func(w http.ResponseWriter, r * http.Request) {
	// Connect to DB
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost"))
	if err != nil {
		respondWithError(w, http.StatusServiceUnavailable, err.Error())
		return
	}
	collection := client.Database("go-rest-mongodb").Collection("Places")
	placeRepository := repository.PlaceRepository(collection)

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
	// Connect to DB
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost"))
	if err != nil {
		respondWithError(w, http.StatusServiceUnavailable, err.Error())
		return
	}
	collection := client.Database("go-rest-mongodb").Collection("Places")
	placeRepository := repository.PlaceRepository(collection)

	defer r.Body.Close()
	var place models.Place
	if err := json.NewDecoder(r.Body).Decode(&place); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	insertResult, err := placeRepository.Insert(place);
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
	// Connect to DB
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost"))
	if err != nil {
		respondWithError(w, http.StatusServiceUnavailable, err.Error())
		return
	}
	collection := client.Database("go-rest-mongodb").Collection("Places")
	placeRepository := repository.PlaceRepository(collection)

	//Vars returns the route variables for the current request, if any.
	vars := mux.Vars(r)
	//Get id from the current request
	id := vars["id"]
	fmt.Println(id)
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
