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

func AddParcel(w http.ResponseWriter, r *http.Request) {
	var newEvent *dto.Parcel
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
	go insertIntoES(newEvent, c)

	parcelProfiler := db.ParcelsImpl{
		DB: global.DB,
	}
	err = parcelProfiler.AddParcel(newEvent)
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

func insertIntoES(parcel *dto.Parcel, c chan error) {
	esObj := dto.ESParcel{
		MyLoc: dto.Loc{
			Lat: parcel.SourceLatitude,
			Long: parcel.SourceLongitude,
		},
		UserName: parcel.UserName,
		Note: parcel.Note,
		Length: parcel.Length,
		Breadth: parcel.Breadth,
		Height: parcel.Height,
		Weight: parcel.Weight,
		Category: parcel.Category,
		SourceAddress: parcel.SourceAddress,
		DestinationAddress: parcel.DestinationAddress,
		SourceLatitude: parcel.SourceLatitude,
		SourceLongitude: parcel.SourceLongitude,
		DestinationLatitude: parcel.DestinationLatitude,
		DestinationLongitude: parcel.DestinationLongitude,
		CreatedAt: parcel.CreatedAt,
		Status: parcel.Status,
		Price: parcel.Price,
		CompletedAt: parcel.CompletedAt,
	}

	payload, _ := json.Marshal(esObj)
	b := bytes.NewBuffer(payload)
	_, err := global.ES.Index("test", b, global.ES.Index.WithDocumentID(parcel.ID))
	if err != nil {
		fmt.Println("error adding parcel to elastic search:", err.Error())
	}
	c <- err
}
