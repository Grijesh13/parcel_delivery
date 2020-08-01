package apis

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"parcelDelivery/db"
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
	userParcelsSorted := make(map[string][]dto2.Parcel)
	for i := 0; i < len(userParcels); i++ {
		if parcels, ok := userParcelsSorted[userParcels[i].Status]; ok {
			parcels = append(parcels, *userParcels[i])
			userParcelsSorted[userParcels[i].Status] = parcels
		} else {
			var newParcels []dto2.Parcel
			newParcels = append(newParcels, *userParcels[i])
			userParcelsSorted[userParcels[i].Status] = newParcels
		}
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(userParcelsSorted)
}
