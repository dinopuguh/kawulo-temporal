package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TemporalGroup struct {
	ID       TemporalID `bson:"_id"`
	Location Location   `bson:"location"`
	Service  float64    `bson:"service"`
	Value    float64    `bson:"value"`
	Food     float64    `bson:"food"`
	Vader    float64    `bson:"vader"`
	Wordnet  float64    `bson:"wordnet"`
	Count    int32      `bson:"count"`
}

type Temporal struct {
	ID           primitive.ObjectID `bson:"_id"`
	LocationId   string             `bson:"location_id"`
	Location     Location           `bson:"location"`
	RestaurantId string             `bson:"restaurant_id"`
	Restaurant   Restaurant         `bson:"restaurant"`
	Month        int32              `bson:"month"`
	Year         int32              `bson:"year"`
	Service      float64            `bson:"service"`
	Value        float64            `bson:"value"`
	Food         float64            `bson:"food"`
	Vader        float64            `bson:"vader"`
	Wordnet      float64            `bson:"wordnet"`
	CreatedAt    time.Time          `bson:"created_at"`
}

type TemporalID struct {
	RestaurantId string `bson:"restaurant_id"`
	Month        int32  `bson:"month"`
	Year         int32  `bson:"year"`
}
