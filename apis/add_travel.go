package apis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"parcelDelivery/db"
	"parcelDelivery/dto"
	"parcelDelivery/global"
	"time"
)

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
	currentTS := currentTime.Format("2006-01-02 15:04:05")
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
			_ = json.NewEncoder(w).Encode("new parcel added")
		}
	}
}

func insertTravelIntoES(travel *dto.Travel, c chan error) {
	esObj := dto.ESTravel{
		MyLoc: dto.Loc{
			Lat: travel.SourceLatitude,
			Long: travel.SourceLongitude,
		},
		UserName: travel.UserName,
		Note: travel.Note,
		Mode: travel.Mode,
		SourceAddress: travel.SourceAddress,
		DestinationAddress: travel.DestinationAddress,
		SourceLatitude: travel.SourceLatitude,
		SourceLongitude: travel.SourceLongitude,
		DestinationLatitude: travel.DestinationLatitude,
		DestinationLongitude: travel.DestinationLongitude,
		CreatedAt: travel.CreatedAt,
		Status: travel.Status,
		CompletedAt: travel.CompletedAt,
	}

	payload, _ := json.Marshal(esObj)
	b := bytes.NewBuffer(payload)
	_, err := global.ES.Index(global.ESTravelIndex, b, global.ES.Index.WithDocumentID(travel.ID))
	if err != nil {
		fmt.Println("error adding travel to elastic search:", err.Error())
	}
	c <- err
}
