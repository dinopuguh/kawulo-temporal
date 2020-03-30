package main

import (
	"log"
	"time"

	"github.com/dinopuguh/kawulo-temporal/database"
	"github.com/dinopuguh/kawulo-temporal/models"
	"github.com/dinopuguh/kawulo-temporal/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	locs := services.FindIndonesianLocations(db)

	for _, loc := range locs {
		log.Println("------------------------", loc.Name)
		rests := services.FindRestaurantByLocation(db, loc.LocationId)

		for _, rest := range rests {
			log.Println("---------------------", rest.Name)
			sentiments := services.GroupSentimentByLocation(db, rest.LocationId)
			for t, sent := range sentiments {

				thisMonth := sent.ID.Month
				thisYear := sent.ID.Year

				nextMonth := int32(time.Now().Month())
				nextYear := int32(time.Now().Year())

				if t+1 < len(sentiments) {
					nextMonth = sentiments[t+1].ID.Month
					nextYear = sentiments[t+1].ID.Year
				}

				for year := thisYear; year < nextYear+1; year++ {
					for month := thisMonth; month < 13; month++ {
						if nextMonth < month && year == nextYear {
							break
						}

						log.Printf("------ [%s] - %d (%d/%d) -> (%d/%d)", rest.LocationId, t+1, month, year, nextMonth, nextYear)

						temporal := models.Temporal{
							ID:           primitive.NewObjectID(),
							LocationId:   rest.LocationID,
							Location:     sent.Location,
							RestaurantId: rest.LocationId,
							Restaurant:   rest,
							Month:        month,
							Year:         year,
							Service:      sent.Service,
							Value:        sent.Value,
							Food:         sent.Food,
							Vader:        sent.Vader,
							Wordnet:      sent.Wordnet,
							CreatedAt:    time.Now(),
						}

						log.Println(temporal.Month, temporal.Year, temporal.Vader)

						temporalExist, temporalByDate := services.FindTemporalByRestaurantDate(db, rest.LocationId, month, year)

						if temporalExist == true {
							services.UpdateTemporalById(db, temporalByDate.ID, temporal)
						} else {
							services.InsertTemporal(db, temporal)
						}

					}
					thisMonth = 1
				}

			}
		}
	}
}
