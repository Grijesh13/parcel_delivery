package apis

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"parcelDelivery/db"
	"parcelDelivery/dto"
	"parcelDelivery/global"
	"time"
)

// AddTravel ...
func AddTravel(w http.ResponseWriter, r *http.Request) {
	var newEvent *dto.Travel
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("error in reading data")
		return
	}
	err = json.Unmarshal(reqBody, &newEvent)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("error in unmarshalling data")
		return
	}

	// set the created_at for the new user to be added
	currentTime := time.Now()
	currentTS := currentTime.Format("2006-01-02")
	newEvent.CreatedAt = currentTS

	// set the status
	newEvent.Status = global.StatusPending

	c := make(chan error)
	go insertTravelIntoES(newEvent, c)

	travelProfiler := db.TravelsImpl{
		DB: global.DB,
	}
	err = travelProfiler.AddTravel(newEvent)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode("problem")
		return
	}

	for i := 0; i < 1; i++ {
		select {
		case e := <-c:
			if e != nil {
				w.WriteHeader(http.StatusInternalServerError)
				_ = json.NewEncoder(w).Encode("problem")
				return
			}
			w.WriteHeader(http.StatusOK)
			_ = json.NewEncoder(w).Encode("new travel added")
		}
	}
}

func insertTravelIntoES(travel *dto.Travel, c chan error) {
	esObj := dto.ESTravel{
		MySrcLoc: dto.Loc{
			Lat:  travel.SourceLatitude,
			Long: travel.SourceLongitude,
		},
		MyDestLoc: dto.Loc{
			Lat:  travel.DestinationLatitude,
			Long: travel.DestinationLongitude,
		},
		UserName:             travel.UserName,
		Note:                 travel.Note,
		Mode:                 travel.Mode,
		SourceAddress:        travel.SourceAddress,
		DestinationAddress:   travel.DestinationAddress,
		SourceLatitude:       travel.SourceLatitude,
		SourceLongitude:      travel.SourceLongitude,
		DestinationLatitude:  travel.DestinationLatitude,
		DestinationLongitude: travel.DestinationLongitude,
		CreatedAt:            travel.CreatedAt,
		Status:               travel.Status,
		StartDate:            travel.StartDate,
		EndDate:              travel.EndDate,
		CompletedAt:          travel.CompletedAt,
	}

	//payload, _ := json.Marshal(esObj)
	//b := bytes.NewBuffer(payload)
	//_, err := global.ES.Index(global.ESTravelIndex, b, global.ES.Index.WithDocumentID(travel.ID))

	_, err := global.ES2.Index().
		Index(global.ESTravelIndex).
		Id(travel.ID).
		BodyJson(esObj).
		Do(context.Background())

	if err != nil {
		fmt.Println("error adding travel to elastic search:", err.Error())
	}
	c <- err
}
