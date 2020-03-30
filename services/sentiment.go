package services

import (
	"log"

	"github.com/dinopuguh/kawulo-temporal/database"
	"github.com/dinopuguh/kawulo-temporal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GroupSentimentByLocation(db *mongo.Database, restId string) []models.TemporalGroup {
	ctx := database.Ctx

	matchStage := bson.D{
		primitive.E{
			Key: "$match",
			Value: bson.D{
				primitive.E{
					Key:   "restaurant_id",
					Value: restId,
				},
			},
		},
	}
	groupStage := bson.D{
		primitive.E{
			Key: "$group",
			Value: bson.D{
				primitive.E{
					Key: "_id",
					Value: bson.D{
						primitive.E{
							Key:   "restaurant_id",
							Value: "$restaurant_id",
						},
						primitive.E{
							Key:   "month",
							Value: "$month",
						},
						primitive.E{
							Key:   "year",
							Value: "$year",
						},
					},
				},
				primitive.E{
					Key: "location",
					Value: bson.D{
						primitive.E{
							Key:   "$first",
							Value: "$location",
						},
					},
				},
				primitive.E{
					Key: "service",
					Value: bson.D{
						primitive.E{
							Key:   "$avg",
							Value: "$service",
						},
					},
				},
				primitive.E{
					Key: "value",
					Value: bson.D{
						primitive.E{
							Key:   "$avg",
							Value: "$value",
						},
					},
				},
				primitive.E{
					Key: "food",
					Value: bson.D{
						primitive.E{
							Key:   "$avg",
							Value: "$food",
						},
					},
				},
				primitive.E{
					Key: "vader",
					Value: bson.D{
						primitive.E{
							Key:   "$avg",
							Value: "$vader",
						},
					},
				},
				primitive.E{
					Key: "wordnet",
					Value: bson.D{
						primitive.E{
							Key:   "$avg",
							Value: "$wordnet",
						},
					},
				},
				primitive.E{
					Key: "count",
					Value: bson.D{
						primitive.E{
							Key:   "$sum",
							Value: 1,
						},
					},
				},
			},
		},
	}
	sortStage := bson.D{
		primitive.E{
			Key: "$sort",
			Value: bson.D{
				primitive.E{
					Key:   "_id.year",
					Value: 1,
				},
				primitive.E{
					Key:   "_id.month",
					Value: 1,
				},
			},
		},
	}

	csr, err := db.Collection("sentiment").Aggregate(ctx, mongo.Pipeline{matchStage, groupStage, sortStage})
	if err != nil {
		log.Fatal(err.Error())
	}
	defer csr.Close(ctx)

	var result []models.TemporalGroup
	if err = csr.All(ctx, &result); err != nil {
		panic(err)
	}

	return result
}
