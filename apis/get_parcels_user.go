package apis

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"parcelDelivery/db"
	dto3 "parcelDelivery/dto/response"
	dto2 "parcelDelivery/dto"
	"parcelDelivery/global"
	dto "parcelDelivery/request_dto"
)

func GetParcelsForUser(w http.ResponseWriter, r *http.Request) {
	var newEvent *dto.User
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

	parcelImpl := db.ParcelsImpl{
		DB: global.DB,
	}
	userParcels := parcelImpl.GetParcels(newEvent.UserName)
	userParcelsSorted := make(map[string][]dto3.Parcel)
	for i := 0; i < len(userParcels); i++ {
		respParcel := convertDataToResponseParcel(userParcels[i])
		if parcels, ok := userParcelsSorted[userParcels[i].Status]; ok {
			parcels = append(parcels, *respParcel)
			userParcelsSorted[userParcels[i].Status] = parcels
		} else {
			var newParcels []dto3.Parcel
			newParcels = append(newParcels, *respParcel)
			userParcelsSorted[userParcels[i].Status] = newParcels
		}
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(userParcelsSorted)
}

func convertDataToResponseParcel(userParcel *dto2.Parcel) *dto3.Parcel {
	var items []dto3.Item
	er := json.Unmarshal([]byte(userParcel.SQLItems), &items)
	if er != nil {
		fmt.Printf("error unmarshalling sql_items for parcels with error = %s", er.Error())
		return nil
	}

  respParcel := dto3.Parcel{}
	respParcel.Items = items

	numItems := 0
	netWeight := 0
	categoriesMap := make(map[string]bool)
	var categories []string
	for i := 0; i < len(items); i++ {
		numItems += items[i].Number
		netWeight += items[i].Number * items[i].Weight
		categoriesMap[items[i].Category] = true
	}
	for category := range categoriesMap {
		categories = append(categories, category)
	}

	respParcel.ID = userParcel.ID
	respParcel.UserName = userParcel.UserName
	respParcel.Note = userParcel.Note
	respParcel.SourceAddress = userParcel.SourceAddress
	respParcel.DestinationAddress = userParcel.DestinationAddress
	respParcel.SourceLatitude = userParcel.SourceLatitude
	respParcel.SourceLongitude = userParcel.SourceLongitude
	respParcel.DestinationLatitude = userParcel.DestinationLatitude
	respParcel.DestinationLongitude = userParcel.DestinationLongitude
	respParcel.PickUpStart = userParcel.PickUpStart
	respParcel.PickUpEnd = userParcel.PickUpEnd
	respParcel.CreatedAt = userParcel.CreatedAt
	respParcel.Status = userParcel.Status
	respParcel.Price = userParcel.Price
	respParcel.IsNegotiable = userParcel.IsNegotiable
	respParcel.CompletedAt = userParcel.CompletedAt
	respParcel.NumberItems = numItems
	respParcel.NetWeight = netWeight
	respParcel.Categories = categories

	return &respParcel
}
