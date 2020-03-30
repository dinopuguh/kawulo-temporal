package services

import (
	"log"

	"github.com/dinopuguh/kawulo-temporal/database"
	"github.com/dinopuguh/kawulo-temporal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindAllLocations(db *mongo.Database) []models.Location {
	ctx := database.Ctx

	csr, err := db.Collection("location").Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err.Error())
	}
	defer csr.Close(ctx)

	result := make([]models.Location, 0)
	for csr.Next(ctx) {
		var row models.Location
		err := csr.Decode(&row)
		if err != nil {
			log.Fatal(err.Error())
		}

		result = append(result, row)
	}

	return result
}

func FindIndonesianLocations(db *mongo.Database) []models.Location {
	ctx := database.Ctx

	cities := []string{
		"Surabaya",
		"Jakarta",
		"Bandung",
		"Semarang",
		"Yogyakarta Region",
		"Denpasar",
		"Banda Aceh",
		"Medan",
		"Padang",
		"Pekanbaru",
		"Palembang",
		"Bengkulu",
		"Bandar Lampung",
		"Pangkal Pinang",
		"Tanjung Pinang",
		"Serang",
		"Mataram",
		"Kupang",
		"Pontianak",
		"Banjarmasin",
		"Samarinda",
		"Manado",
		"Palu",
		"Makassar",
		"Kendari",
		"Gorontalo",
		"Mamuju",
		"Ambon",
		"Jayapura",
		"Manokwari",
	}

	result := make([]models.Location, 0)

	for _, city := range cities {
		csr, err := db.Collection("location").Find(ctx, bson.M{"name": city})
		if err != nil {
			log.Fatal(err.Error())
		}
		defer csr.Close(ctx)

		for csr.Next(ctx) {
			var row models.Location
			err := csr.Decode(&row)
			if err != nil {
				log.Fatal(err.Error())
			}

			result = append(result, row)
		}
	}

	return result
}

func FindLocationById(db *mongo.Database, locId string) models.Location {
	ctx := database.Ctx

	var result models.Location
	err := db.Collection("location").FindOne(ctx, bson.M{"location_id": locId}).Decode(&result)
	if err != nil {
		log.Fatal(err.Error())
	}

	return result
}

func FindLocationByQuery(db *mongo.Database, query string) []models.Location {
	ctx := database.Ctx

	filter := bson.M{"name": primitive.Regex{Pattern: "^" + query + ".*", Options: "i"}}
	csr, err := db.Collection("location").Find(ctx, filter)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer csr.Close(ctx)

	result := make([]models.Location, 0)
	for csr.Next(ctx) {
		var row models.Location
		err := csr.Decode(&row)
		if err != nil {
			log.Fatal(err.Error())
		}

		result = append(result, row)
	}

	return result
}
