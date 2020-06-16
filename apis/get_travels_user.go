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

func GetTravelsForUser(w http.ResponseWriter, r *http.Request) {
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

	travelImpl := db.TravelsImpl{
		DB : global.DB,
	}
	userTravels := travelImpl.GetTravels(newEvent.UserName)
	userTravelsSorted := make(map[string][]dto2.Travel)
	for i := 0; i < len(userTravels); i++ {
		if travels, ok := userTravelsSorted[userTravels[i].Status]; ok {
			travels = append(travels, *userTravels[i])
			userTravelsSorted[userTravels[i].Status] = travels
		} else {
			var newTravels []dto2.Travel
			newTravels = append(newTravels, *userTravels[i])
			userTravelsSorted[userTravels[i].Status] = newTravels
		}
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(userTravelsSorted)
}
