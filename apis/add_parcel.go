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

// AddParcel ...
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
	go insertParcelIntoES(newEvent, c)

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

func insertParcelIntoES(parcel *dto.Parcel, c chan error) {
	numItems := 0
	netWeight := 0
	categoriesMap := make(map[string]bool)
	var categories []string
	for i := 0; i < len(parcel.Items); i++ {
		numItems += parcel.Items[i].Number
		netWeight += parcel.Items[i].Number * parcel.Items[i].Weight
		categoriesMap[parcel.Items[i].Category] = true
	}
	for category := range categoriesMap {
		categories = append(categories, category)
	}

	sDate, _ := time.Parse("2006-01-02 15:04:05", parcel.ShipDate)
	uDate := sDate.Format("2006-01-02")

	esObj := dto.ESParcel{
		MySrcLoc: dto.Loc{
			Lat:  parcel.SourceLatitude,
			Long: parcel.SourceLongitude,
		},
		MyDestLoc: dto.Loc{
			Lat:  parcel.DestinationLatitude,
			Long: parcel.DestinationLongitude,
		},
		PickUpStart:          uDate,
		UserName:             parcel.UserName,
		Note:                 parcel.Note,
		SourceAddress:        parcel.SourceAddress,
		DestinationAddress:   parcel.DestinationAddress,
		SourceLatitude:       parcel.SourceLatitude,
		SourceLongitude:      parcel.SourceLongitude,
		DestinationLatitude:  parcel.DestinationLatitude,
		DestinationLongitude: parcel.DestinationLongitude,
		CreatedAt:            parcel.CreatedAt,
		Status:               parcel.Status,
		Price:                parcel.Price,
		CompletedAt:          parcel.CompletedAt,
		IsNegotiable:         parcel.IsNegotiable,
		ShipDate:             parcel.ShipDate,
		NumberItems:          numItems,
		NetWeight:            netWeight,
		Categories:           categories,
		Items:                parcel.Items,
	}

	//payload, _ := json.Marshal(esObj)
	//b := bytes.NewBuffer(payload)
	//_, err := global.ES.Index(global.ESParcelIndex, b, global.ES.Index.WithDocumentID(parcel.ID))

	_, err := global.ES2.Index().
		Index(global.ESParcelIndex).
		Id(parcel.ID).
		BodyJson(esObj).
		Do(context.Background())

	if err != nil {
		fmt.Println("error adding parcel to elastic search:", err.Error())
	}
	c <- err
}
