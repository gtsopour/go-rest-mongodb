package handler

import (
  "fmt"
	"encoding/json"
	"go-rest-mongodb/config"
	"go-rest-mongodb/model"
	"go-rest-mongodb/repository/place"
	 "gopkg.in/mgo.v2/bson"
	"net/http"
)

var GetAllPlaces = func(w http.ResponseWriter, r * http.Request) {
//   db, err := config.GetMongoDB()
// 	if err != nil {
// 		json.NewEncoder(w).Encode(err)
// 	}
// 	fmt.Println(db)
// 	placeRepository := repository.PlaceRepository(db, "Places")
// 	places, err := placeRepository.FindAll()
//   if err != nil {
//     json.NewEncoder(w).Encode(err)
//   }
//   fmt.Println(places)
//   json.NewEncoder(w).Encode(places)

  db, err := config.GetMongoDB()
  if err != nil {
    json.NewEncoder(w).Encode(err)
  }
  placeRepository := repository.PlaceRepository(db, "Places")
  places, err := placeRepository.FindAll()
  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }
  respondWithJson(w, http.StatusOK, places)
}

var GetPlaceById = func(w http.ResponseWriter, r * http.Request) {
	fmt.Fprintln(w, "Not implemented")
}

var CreatePlace = func(w http.ResponseWriter, r * http.Request) {
	db, err := config.GetMongoDB()
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	placeRepository := repository.PlaceRepository(db, "Places")
// 	decoder := json.NewDecoder(r.Body)

// 	var p model.Place
// 	var paramBody model.Place
//
// 	decoder.Decode(&paramBody)
//
// 	p.Title = paramBody.Title
// 	p.Description = paramBody.Description
// // 	p.CreatedAt = time.Now()
// // 	p.UpdatedAt = time.Now()
//   fmt.Println(p)
// 	err = placeRepository.Insert(p)
//   fmt.Println(err)
// 	if err != nil {
// 		json.NewEncoder(w).Encode(err)
// 	} else {
// 		json.NewEncoder(w).Encode("")
// 	}

	defer r.Body.Close()
  var place model.Place
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
	fmt.Fprintln(w, "Not implemented")
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
