package apis

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"parcelDelivery/db"
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
		DB : global.DB,
	}
	userParcels := parcelImpl.GetParcels(newEvent.UserName)

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(userParcels)
}
