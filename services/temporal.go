package services

import (
	"log"

	"github.com/dinopuguh/kawulo-temporal/database"
	"github.com/dinopuguh/kawulo-temporal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindTemporalByRestaurantDate(db *mongo.Database, restId string, month int32, year int32) (bool, models.Temporal) {
	ctx := database.Ctx

	var result models.Temporal
	err := db.Collection("temporal").FindOne(ctx, bson.M{"restaurant_id": restId, "month": month, "year": year}).Decode(&result)
	if err != nil {
		return false, result
	}

	return true, result
}

func InsertTemporal(db *mongo.Database, temporal models.Temporal) {
	ctx := database.Ctx

	crs, err := db.Collection("temporal").InsertOne(ctx, temporal)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Insert temporal success -", crs.InsertedID)
}

func UpdateTemporalById(db *mongo.Database, temporalId primitive.ObjectID, temporal models.Temporal) {
	ctx := database.Ctx

	updated := bson.D{
		{Key: "month", Value: temporal.Month},
		{Key: "year", Value: temporal.Year},
		{Key: "service", Value: temporal.Service},
		{Key: "value", Value: temporal.Value},
		{Key: "food", Value: temporal.Food},
		{Key: "vader", Value: temporal.Vader},
		{Key: "wordnet", Value: temporal.Wordnet},
		{Key: "created_at", Value: temporal.CreatedAt},
	}

	_, err := db.Collection("temporal").UpdateOne(ctx, bson.M{"_id": temporalId}, bson.D{{Key: "$set", Value: updated}})

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Update temporal success -", temporalId)
}
