package apis

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"parcelDelivery/db"
	"parcelDelivery/dto"
	"parcelDelivery/global"
	"strings"
	"time"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
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

	// set the created_at for the new user to be added
	currentTime := time.Now()
	currentTS := currentTime.Format("2006-01-02 15:04:05")
	newEvent.CreatedAt = currentTS

	userProfiler := db.UserProfileImpl{
		DB: global.DB,
	}
	err = userProfiler.AddUser(newEvent)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			w.WriteHeader(http.StatusConflict)
			_ = json.NewEncoder(w).Encode("username already in use")
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode("problem")
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode("new user added")
}
