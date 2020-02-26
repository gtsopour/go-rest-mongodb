package repository

import (
	"context"
	"fmt"
	"go-rest-mongodb/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type placeRepository struct {
	collection *mongo.Collection
}

func PlaceRepository(collection *mongo.Collection) *placeRepository {
	return &placeRepository {
		collection: collection,
	}
}

// Get all Places
func (p *placeRepository) FindAll() ([]models.Place, error) {
	var places []models.Place

	findOptions := options.Find()
	findOptions.SetLimit(100)

	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	// Finding multiple documents returns a cursor
	cur, err := p.collection.Find(ctx, bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Iterate through the cursor
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result models.Place
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		places = append(places, result)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	return places, err
}

// Create a new Place
func (p *placeRepository) Insert(place models.Place) (interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := p.collection.InsertOne(ctx, &place)
	fmt.Println("Inserted a single document: ", result.InsertedID)
	return result.InsertedID, err
}

// Delete an existing Place
func (p *placeRepository) Delete(id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objectId}
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := p.collection.DeleteOne(ctx, filter)
	fmt.Println("Deleted a single document: ", result.DeletedCount)
	return err
}
